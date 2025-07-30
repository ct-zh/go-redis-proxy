package integration

import (
	"context"
	"testing"
	"time"

	"github.com/ct-zh/go-redis-proxy/internal/adapter"
	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// IntegrationTestSuite 集成测试套件
type IntegrationTestSuite struct {
	suite.Suite
	connectionService *connection.Service
	stringService     *service.DynamicRedisStringService
	tokenService      *service.SimpleTokenService
	legacyAdapter     *adapter.LegacyCompatibilityAdapter
	ctx               context.Context
}

// SetupSuite 测试套件初始化
func (suite *IntegrationTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	
	// 初始化服务
	suite.connectionService = connection.NewService()
	dynamicDAO := dao.NewDynamicRedisDAO()
	suite.stringService = service.NewDynamicRedisStringService(suite.connectionService, dynamicDAO)
	suite.tokenService = service.NewSimpleTokenService(suite.connectionService)
	suite.legacyAdapter = adapter.NewLegacyCompatibilityAdapter()
}

// TearDownSuite 测试套件清理
func (suite *IntegrationTestSuite) TearDownSuite() {
	if suite.connectionService != nil {
		suite.connectionService.Close()
	}
	if suite.legacyAdapter != nil {
		suite.legacyAdapter.Close()
	}
}

// TestCompleteWorkflow 测试完整工作流程
func (suite *IntegrationTestSuite) TestCompleteWorkflow() {
	// 1. 创建Token
	tokenReq := &types.TokenRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	tokenResp, err := suite.tokenService.CreateToken(suite.ctx, tokenReq)
	if err != nil {
		suite.T().Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}
	
	require.NotNil(suite.T(), tokenResp)
	assert.NotEmpty(suite.T(), tokenResp.Token)

	// 2. 验证Token
	valid, err := suite.tokenService.ValidateToken(suite.ctx, tokenResp.Token)
	require.NoError(suite.T(), err)
	assert.True(suite.T(), valid)

	// 3. 使用字符串服务进行Redis操作
	setReq := &types.StringSetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key:   "integration_test_key",
		Value: "integration_test_value",
		TTL:   300,
	}

	setResp, err := suite.stringService.Set(suite.ctx, setReq)
	require.NoError(suite.T(), err)
	assert.NotNil(suite.T(), setResp)

	// 4. 获取值验证
	getReq := &types.StringGetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key: "integration_test_key",
	}

	getResp, err := suite.stringService.Get(suite.ctx, getReq)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), "integration_test_value", getResp.Value)

	// 5. 刷新Token
	refreshReq := &types.TokenRefreshRequest{
		Token:        tokenResp.Token,
		RedisRequest: tokenReq.RedisRequest,
		TokenTTL:     7200,
	}

	newTokenResp, err := suite.tokenService.RefreshToken(suite.ctx, refreshReq)
	require.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), tokenResp.Token, newTokenResp.Token)

	// 6. 验证新Token有效，旧Token无效
	valid, err = suite.tokenService.ValidateToken(suite.ctx, newTokenResp.Token)
	require.NoError(suite.T(), err)
	assert.True(suite.T(), valid)

	valid, err = suite.tokenService.ValidateToken(suite.ctx, tokenResp.Token)
	require.NoError(suite.T(), err)
	assert.False(suite.T(), valid)

	// 7. 清理资源
	err = suite.tokenService.DeleteToken(suite.ctx, newTokenResp.Token)
	require.NoError(suite.T(), err)

	// 删除测试数据
	delReq := &types.StringDelRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key: "integration_test_key",
	}
	_, err = suite.stringService.Del(suite.ctx, delReq)
	require.NoError(suite.T(), err)
}

// TestLegacyCompatibility 测试向后兼容性
func (suite *IntegrationTestSuite) TestLegacyCompatibility() {
	// 1. 使用传统接口创建连接
	connectReq := &types.ConnectRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		TokenTTL: 3600,
	}

	connectResp, err := suite.legacyAdapter.Connect(suite.ctx, connectReq)
	if err != nil {
		suite.T().Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}
	
	require.NotNil(suite.T(), connectResp)
	assert.NotEmpty(suite.T(), connectResp.Token)

	// 2. 使用传统接口进行Redis操作
	setReq := &types.StringSetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key:   "legacy_test_key",
		Value: "legacy_test_value",
		TTL:   300,
	}

	setResp, err := suite.legacyAdapter.StringSet(suite.ctx, setReq)
	require.NoError(suite.T(), err)
	assert.NotNil(suite.T(), setResp)

	getReq := &types.StringGetRequest{
		RedisRequest: types.RedisRequest{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		Key: "legacy_test_key",
	}

	getResp, err := suite.legacyAdapter.StringGet(suite.ctx, getReq)
	require.NoError(suite.T(), err)
	assert.Equal(suite.T(), "legacy_test_value", getResp.Value)

	// 3. 验证Token
	valid, err := suite.legacyAdapter.ValidateToken(suite.ctx, connectResp.Token)
	require.NoError(suite.T(), err)
	assert.True(suite.T(), valid)

	// 4. 断开连接
	disconnectReq := &types.DisconnectRequest{
		Token: connectResp.Token,
	}

	err = suite.legacyAdapter.Disconnect(suite.ctx, disconnectReq)
	require.NoError(suite.T(), err)

	// 5. 验证Token已失效
	valid, err = suite.legacyAdapter.ValidateToken(suite.ctx, connectResp.Token)
	require.NoError(suite.T(), err)
	assert.False(suite.T(), valid)
}

// TestBatchOperations 测试批量操作
func (suite *IntegrationTestSuite) TestBatchOperations() {
	batchOps := &adapter.BatchStringOperations{
		Sets: []*types.StringSetRequest{
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_1",
				Value:        "batch_value_1",
				TTL:          300,
			},
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_2",
				Value:        "batch_value_2",
				TTL:          300,
			},
		},
		Gets: []*types.StringGetRequest{
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_1",
			},
			{
				RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
				Key:          "batch_key_2",
			},
		},
	}

	results := suite.legacyAdapter.ExecuteBatchStringOperations(suite.ctx, batchOps)
	
	if len(results.Errors) > 0 {
		// 检查是否是Redis连接错误
		for _, err := range results.Errors {
			if err.Error() == "dial tcp [::1]:6379: connect: connection refused" ||
				err.Error() == "dial tcp 127.0.0.1:6379: connect: connection refused" {
				suite.T().Skipf("跳过测试，Redis服务器不可用: %v", err)
				return
			}
		}
		suite.T().Errorf("批量操作出现错误: %v", results.Errors)
	}

	assert.Equal(suite.T(), 2, len(results.SetResults), "应该有2个SET结果")
	assert.Equal(suite.T(), 2, len(results.GetResults), "应该有2个GET结果")
}

// TestConnectionReuse 测试连接复用
func (suite *IntegrationTestSuite) TestConnectionReuse() {
	req := types.RedisRequest{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	// 获取第一个连接
	conn1, err := suite.connectionService.GetConnection(suite.ctx, req)
	if err != nil {
		suite.T().Skipf("跳过测试，Redis服务器不可用: %v", err)
		return
	}

	// 获取第二个连接
	conn2, err := suite.connectionService.GetConnection(suite.ctx, req)
	require.NoError(suite.T(), err)

	// 验证连接复用
	assert.Equal(suite.T(), conn1, conn2, "相同配置应该复用连接")

	// 获取连接统计
	stats := suite.connectionService.GetStats()
	assert.GreaterOrEqual(suite.T(), stats.ActiveConnections, 1)
	assert.GreaterOrEqual(suite.T(), stats.TotalConnections, 1)
}

// TestPerformanceComparison 测试性能对比
func (suite *IntegrationTestSuite) TestPerformanceComparison() {
	iterations := 100
	
	// 测试新接口性能
	start := time.Now()
	for i := 0; i < iterations; i++ {
		setReq := &types.StringSetRequest{
			RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
			Key:          "perf_new_key",
			Value:        "perf_new_value",
			TTL:          300,
		}
		
		_, err := suite.stringService.Set(suite.ctx, setReq)
		if err != nil {
			suite.T().Skipf("跳过性能测试，Redis服务器不可用: %v", err)
			return
		}
	}
	newDuration := time.Since(start)

	// 测试传统接口性能
	start = time.Now()
	for i := 0; i < iterations; i++ {
		setReq := &types.StringSetRequest{
			RedisRequest: types.RedisRequest{Addr: "localhost:6379"},
			Key:          "perf_legacy_key",
			Value:        "perf_legacy_value",
			TTL:          300,
		}
		
		_, err := suite.legacyAdapter.StringSet(suite.ctx, setReq)
		if err != nil {
			suite.T().Skipf("跳过性能测试，Redis服务器不可用: %v", err)
			return
		}
	}
	legacyDuration := time.Since(start)

	suite.T().Logf("新接口性能: %v", newDuration)
	suite.T().Logf("传统接口性能: %v", legacyDuration)
	
	// 新接口应该不比传统接口慢太多（允许10%的性能损失）
	assert.True(suite.T(), newDuration <= legacyDuration*11/10, 
		"新接口性能不应该比传统接口慢太多")
}

// TestHealthAndStats 测试健康检查和统计
func (suite *IntegrationTestSuite) TestHealthAndStats() {
	// 健康检查
	err := suite.connectionService.HealthCheck(suite.ctx)
	assert.NoError(suite.T(), err)

	err = suite.tokenService.HealthCheck(suite.ctx)
	assert.NoError(suite.T(), err)

	err = suite.legacyAdapter.HealthCheck(suite.ctx)
	assert.NoError(suite.T(), err)

	// 统计信息
	stats, err := suite.tokenService.GetConnectionStats(suite.ctx)
	require.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), stats.ActiveConnections, 0)
	assert.GreaterOrEqual(suite.T(), stats.TotalConnections, 0)
	assert.NotNil(suite.T(), stats.ConnectionsByAddr)

	legacyStats, err := suite.legacyAdapter.GetConnectionStats(suite.ctx)
	require.NoError(suite.T(), err)
	assert.GreaterOrEqual(suite.T(), legacyStats.ActiveConnections, 0)
}

// TestConcurrentAccess 测试并发访问
func (suite *IntegrationTestSuite) TestConcurrentAccess() {
	concurrency := 10
	done := make(chan bool, concurrency)
	errors := make(chan error, concurrency)

	for i := 0; i < concurrency; i++ {
		go func(id int) {
			defer func() { done <- true }()

			// 创建Token
			tokenReq := &types.TokenRequest{
				RedisRequest: types.RedisRequest{
					Addr: "localhost:6379",
					DB:   0,
				},
				TokenTTL: 3600,
			}

			tokenResp, err := suite.tokenService.CreateToken(suite.ctx, tokenReq)
			if err != nil {
				errors <- err
				return
			}

			// 进行Redis操作
			setReq := &types.StringSetRequest{
				RedisRequest: types.RedisRequest{
					Addr: "localhost:6379",
					DB:   0,
				},
				Key:   "concurrent_key",
				Value: "concurrent_value",
				TTL:   300,
			}

			_, err = suite.stringService.Set(suite.ctx, setReq)
			if err != nil {
				errors <- err
				return
			}

			// 清理Token
			err = suite.tokenService.DeleteToken(suite.ctx, tokenResp.Token)
			if err != nil {
				errors <- err
				return
			}
		}(i)
	}

	// 等待所有goroutine完成
	for i := 0; i < concurrency; i++ {
		select {
		case <-done:
			// 正常完成
		case err := <-errors:
			if err.Error() != "dial tcp [::1]:6379: connect: connection refused" &&
				err.Error() != "dial tcp 127.0.0.1:6379: connect: connection refused" {
				suite.T().Errorf("并发测试失败: %v", err)
			}
		case <-time.After(10 * time.Second):
			suite.T().Error("并发测试超时")
			return
		}
	}
}

// TestIntegrationSuite 运行集成测试套件
func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}