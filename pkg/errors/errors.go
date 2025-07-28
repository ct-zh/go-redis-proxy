package errors

import (
	"errors"
	"fmt"
)

// ErrorCode represents different types of errors
type ErrorCode int

const (
	// Internal errors
	ErrCodeInternal ErrorCode = iota + 1000
	ErrCodeDatabase
	ErrCodeRedisConnection
	ErrCodeRedisOperation
	
	// Client errors
	ErrCodeInvalidRequest ErrorCode = iota + 2000
	ErrCodeMissingParameter
	ErrCodeInvalidParameter
	ErrCodeKeyNotFound
)

// AppError represents a structured application error
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Details string    `json:"details,omitempty"`
	Cause   error     `json:"-"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("[%d] %s: %s (%v)", e.Code, e.Message, e.Details, e.Cause)
	}
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Details)
}

// Unwrap returns the underlying cause
func (e *AppError) Unwrap() error {
	return e.Cause
}

// NewAppError creates a new application error
func NewAppError(code ErrorCode, message string, details string, cause error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Details: details,
		Cause:   cause,
	}
}

// Predefined error constructors

// NewInternalError creates a new internal server error
func NewInternalError(details string, cause error) *AppError {
	return NewAppError(ErrCodeInternal, "Internal server error", details, cause)
}

// NewDatabaseError creates a new database error
func NewDatabaseError(details string, cause error) *AppError {
	return NewAppError(ErrCodeDatabase, "Database error", details, cause)
}

// NewRedisConnectionError creates a new Redis connection error
func NewRedisConnectionError(details string, cause error) *AppError {
	return NewAppError(ErrCodeRedisConnection, "Redis connection error", details, cause)
}

// NewRedisOperationError creates a new Redis operation error
func NewRedisOperationError(details string, cause error) *AppError {
	return NewAppError(ErrCodeRedisOperation, "Redis operation error", details, cause)
}

// NewInvalidRequestError creates a new invalid request error
func NewInvalidRequestError(details string, cause error) *AppError {
	return NewAppError(ErrCodeInvalidRequest, "Invalid request", details, cause)
}

// NewMissingParameterError creates a new missing parameter error
func NewMissingParameterError(parameter string) *AppError {
	return NewAppError(ErrCodeMissingParameter, "Missing required parameter", parameter, nil)
}

// NewInvalidParameterError creates a new invalid parameter error
func NewInvalidParameterError(parameter string, details string) *AppError {
	return NewAppError(ErrCodeInvalidParameter, "Invalid parameter", fmt.Sprintf("%s: %s", parameter, details), nil)
}

// NewKeyNotFoundError creates a new key not found error
func NewKeyNotFoundError(key string) *AppError {
	return NewAppError(ErrCodeKeyNotFound, "Key not found", key, nil)
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

// GetErrorCode extracts the error code from an error
func GetErrorCode(err error) ErrorCode {
	if appErr, ok := IsAppError(err); ok {
		return appErr.Code
	}
	return ErrCodeInternal
}