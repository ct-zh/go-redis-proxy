package errors

import (
	"fmt"
	"sync"
)

// ErrorManager 错误管理器
type ErrorManager struct {
	registry *ErrorRegistry
}

var (
	globalManager *ErrorManager
	once          sync.Once
)

// GetGlobalManager 获取全局错误管理器实例（单例模式）
func GetGlobalManager() *ErrorManager {
	once.Do(func() {
		globalManager = &ErrorManager{
			registry: NewErrorRegistry(),
		}
		// 初始化预定义错误码
		globalManager.initPredefinedErrors()
	})
	return globalManager
}

// NewErrorManager 创建错误管理器实例
func NewErrorManager() *ErrorManager {
	manager := &ErrorManager{
		registry: NewErrorRegistry(),
	}
	manager.initPredefinedErrors()
	return manager
}

// initPredefinedErrors 初始化预定义的错误码
func (m *ErrorManager) initPredefinedErrors() {
	// 系统级错误
	m.registry.Register(CodeInternalError, "内部服务器错误", "system")
	m.registry.Register(CodeInvalidParams, "参数验证失败", "system")
	m.registry.Register(CodeMethodNotAllowed, "方法不允许", "system")
	m.registry.Register(CodeUnauthorized, "未授权访问", "system")
	
	// Redis连接错误
	m.registry.Register(CodeRedisConnectFailed, "Redis连接失败", "redis")
	m.registry.Register(CodeRedisTimeout, "Redis操作超时", "redis")
	m.registry.Register(CodeRedisAuthFailed, "Redis认证失败", "redis")
	m.registry.Register(CodeRedisDBSelectFailed, "Redis数据库选择失败", "redis")
	
	// String操作错误
	m.registry.Register(CodeStringKeyNotFound, "键不存在", "string")
	m.registry.Register(CodeStringTypeMismatch, "类型不匹配", "string")
	m.registry.Register(CodeStringSetFailed, "设置失败", "string")
	m.registry.Register(CodeStringGetFailed, "获取失败", "string")
	m.registry.Register(CodeStringDelFailed, "删除失败", "string")
	m.registry.Register(CodeStringIncrFailed, "自增失败", "string")
	m.registry.Register(CodeStringDecrFailed, "自减失败", "string")
	m.registry.Register(CodeStringExpireFailed, "设置过期时间失败", "string")
	
	// List操作错误
	m.registry.Register(CodeListIndexOutOfRange, "索引超出范围", "list")
	m.registry.Register(CodeListTypeMismatch, "类型不匹配", "list")
	m.registry.Register(CodeListPushFailed, "推入失败", "list")
	m.registry.Register(CodeListPopFailed, "弹出失败", "list")
	m.registry.Register(CodeListRemoveFailed, "删除失败", "list")
	m.registry.Register(CodeListTrimFailed, "裁剪失败", "list")
	
	// Hash操作错误
	m.registry.Register(CodeHashKeyNotFound, "Hash键不存在", "hash")
	m.registry.Register(CodeHashFieldNotFound, "Hash字段不存在", "hash")
	m.registry.Register(CodeHashTypeMismatch, "类型不匹配", "hash")
	m.registry.Register(CodeHashSetFailed, "设置失败", "hash")
	m.registry.Register(CodeHashGetFailed, "获取失败", "hash")
	m.registry.Register(CodeHashDelFailed, "删除失败", "hash")
	m.registry.Register(CodeHashDeleteFailed, "删除失败", "hash")
	m.registry.Register(CodeHashIncrementFailed, "增量操作失败", "hash")
	
	// Set操作错误
	m.registry.Register(CodeSetTypeMismatch, "类型不匹配", "set")
	m.registry.Register(CodeSetAddFailed, "添加失败", "set")
	m.registry.Register(CodeSetRemoveFailed, "删除失败", "set")
	m.registry.Register(CodeSetMemberNotFound, "成员不存在", "set")
	
	// ZSet操作错误
	m.registry.Register(CodeZSetTypeMismatch, "类型不匹配", "zset")
	m.registry.Register(CodeZSetAddFailed, "添加失败", "zset")
	m.registry.Register(CodeZSetRemoveFailed, "删除失败", "zset")
	m.registry.Register(CodeZSetMemberNotFound, "成员不存在", "zset")
	m.registry.Register(CodeZSetRankNotFound, "排名不存在", "zset")
}

// NewBusinessError 创建业务错误
func (m *ErrorManager) NewBusinessError(code int, args ...interface{}) BusinessError {
	info, exists := m.registry.Get(code)
	if !exists {
		// 如果错误码未注册，返回通用错误
		return NewBusinessError(CodeInternalError, fmt.Sprintf("未知错误码: %d", code))
	}
	
	message := info.Message
	if len(args) > 0 {
		message = fmt.Sprintf(info.Message, args...)
	}
	
	return NewBusinessError(code, message)
}

// Register 注册自定义错误码
func (m *ErrorManager) Register(code int, message, module string) error {
	return m.registry.Register(code, message, module)
}

// Validate 验证错误管理器
func (m *ErrorManager) Validate() error {
	return m.registry.Validate()
}

// GetAllErrors 获取所有错误信息
func (m *ErrorManager) GetAllErrors() map[int]*ErrorInfo {
	return m.registry.GetAll()
}

// 全局便捷方法

// NewError 创建业务错误（使用全局管理器）
func NewError(code int, args ...interface{}) BusinessError {
	return GetGlobalManager().NewBusinessError(code, args...)
}

// RegisterError 注册错误码（使用全局管理器）
func RegisterError(code int, message, module string) error {
	return GetGlobalManager().Register(code, message, module)
}

// ValidateRegistry 验证错误注册表（使用全局管理器）
func ValidateRegistry() error {
	return GetGlobalManager().Validate()
}