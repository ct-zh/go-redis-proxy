package errors

import (
	"fmt"
	"sync"
)

// ErrorInfo 错误信息结构
type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Module  string `json:"module"`
}

// ErrorRegistry 错误注册表
type ErrorRegistry struct {
	mu     sync.RWMutex
	errors map[int]*ErrorInfo
}

// NewErrorRegistry 创建错误注册表实例
func NewErrorRegistry() *ErrorRegistry {
	return &ErrorRegistry{
		errors: make(map[int]*ErrorInfo),
	}
}

// Register 注册错误信息
func (r *ErrorRegistry) Register(code int, message, module string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, exists := r.errors[code]; exists {
		return fmt.Errorf("error code %d already registered in module %s", code, r.errors[code].Module)
	}
	
	r.errors[code] = &ErrorInfo{
		Code:    code,
		Message: message,
		Module:  module,
	}
	
	return nil
}

// Get 获取错误信息
func (r *ErrorRegistry) Get(code int) (*ErrorInfo, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	info, exists := r.errors[code]
	return info, exists
}

// GetAll 获取所有错误信息
func (r *ErrorRegistry) GetAll() map[int]*ErrorInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	result := make(map[int]*ErrorInfo)
	for k, v := range r.errors {
		result[k] = v
	}
	return result
}

// Validate 验证错误注册表的完整性
func (r *ErrorRegistry) Validate() error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	if len(r.errors) == 0 {
		return fmt.Errorf("no error codes registered")
	}
	
	// 检查是否有重复的错误码（这个在Register时已经检查了，这里是双重保险）
	for code, info := range r.errors {
		if info.Code != code {
			return fmt.Errorf("inconsistent error code: registered %d but info shows %d", code, info.Code)
		}
	}
	
	return nil
}