package errors

// BusinessError 定义业务错误接口
type BusinessError interface {
	error
	Code() int
	Message() string
}

// businessError 实现BusinessError接口
type businessError struct {
	code    int
	message string
}

// Error 实现error接口
func (e *businessError) Error() string {
	return e.message
}

// Code 返回错误码
func (e *businessError) Code() int {
	return e.code
}

// Message 返回错误消息
func (e *businessError) Message() string {
	return e.message
}

// NewBusinessError 创建业务错误实例
func NewBusinessError(code int, message string) BusinessError {
	return &businessError{
		code:    code,
		message: message,
	}
}