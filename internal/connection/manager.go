package connection

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

// ConnectionManager 连接管理器，负责管理Redis连接池
type ConnectionManager struct {
	pools map[string]*redis.Client
	mutex sync.RWMutex
}

// NewConnectionManager 创建连接管理器
func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		pools: make(map[string]*redis.Client),
	}
}

// GetConnection 获取Redis连接
func (m *ConnectionManager) GetConnection(config *ConnectionConfig) (*redis.Client, error) {
	hash := config.Hash()
	
	// 先尝试读锁获取现有连接
	m.mutex.RLock()
	if client, exists := m.pools[hash]; exists {
		m.mutex.RUnlock()
		return client, nil
	}
	m.mutex.RUnlock()
	
	// 使用写锁创建新连接
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	// 双重检查，防止并发创建
	if client, exists := m.pools[hash]; exists {
		return client, nil
	}
	
	// 创建新的Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		MaxRetries:   config.MaxRetries,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,
		IdleTimeout:  5 * time.Minute,
	})
	
	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	
	m.pools[hash] = client
	return client, nil
}

// CloseConnection 关闭指定配置的连接
func (m *ConnectionManager) CloseConnection(config *ConnectionConfig) error {
	hash := config.Hash()
	
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	if client, exists := m.pools[hash]; exists {
		delete(m.pools, hash)
		return client.Close()
	}
	
	return nil
}

// CloseAll 关闭所有连接
func (m *ConnectionManager) CloseAll() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	var lastErr error
	for hash, client := range m.pools {
		if err := client.Close(); err != nil {
			lastErr = err
		}
		delete(m.pools, hash)
	}
	
	return lastErr
}

// GetStats 获取连接池统计信息
func (m *ConnectionManager) GetStats() map[string]interface{} {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	
	stats := make(map[string]interface{})
	stats["total_pools"] = len(m.pools)
	
	poolStats := make(map[string]interface{})
	for hash, client := range m.pools {
		poolStat := client.PoolStats()
		poolStats[hash] = map[string]interface{}{
			"hits":         poolStat.Hits,
			"misses":       poolStat.Misses,
			"timeouts":     poolStat.Timeouts,
			"total_conns":  poolStat.TotalConns,
			"idle_conns":   poolStat.IdleConns,
			"stale_conns":  poolStat.StaleConns,
		}
	}
	stats["pools"] = poolStats
	
	return stats
}

// HealthCheck 健康检查，检查所有连接池的状态
func (m *ConnectionManager) HealthCheck(ctx context.Context) map[string]error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	
	results := make(map[string]error)
	for hash, client := range m.pools {
		if err := client.Ping(ctx).Err(); err != nil {
			results[hash] = err
		} else {
			results[hash] = nil
		}
	}
	
	return results
}