package logging

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// Logging field keys used in structured log entries
const (
	TraceIDKey = "trace_id" // Key for trace ID
	UserIDKey  = "user_id"  // Key for user ID
	TagKey     = "tag"      // Key for custom tags
	StackKey   = "stack"    // Key for stack trace

)

// Context key types (unexported) to avoid collisions with other context keys.
// Using custom struct types ensures type safety and uniqueness.
type (
	CtxTraceIDKey struct{} // Context key type for trace ID
	CtxUserIDKey  struct{} // Context key type for user ID
	CtxTagKey     struct{} // Context key type for custom tags
	CtxStackKey   struct{} // Context key type for stack traces
)

// NewTraceIDContext returns a new context with the given trace ID.
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, CtxTraceIDKey{}, traceID)
}

// FormTraceIDContext retrieves the trace ID from the context, if present.
func FormTraceIDContext(ctx context.Context) string {
	if v := ctx.Value(CtxTraceIDKey{}); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewUserIDContext returns a new context with the given user ID.
func NewUserIDContext(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, CtxUserIDKey{}, userID)
}

// FormUserIDContext retrieves the user ID from the context, if present.
func FormUserIDContext(ctx context.Context) string {
	if v := ctx.Value(CtxUserIDKey{}); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewTagContext returns a new context with the given custom tag.
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, CtxTagKey{}, tag)
}

// FormTagContext retrieves the custom tag from the context, if present.
func FormTagContext(ctx context.Context) string {
	if v := ctx.Value(CtxTagKey{}); v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// NewStackContext returns a new context with the given error as the stack trace.
func NewStackContext(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, CtxStackKey{}, err)
}

// FormStackContext retrieves the stack trace (error) from the context, if present.
func FormStackContext(ctx context.Context) error {
	if v := ctx.Value(CtxStackKey{}); v != nil {
		if e, ok := v.(error); ok {
			return e
		}
	}
	return nil
}

func Logger() *zap.Logger {

	return zap.L()
}

// WithContext extracts contextual logging fields and returns a new logger instance
// enriched with trace ID, user ID, tag, and stack trace (if available).
func WithContext(ctx context.Context) *zap.Logger {
	var fields []zap.Field

	if traceID := FormTraceIDContext(ctx); traceID != "" {
		fields = append(fields, zap.String(TraceIDKey, traceID))
	}
	if userID := FormUserIDContext(ctx); userID != "" {
		fields = append(fields, zap.String(UserIDKey, userID))
	}
	if tag := FormTagContext(ctx); tag != "" {
		fields = append(fields, zap.String(TagKey, tag))
	}
	if stack := FormStackContext(ctx); stack != nil {
		fields = append(fields, zap.String(StackKey, fmt.Sprintf("%+v", stack)))
	}

	// Return a logger with additional context-specific fields attached.
	return Logger().With(fields...)
}
