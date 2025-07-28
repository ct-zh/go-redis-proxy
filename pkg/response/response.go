package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	apperrors "github.com/ct-zh/go-redis-proxy/pkg/errors"
)

// BaseResponse represents the standard API response structure
type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success sends a successful response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

// SuccessWithMessage sends a successful response with custom message
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Error sends an error response
func Error(c *gin.Context, httpCode int, message string, err error) {
	response := BaseResponse{
		Code:    httpCode,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(httpCode, response)
}

// BadRequest sends a 400 bad request response
func BadRequest(c *gin.Context, message string, err error) {
	Error(c, http.StatusBadRequest, message, err)
}

// Unauthorized sends a 401 unauthorized response
func Unauthorized(c *gin.Context, message string, err error) {
	Error(c, http.StatusUnauthorized, message, err)
}

// Forbidden sends a 403 forbidden response
func Forbidden(c *gin.Context, message string, err error) {
	Error(c, http.StatusForbidden, message, err)
}

// NotFound sends a 404 not found response
func NotFound(c *gin.Context, message string, err error) {
	Error(c, http.StatusNotFound, message, err)
}

// InternalServerError sends a 500 internal server error response
func InternalServerError(c *gin.Context, message string, err error) {
	Error(c, http.StatusInternalServerError, message, err)
}

// HandleAppError handles application errors and sends appropriate HTTP responses
func HandleAppError(c *gin.Context, err error) {
	if appErr, ok := apperrors.IsAppError(err); ok {
		switch appErr.Code {
		case apperrors.ErrCodeInvalidRequest,
			 apperrors.ErrCodeMissingParameter,
			 apperrors.ErrCodeInvalidParameter:
			BadRequest(c, appErr.Message, appErr)
		case apperrors.ErrCodeKeyNotFound:
			NotFound(c, appErr.Message, appErr)
		case apperrors.ErrCodeRedisConnection,
			 apperrors.ErrCodeRedisOperation,
			 apperrors.ErrCodeDatabase,
			 apperrors.ErrCodeInternal:
			InternalServerError(c, appErr.Message, appErr)
		default:
			InternalServerError(c, "Unknown error", appErr)
		}
		return
	}

	// Handle non-application errors
	InternalServerError(c, "Internal server error", err)
}

// JSON handles business data and errors, formats them into standard response
func JSON(c *gin.Context, data interface{}, err error) {
	if err != nil {
		// Try to convert to BusinessError
		if bizErr, ok := err.(interface {
			Code() int
			Message() string
		}); ok {
			c.JSON(http.StatusOK, BaseResponse{
				Code:    bizErr.Code(),
				Message: bizErr.Message(),
				Data:    data,
			})
			return
		}
		
		// Handle standard error
		c.JSON(http.StatusOK, BaseResponse{
			Code:    500,
			Message: err.Error(),
			Data:    data,
		})
		return
	}
	
	// Success response
	c.JSON(http.StatusOK, BaseResponse{
		Code:    200,
		Message: "Success",
		Data:    data,
	})
}

// ValidateRequest validates required fields and sends error response if validation fails
func ValidateRequest(c *gin.Context, conditions ...func() (bool, string)) bool {
	for _, condition := range conditions {
		if valid, message := condition(); !valid {
			BadRequest(c, "Validation failed", apperrors.NewInvalidRequestError(message, nil))
			return false
		}
	}
	return true
}

// RequiredField creates a validation condition for required fields
func RequiredField(value string, fieldName string) func() (bool, string) {
	return func() (bool, string) {
		if value == "" {
			return false, fieldName + " is required"
		}
		return true, ""
	}
}

// RequiredSlice creates a validation condition for required slices
func RequiredSlice(value []string, fieldName string) func() (bool, string) {
	return func() (bool, string) {
		if len(value) == 0 {
			return false, fieldName + " is required and cannot be empty"
		}
		return true, ""
	}
}

// PositiveInt creates a validation condition for positive integers
func PositiveInt(value int64, fieldName string) func() (bool, string) {
	return func() (bool, string) {
		if value <= 0 {
			return false, fieldName + " must be greater than 0"
		}
		return true, ""
	}
}