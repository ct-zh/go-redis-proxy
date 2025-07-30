package adapter

import (
	"context"

	"github.com/ct-zh/go-redis-proxy/internal/connection"
	"github.com/ct-zh/go-redis-proxy/internal/dao"
	"github.com/ct-zh/go-redis-proxy/internal/service"
	"github.com/ct-zh/go-redis-proxy/pkg/types"
)

// LegacyCompatibilityAdapter 向后兼容性适配器
// 提供与旧版本API兼容的接口，同时使用新的连接管理系统
type LegacyCompatibilityAdapter struct {
	connectionService *connection.Service
	dynamicDAO        dao.DynamicRedisDAO
	stringService     *service.DynamicRedisStringService
	tokenService      *service.SimpleTokenService
}

// NewLegacyCompatibilityAdapter 创建向后兼容性适配器
func NewLegacyCompatibilityAdapter() *LegacyCompatibilityAdapter {
	// 初始化新的连接管理系统
	connectionService := connection.NewService()
	dynamicDAO := dao.NewDynamicRedisDAO()
	stringService := service.NewDynamicRedisStringService(connectionService, dynamicDAO)
	tokenService := service.NewSimpleTokenService(connectionService)

	return &LegacyCompatibilityAdapter{
		connectionService: connectionService,
		dynamicDAO:        dynamicDAO,
		stringService:     stringService,
		tokenService:      tokenService,
	}
}

// === 兼容旧版本的Redis操作接口 ===

// StringGet 兼容旧版本的字符串GET操作
func (a *LegacyCompatibilityAdapter) StringGet(ctx context.Context, req *types.StringGetRequest) (*types.StringGetData, error) {
	return a.stringService.Get(ctx, req)
}

// StringSet 兼容旧版本的字符串SET操作
func (a *LegacyCompatibilityAdapter) StringSet(ctx context.Context, req *types.StringSetRequest) (*types.StringSetData, error) {
	return a.stringService.Set(ctx, req)
}

// StringDel 兼容旧版本的字符串DEL操作
func (a *LegacyCompatibilityAdapter) StringDel(ctx context.Context, req *types.StringDelRequest) (*types.StringDelData, error) {
	return a.stringService.Del(ctx, req)
}

// StringExists 兼容旧版本的字符串EXISTS操作
func (a *LegacyCompatibilityAdapter) StringExists(ctx context.Context, req *types.StringExistsRequest) (*types.StringExistsData, error) {
	return a.stringService.Exists(ctx, req)
}

// StringIncr 兼容旧版本的字符串INCR操作
func (a *LegacyCompatibilityAdapter) StringIncr(ctx context.Context, req *types.StringIncrRequest) (*types.StringIncrData, error) {
	return a.stringService.Incr(ctx, req)
}

// StringDecr 兼容旧版本的字符串DECR操作
func (a *LegacyCompatibilityAdapter) StringDecr(ctx context.Context, req *types.StringDecrRequest) (*types.StringDecrData, error) {
	return a.stringService.Decr(ctx, req)
}

// StringExpire 兼容旧版本的字符串EXPIRE操作
func (a *LegacyCompatibilityAdapter) StringExpire(ctx context.Context, req *types.StringExpireRequest) (*types.StringExpireData, error) {
	return a.stringService.Expire(ctx, req)
}

// === 兼容旧版本的Token管理接口 ===

// Connect 兼容旧版本的连接建立接口
func (a *LegacyCompatibilityAdapter) Connect(ctx context.Context, req *types.ConnectRequest) (*types.ConnectResponse, error) {
	return a.tokenService.CreateToken(ctx, req)
}

// Disconnect 兼容旧版本的断开连接接口
func (a *LegacyCompatibilityAdapter) Disconnect(ctx context.Context, req *types.DisconnectRequest) error {
	return a.tokenService.DeleteToken(ctx, req)
}

// RefreshToken 兼容旧版本的Token刷新接口
func (a *LegacyCompatibilityAdapter) RefreshToken(ctx context.Context, req *types.TokenRefreshRequest) (*types.ConnectResponse, error) {
	return a.tokenService.RefreshToken(ctx, req)
}

// ValidateToken 兼容旧版本的Token验证接口
func (a *LegacyCompatibilityAdapter) ValidateToken(ctx context.Context, token string) (bool, error) {
	return a.tokenService.ValidateToken(ctx, token)
}

// === 兼容旧版本的统计和监控接口 ===

// GetConnectionStats 兼容旧版本的连接统计接口
func (a *LegacyCompatibilityAdapter) GetConnectionStats(ctx context.Context) (*types.ConnectionStats, error) {
	return a.tokenService.GetConnectionStats(ctx)
}

// HealthCheck 兼容旧版本的健康检查接口
func (a *LegacyCompatibilityAdapter) HealthCheck(ctx context.Context) error {
	return a.tokenService.HealthCheck(ctx)
}

// === 资源管理 ===

// Close 关闭适配器并清理资源
func (a *LegacyCompatibilityAdapter) Close() error {
	return a.connectionService.Close()
}

// === 新功能的便捷访问方法 ===

// GetConnectionService 获取底层连接服务（用于高级用法）
func (a *LegacyCompatibilityAdapter) GetConnectionService() *connection.Service {
	return a.connectionService
}

// GetStringService 获取字符串服务（用于高级用法）
func (a *LegacyCompatibilityAdapter) GetStringService() *service.DynamicRedisStringService {
	return a.stringService
}

// GetTokenService 获取Token服务（用于高级用法）
func (a *LegacyCompatibilityAdapter) GetTokenService() *service.SimpleTokenService {
	return a.tokenService
}

// === 批量操作支持（新功能） ===

// BatchStringOperations 批量字符串操作
type BatchStringOperations struct {
	Gets    []*types.StringGetRequest
	Sets    []*types.StringSetRequest
	Dels    []*types.StringDelRequest
	Expires []*types.StringExpireRequest
}

// BatchStringResults 批量字符串操作结果
type BatchStringResults struct {
	GetResults    []*types.StringGetData
	SetResults    []*types.StringSetData
	DelResults    []*types.StringDelData
	ExpireResults []*types.StringExpireData
	Errors        []error
}

// ExecuteBatchStringOperations 执行批量字符串操作
func (a *LegacyCompatibilityAdapter) ExecuteBatchStringOperations(ctx context.Context, operations *BatchStringOperations) *BatchStringResults {
	results := &BatchStringResults{
		GetResults:    make([]*types.StringGetData, len(operations.Gets)),
		SetResults:    make([]*types.StringSetData, len(operations.Sets)),
		DelResults:    make([]*types.StringDelData, len(operations.Dels)),
		ExpireResults: make([]*types.StringExpireData, len(operations.Expires)),
		Errors:        make([]error, 0),
	}

	// 执行GET操作
	for i, req := range operations.Gets {
		result, err := a.StringGet(ctx, req)
		if err != nil {
			results.Errors = append(results.Errors, err)
		} else {
			results.GetResults[i] = result
		}
	}

	// 执行SET操作
	for i, req := range operations.Sets {
		result, err := a.StringSet(ctx, req)
		if err != nil {
			results.Errors = append(results.Errors, err)
		} else {
			results.SetResults[i] = result
		}
	}

	// 执行DEL操作
	for i, req := range operations.Dels {
		result, err := a.StringDel(ctx, req)
		if err != nil {
			results.Errors = append(results.Errors, err)
		} else {
			results.DelResults[i] = result
		}
	}

	// 执行EXPIRE操作
	for i, req := range operations.Expires {
		result, err := a.StringExpire(ctx, req)
		if err != nil {
			results.Errors = append(results.Errors, err)
		} else {
			results.ExpireResults[i] = result
		}
	}

	return results
}