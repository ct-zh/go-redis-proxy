package service

import (
	"context"
	"testing"
	"time"

	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSimpleTokenService_NewSimpleTokenService(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	assert.NotNil(t, tokenService)
	assert.Equal(t, connectionService, tokenService.connectionService)
}

func TestSimpleTokenService_CreateToken(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	req := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	resp, err := tokenService.CreateToken(ctx, req)
	if err != nil {
		t.Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)
	assert.NotEmpty(t, resp.ConnID)
	assert.True(t, resp.ExpiresAt.After(time.Now()))
	assert.True(t, resp.ExpiresAt.Before(time.Now().Add(time.Hour+time.Minute)))
}

func TestSimpleTokenService_ValidateToken(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	req := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	// 创建Token
	resp, err := tokenService.CreateToken(ctx, req)
	if err != nil {
		t.Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}

	// 验证有效Token
	valid, err := tokenService.ValidateToken(ctx, resp.Token)
	require.NoError(t, err)
	assert.True(t, valid)

	// 验证无效Token
	valid, err = tokenService.ValidateToken(ctx, "invalid-token")
	require.NoError(t, err)
	assert.False(t, valid)
}

func TestSimpleTokenService_DeleteToken(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	req := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	// 创建Token
	resp, err := tokenService.CreateToken(ctx, req)
	if err != nil {
		t.Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}

	// 验证Token有效
	valid, err := tokenService.ValidateToken(ctx, resp.Token)
	require.NoError(t, err)
	assert.True(t, valid)

	// 删除Token
	deleteReq := &types.DisconnectRequest{
		Token: resp.Token,
	}
	err = tokenService.DeleteToken(ctx, deleteReq)
	require.NoError(t, err)

	// 验证Token已失效
	valid, err = tokenService.ValidateToken(ctx, resp.Token)
	require.NoError(t, err)
	assert.False(t, valid)
}

func TestSimpleTokenService_RefreshToken(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	// 创建初始Token
	createReq := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	createResp, err := tokenService.CreateToken(ctx, createReq)
	if err != nil {
		t.Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}

	// 刷新Token
	refreshReq := &types.TokenRefreshRequest{
		Token:    createResp.Token,
		TokenTTL: 7200,
	}

	refreshResp, err := tokenService.RefreshToken(ctx, refreshReq)
	require.NoError(t, err)
	assert.NotNil(t, refreshResp)
	assert.NotEmpty(t, refreshResp.Token)
	assert.NotEqual(t, createResp.Token, refreshResp.Token, "刷新后应该生成新Token")

	// 验证旧Token已失效
	valid, err := tokenService.ValidateToken(ctx, createResp.Token)
	require.NoError(t, err)
	assert.False(t, valid, "旧Token应该失效")

	// 验证新Token有效
	valid, err = tokenService.ValidateToken(ctx, refreshResp.Token)
	require.NoError(t, err)
	assert.True(t, valid, "新Token应该有效")
}

func TestSimpleTokenService_RefreshInvalidToken(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	// 尝试刷新无效Token
	refreshReq := &types.TokenRefreshRequest{
		Token:    "invalid-token",
		TokenTTL: 7200,
	}

	_, err := tokenService.RefreshToken(ctx, refreshReq)
	assert.Error(t, err, "刷新无效Token应该返回错误")
}

func TestSimpleTokenService_GetConnectionStats(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	stats, err := tokenService.GetConnectionStats(ctx)
	require.NoError(t, err)
	assert.NotNil(t, stats)
	assert.GreaterOrEqual(t, stats.TotalConnections, 0)
	assert.GreaterOrEqual(t, stats.ActiveConnections, 0)
	assert.GreaterOrEqual(t, stats.PoolSize, 0)
	assert.GreaterOrEqual(t, stats.IdleConnections, 0)
}

func TestSimpleTokenService_HealthCheck(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	err := tokenService.HealthCheck(ctx)
	assert.NoError(t, err, "健康检查应该通过")
}

func TestSimpleTokenService_ConcurrentOperations(t *testing.T) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	// 并发创建Token
	tokenChan := make(chan string, 10)
	errChan := make(chan error, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			req := &types.ConnectRequest{
				RedisRequest: types.RedisRequest{
					Addr:     "localhost:6379",
					Password: "",
					DB:       id % 16,
				},
				TokenTTL: 3600,
			}
			
			resp, err := tokenService.CreateToken(ctx, req)
			if err != nil {
				errChan <- err
				return
			}
			tokenChan <- resp.Token
		}(i)
	}

	// 收集结果
	tokens := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		select {
		case token := <-tokenChan:
			tokens = append(tokens, token)
		case err := <-errChan:
			if err.Error() != "dial tcp [::1]:6379: connect: connection refused" &&
				err.Error() != "dial tcp 127.0.0.1:6379: connect: connection refused" {
				t.Errorf("并发创建Token失败: %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Error("并发测试超时")
			return
		}
	}

	// 验证所有Token都有效且唯一
	tokenSet := make(map[string]bool)
	for _, token := range tokens {
		assert.False(t, tokenSet[token], "Token应该唯一: %s", token)
		tokenSet[token] = true

		valid, err := tokenService.ValidateToken(ctx, token)
		require.NoError(t, err)
		assert.True(t, valid, "Token应该有效: %s", token)
	}
}

func BenchmarkSimpleTokenService_CreateToken(b *testing.B) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	req := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := tokenService.CreateToken(ctx, req)
		if err != nil {
			b.Skipf("跳过基准测试，Redis服务器不可用: %v", err)
			return
		}
	}
}

func BenchmarkSimpleTokenService_ValidateToken(b *testing.B) {
	connectionService := connection.NewService()
	defer connectionService.Close()

	tokenService := NewSimpleTokenService(connectionService)
	ctx := context.Background()

	req := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	// 预先创建Token
	resp, err := tokenService.CreateToken(ctx, req)
	if err != nil {
		b.Skipf("跳过基准测试，Redis服务器不可用: %v", err)
		return
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tokenService.ValidateToken(ctx, resp.Token)
	}
}