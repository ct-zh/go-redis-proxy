package connection

import (
	"context"
	"sync"
	"testing"

	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestService_NewService(t *testing.T) {
	service := NewService()
	assert.NotNil(t, service)
	assert.NotNil(t, service.resolver)
	assert.NotNil(t, service.manager)
	
	defer service.Close()
}

func TestService_GetConnection(t *testing.T) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	
	// 使用包含RedisRequest字段的结构体
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	// 由于没有实际的Redis服务器，这个测试会失败
	// 但我们可以验证解析逻辑
	_, err := service.GetConnection(ctx, req)
	if err != nil {
		t.Logf("跳过测试，Redis服务器不可用: %v", err)
		t.Skip("Redis服务器不可用")
	}
}

func TestService_CreateToken(t *testing.T) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	token, err := service.CreateToken(ctx, req)
	if err != nil {
		t.Logf("跳过测试，Redis服务器不可用: %v", err)
		t.Skip("Redis服务器不可用")
	}

	assert.NotEmpty(t, token)
	assert.True(t, service.IsValidToken(ctx, token))
}

func TestService_TokenOperations(t *testing.T) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	// 创建token
	token, err := service.CreateToken(ctx, req)
	if err != nil {
		t.Logf("跳过测试，Redis服务器不可用: %v", err)
		t.Skip("Redis服务器不可用")
	}

	// 验证token
	assert.True(t, service.IsValidToken(ctx, token))

	// 删除token
	err = service.DeleteToken(ctx, token)
	assert.NoError(t, err)

	// 验证token已失效
	assert.False(t, service.IsValidToken(ctx, token))
}

func TestService_InvalidToken(t *testing.T) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	
	// 测试无效token
	assert.False(t, service.IsValidToken(ctx, "invalid_token"))
}

func TestService_GetStats(t *testing.T) {
	service := NewService()
	defer service.Close()

	stats := service.GetStats()
	assert.NotNil(t, stats)
	assert.Contains(t, stats, "connection_manager")
	assert.Contains(t, stats, "token_store")
}

func TestService_HealthCheck(t *testing.T) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	health := service.HealthCheck(ctx)
	assert.NotNil(t, health)
	assert.Contains(t, health, "connections")
	assert.Contains(t, health, "token_store")
}

func TestService_ConcurrentAccess(t *testing.T) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	var wg sync.WaitGroup
	tokens := make([]string, 10)
	errors := make([]error, 10)

	// 并发创建token
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			token, err := service.CreateToken(ctx, req)
			if err != nil {
				t.Logf("并发创建Token失败: %v", err)
				errors[index] = err
				return
			}
			tokens[index] = token
		}(i)
	}

	wg.Wait()

	// 检查结果
	successCount := 0
	for i, token := range tokens {
		if errors[i] == nil && token != "" {
			successCount++
			assert.True(t, service.IsValidToken(ctx, token))
		}
	}

	t.Logf("成功创建 %d 个Token", successCount)
}

func TestService_ConnectionKey(t *testing.T) {
	service := NewService()
	defer service.Close()

	// 测试连接键生成逻辑
	config1 := &ConnectionConfig{Host: "localhost", Port: 6379, DB: 0}
	config2 := &ConnectionConfig{Host: "localhost", Port: 6379, DB: 1}
	config3 := &ConnectionConfig{Host: "localhost", Port: 6379, DB: 0}

	// 相同配置应该生成相同的键
	key1 := config1.Hash()
	key3 := config3.Hash()
	assert.Equal(t, key1, key3)

	// 不同配置应该生成不同的键
	key2 := config2.Hash()
	assert.NotEqual(t, key1, key2)
}

// 基准测试
func BenchmarkService_GetConnection(b *testing.B) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.GetConnection(ctx, req)
	}
}

func BenchmarkService_CreateToken(b *testing.B) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = service.CreateToken(ctx, req)
	}
}

func BenchmarkService_ValidateToken(b *testing.B) {
	service := NewService()
	defer service.Close()

	ctx := context.Background()
	req := struct {
		types.RedisRequest
	}{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	}

	// 预先创建一个token
	token, _ := service.CreateToken(ctx, req)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = service.IsValidToken(ctx, token)
	}
}