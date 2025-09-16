package errors

import (
	"fmt"
	errors2 "github.com/pkg/errors"
	"net/http"
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
	Code       int    `json:"code"`                 // Business-specific error code for client-side handling
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

// UnwrapResponseError extracts a *ResponseError from a generic error using errors.As
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

// WrapResponseError wraps an existing error with business and HTTP status information
// Creates a new ResponseError that preserves the original error context
func WrapResponseError(err error, code, statusCode int, message string, args ...interface{}) *ResponseError {
	if err == nil {
		return nil
	}
	return &ResponseError{
		Code:       code,
		StatusCode: statusCode,
		Message:    fmt.Sprintf(message, args...),
		Err:        err,
	}
}

// Predefined error constructors for common error scenarios
// These provide consistent error codes while allowing custom messages

// NewInternalError creates a 500-level error for server-side issues
func NewInternalError(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeInternalError, http.StatusInternalServerError, message, args...)
}

// NewNotFound creates a 404-level error for missing resources
func NewNotFound(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeNotFound, http.StatusNotFound, message, args...)
}

// NewBadRequest creates a 400-level error for invalid client input
func NewBadRequest(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeBadRequest, http.StatusBadRequest, message, args...)
}

// NewTooManyRequest creates a 429-level error for rate limiting
func NewTooManyRequest(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeTooManyRequests, http.StatusTooManyRequests, message, args...)
}

// NewRequestEntityTooLarge creates a 429-level error for rate limiting
func NewRequestEntityTooLarge(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeRequestEntityTooLarge, http.StatusRequestEntityTooLarge, message, args...)
}

// NewUnauthorized creates a 401-level error for authentication failures
func NewUnauthorized(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeUnauthorized, http.StatusUnauthorized, message, args...)
}

// NewMethodNotAllowed creates a 405-level error for unsupported HTTP methods
func NewMethodNotAllowed(message string, args ...interface{}) *ResponseError {
	return NewResponseError(CodeMethodNotAllowed, http.StatusMethodNotAllowed, message, args...)
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
