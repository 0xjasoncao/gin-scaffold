package errorsx

import (
	"fmt"
	"gin-scaffold/pkg/core/validatorx"
	"net/http"

	errors2 "github.com/pkg/errors"
)

// Define alias

var (
	WithStack = errors2.WithStack
	Errorf    = errors2.Errorf
	Is        = errors2.Is
	As        = errors2.As
	Wrap      = errors2.Wrap
	New       = errors2.New
)

// ResponseError represents a structured error that includes both business and HTTP status information
// It supports error chaining through the Err field and provides chainable methods for modification
type ResponseError struct {
	Code       int    `json:"code,omitempty"`       // Business-specific error code for client-side handling
	StatusCode int    `json:"statusCode,omitempty"` // HTTP status code following standard HTTP semantics
	Message    string `json:"message,omitempty"`    // Human-readable error message (client-friendly)
	Err        error  `json:"-"`                    // Original underlying error (for server-side debugging only)
}

// Error implements the error interface
// Returns the original error message if available, otherwise returns the client message
func (e *ResponseError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// UnwrapResponseError extracts a *ResponseError from a generic error using errorsx.As
// Returns nil if the error is not a *ResponseError or wraps one
func UnwrapResponseError(err error) *ResponseError {
	var v *ResponseError
	if As(err, &v) {
		return v
	}
	return nil
}

// NewResponseError creates a new ResponseError with specified code, status code, and formatted message
func NewResponseError(code, statusCode int, message string, args ...interface{}) *ResponseError {
	return &ResponseError{
		Code:       code,
		StatusCode: statusCode,
		Message:    fmt.Sprintf(message, args...),
	}
}

// Predefined error constructors for common error scenarios
// These provide consistent error codes while allowing custom messages

// NewInternal creates a 500-level error for server-side issues
func NewInternal(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusInternalServerError, http.StatusInternalServerError, message, args...)
}

// NewNotFound creates a 404-level error for missing resources
func NewNotFound(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusNotFound, http.StatusNotFound, message, args...)
}

// NewBadRequest creates a 400-level error for invalid client input
func NewBadRequest(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusBadRequest, http.StatusBadRequest, message, args...)
}

// NewTooManyRequest creates a 429-level error for rate limiting
func NewTooManyRequest(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusTooManyRequests, http.StatusTooManyRequests, message, args...)
}

// NewRequestEntityTooLarge creates a 429-level error for rate limiting
func NewRequestEntityTooLarge(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusRequestEntityTooLarge, http.StatusRequestEntityTooLarge, message, args...)
}

// NewUnauthorized creates a 401-level error for authentication failures
func NewUnauthorized(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusUnauthorized, http.StatusUnauthorized, message, args...)
}

// NewMethodNotAllowed creates a 405-level error for unsupported HTTP methods
func NewMethodNotAllowed(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusMethodNotAllowed, http.StatusMethodNotAllowed, message, args...)
}

func NewForbidden(message string, args ...interface{}) *ResponseError {
	return NewResponseError(http.StatusForbidden, http.StatusForbidden, message, args...)

}

func NewInvalidParams(err any) *ResponseError {
	switch msg := err.(type) {
	case error:
		return NewBadRequest(validatorx.ZhTranslate(msg))
	case string:
		return NewBadRequest(msg)
	default:
		return NewBadRequest("%v", msg)
	}
}

// Chainable modification methods
// These allow fluent-style modification of error properties

// WithError attaches an underlying error to this ResponseError
// Returns the modified ResponseError for method chaining
func (e *ResponseError) WithError(err error) *ResponseError {
	e.Err = err
	return e
}

// WithMessage updates the error message with a formatted string
// Returns the modified ResponseError for method chaining
func (e *ResponseError) WithMessage(message string, args ...interface{}) *ResponseError {
	e.Message = fmt.Sprintf(message, args...)
	return e
}

// WithCode updates the business error code
// Returns the modified ResponseError for method chaining
func (e *ResponseError) WithCode(code int) *ResponseError {
	e.Code = code
	return e
}
