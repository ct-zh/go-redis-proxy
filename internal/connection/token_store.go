package connection

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// MemoryTokenStore 内存Token存储实现
type MemoryTokenStore struct {
	tokens map[string]*TokenInfo
	mutex  sync.RWMutex
}

// TokenInfo Token信息
type TokenInfo struct {
	Config    *ConnectionConfig `json:"config"`
	CreatedAt time.Time         `json:"created_at"`
	ExpiresAt time.Time         `json:"expires_at"`
}

// NewMemoryTokenStore 创建内存Token存储
func NewMemoryTokenStore() *MemoryTokenStore {
	store := &MemoryTokenStore{
		tokens: make(map[string]*TokenInfo),
	}
	
	// 启动清理过期token的goroutine
	go store.cleanupExpiredTokens()
	
	return store
}

// GetConnectionConfig 根据token获取连接配置
func (s *MemoryTokenStore) GetConnectionConfig(ctx context.Context, token string) (*ConnectionConfig, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	tokenInfo, exists := s.tokens[token]
	if !exists {
		return nil, fmt.Errorf("token not found")
	}
	
	// 检查token是否过期
	if time.Now().After(tokenInfo.ExpiresAt) {
		return nil, fmt.Errorf("token expired")
	}
	
	return tokenInfo.Config, nil
}

// StoreConnectionConfig 存储token和连接配置的映射
func (s *MemoryTokenStore) StoreConnectionConfig(ctx context.Context, token string, config *ConnectionConfig) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	// 默认token有效期为24小时
	expiresAt := time.Now().Add(24 * time.Hour)
	
	s.tokens[token] = &TokenInfo{
		Config:    config,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
	
	return nil
}

// DeleteToken 删除token
func (s *MemoryTokenStore) DeleteToken(ctx context.Context, token string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	delete(s.tokens, token)
	return nil
}

// IsValidToken 检查token是否有效
func (s *MemoryTokenStore) IsValidToken(ctx context.Context, token string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	tokenInfo, exists := s.tokens[token]
	if !exists {
		return false
	}
	
	// 检查token是否过期
	return time.Now().Before(tokenInfo.ExpiresAt)
}

// GetTokenInfo 获取token信息（用于调试和管理）
func (s *MemoryTokenStore) GetTokenInfo(ctx context.Context, token string) (*TokenInfo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	tokenInfo, exists := s.tokens[token]
	if !exists {
		return nil, fmt.Errorf("token not found")
	}
	
	return tokenInfo, nil
}

// ListTokens 列出所有token（用于管理）
func (s *MemoryTokenStore) ListTokens(ctx context.Context) map[string]*TokenInfo {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	result := make(map[string]*TokenInfo)
	for token, info := range s.tokens {
		result[token] = info
	}
	
	return result
}

// GetStats 获取存储统计信息
func (s *MemoryTokenStore) GetStats() map[string]interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	
	now := time.Now()
	validTokens := 0
	expiredTokens := 0
	
	for _, info := range s.tokens {
		if now.Before(info.ExpiresAt) {
			validTokens++
		} else {
			expiredTokens++
		}
	}
	
	return map[string]interface{}{
		"total_tokens":   len(s.tokens),
		"valid_tokens":   validTokens,
		"expired_tokens": expiredTokens,
	}
}

// cleanupExpiredTokens 清理过期的token
func (s *MemoryTokenStore) cleanupExpiredTokens() {
	ticker := time.NewTicker(1 * time.Hour) // 每小时清理一次
	defer ticker.Stop()
	
	for range ticker.C {
		s.mutex.Lock()
		now := time.Now()
		
		for token, info := range s.tokens {
			if now.After(info.ExpiresAt) {
				delete(s.tokens, token)
			}
		}
		
		s.mutex.Unlock()
	}
}

// SetTokenExpiry 设置token过期时间
func (s *MemoryTokenStore) SetTokenExpiry(ctx context.Context, token string, expiresAt time.Time) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	tokenInfo, exists := s.tokens[token]
	if !exists {
		return fmt.Errorf("token not found")
	}
	
	tokenInfo.ExpiresAt = expiresAt
	return nil
}