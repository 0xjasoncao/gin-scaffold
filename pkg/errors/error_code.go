package errors

const (
	CodeInternalError         = 50001 // Internal server error
	CodeNotFound              = 40401 // Resource not found
	CodeBadRequest            = 40001 // Invalid request parameters
	CodeTooManyRequests       = 42901 // Too many requests
	CodeUnauthorized          = 40101 // Authentication required
	CodeRequestEntityTooLarge = 41301 // Authentication required
	CodeMethodNotAllowed      = 40501 // HTTP method not allowed
)
