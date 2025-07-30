package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	redis "github.com/go-redis/redis/v8"

	"github.com/ct-zh/go-redis-proxy/pkg/errors"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// ConnectionManager 管理Redis连接池和token
type ConnectionManager struct {
	// 存储token到连接信息的映射
	connections map[string]*types.ConnectionInfo
	// 存储token到Redis客户端的映射
	clients map[string]*redis.Client
	// 读写锁保护并发访问
	mutex sync.RWMutex
	// 清理过期连接的定时器
	cleanupTicker *time.Ticker
	// 停止清理的通道
	stopCleanup chan bool
}

// NewConnectionManager 创建新的连接管理器
func NewConnectionManager() *ConnectionManager {
	cm := &ConnectionManager{
		connections:   make(map[string]*types.ConnectionInfo),
		clients:       make(map[string]*redis.Client),
		cleanupTicker: time.NewTicker(5 * time.Minute), // 每5分钟清理一次过期连接
		stopCleanup:   make(chan bool),
	}

	// 启动清理协程
	go cm.startCleanup()

	return cm
}

// Connect 建立Redis连接并返回token
func (cm *ConnectionManager) Connect(ctx context.Context, req *types.ConnectRequest) (*types.ConnectResponse, error) {
	// 生成唯一token
	token, err := cm.generateToken()
	if err != nil {
		return nil, errors.NewError(errors.CodeTokenGenerationFailed)
	}

	// 生成连接ID
	connID := cm.generateConnID()

	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     req.Addr,
		Password: req.Password,
		DB:       req.DB,
	})

	// 测试连接
	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, errors.NewError(errors.CodeRedisConnectFailed)
	}

	// 设置默认TTL
	ttl := req.TokenTTL
	if ttl <= 0 {
		ttl = 3600 // 默认1小时
	}

	now := time.Now()
	expiresAt := now.Add(time.Duration(ttl) * time.Second)

	// 存储连接信息
	connInfo := &types.ConnectionInfo{
		RedisConfig: req.RedisRequest,
		CreatedAt:   now,
		ExpiresAt:   expiresAt,
		ConnID:      connID,
		LastUsed:    now,
	}

	cm.mutex.Lock()
	cm.connections[token] = connInfo
	cm.clients[token] = client
	cm.mutex.Unlock()

	return &types.ConnectResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		ConnID:    connID,
	}, nil
}

// GetClient 根据token获取Redis客户端
func (cm *ConnectionManager) GetClient(token string) (*redis.Client, error) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	connInfo, exists := cm.connections[token]
	if !exists {
		return nil, errors.NewError(errors.CodeTokenNotFound)
	}

	// 检查token是否过期
	if time.Now().After(connInfo.ExpiresAt) {
		return nil, errors.NewError(errors.CodeTokenExpired)
	}

	client, exists := cm.clients[token]
	if !exists {
		return nil, errors.NewError(errors.CodeConnectionNotFound)
	}

	// 更新最后使用时间
	connInfo.LastUsed = time.Now()

	return client, nil
}

// RefreshToken 刷新token的过期时间
func (cm *ConnectionManager) RefreshToken(req *types.TokenRefreshRequest) (*types.ConnectResponse, error) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	connInfo, exists := cm.connections[req.Token]
	if !exists {
		return nil, errors.NewError(errors.CodeTokenNotFound)
	}

	// 检查token是否已过期
	if time.Now().After(connInfo.ExpiresAt) {
		return nil, errors.NewError(errors.CodeTokenExpired)
	}

	// 设置新的TTL
	ttl := req.TokenTTL
	if ttl <= 0 {
		ttl = 3600 // 默认1小时
	}

	newExpiresAt := time.Now().Add(time.Duration(ttl) * time.Second)
	connInfo.ExpiresAt = newExpiresAt
	connInfo.LastUsed = time.Now()

	return &types.ConnectResponse{
		Token:     req.Token,
		ExpiresAt: newExpiresAt,
		ConnID:    connInfo.ConnID,
	}, nil
}

// Disconnect 断开连接并清理资源
func (cm *ConnectionManager) Disconnect(token string) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	client, exists := cm.clients[token]
	if exists {
		client.Close()
		delete(cm.clients, token)
	}

	delete(cm.connections, token)
	return nil
}

// GetStats 获取连接池统计信息
func (cm *ConnectionManager) GetStats() *types.ConnectionStats {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()

	activeCount := 0
	for _, connInfo := range cm.connections {
		if time.Now().Before(connInfo.ExpiresAt) {
			activeCount++
		}
	}

	return &types.ConnectionStats{
		TotalConnections:  len(cm.connections),
		ActiveConnections: activeCount,
		PoolSize:         len(cm.clients),
		IdleConnections:  len(cm.clients) - activeCount,
	}
}

// generateToken 生成随机token
func (cm *ConnectionManager) generateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// generateConnID 生成连接ID
func (cm *ConnectionManager) generateConnID() string {
	return fmt.Sprintf("conn_%d", time.Now().UnixNano())
}

// startCleanup 启动清理过期连接的协程
func (cm *ConnectionManager) startCleanup() {
	for {
		select {
		case <-cm.cleanupTicker.C:
			cm.cleanupExpiredConnections()
		case <-cm.stopCleanup:
			cm.cleanupTicker.Stop()
			return
		}
	}
}

// cleanupExpiredConnections 清理过期的连接
func (cm *ConnectionManager) cleanupExpiredConnections() {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	now := time.Now()
	expiredTokens := make([]string, 0)

	for token, connInfo := range cm.connections {
		if now.After(connInfo.ExpiresAt) {
			expiredTokens = append(expiredTokens, token)
		}
	}

	for _, token := range expiredTokens {
		if client, exists := cm.clients[token]; exists {
			client.Close()
			delete(cm.clients, token)
		}
		delete(cm.connections, token)
	}
}

// Stop 停止连接管理器
func (cm *ConnectionManager) Stop() {
	cm.stopCleanup <- true

	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	// 关闭所有连接
	for token, client := range cm.clients {
		client.Close()
		delete(cm.clients, token)
		delete(cm.connections, token)
	}
}