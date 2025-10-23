package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"time"

	"gin-scaffold/pkg/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoggerConfig struct {
	Enable bool `mapstructure:"enable"`
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		request := c.Request
		method := request.Method
		path := request.URL.Path
		var fields []zap.Field

		if method == http.MethodGet {
			fields = append(fields, zap.String("query", request.URL.RawQuery))
		}

		if method == http.MethodPost || method == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType != "multipart/form-data" {
				if v, ok := c.Get(ReqBodyKey); ok {
					if b, ok := v.([]byte); ok {
						if json.Valid(b) {
							var compactJSON bytes.Buffer
							err := json.Compact(&compactJSON, b)
							if err == nil {
								fields = append(fields, zap.String("body", compactJSON.String()))
							} else {
								fields = append(fields, zap.ByteString("body", b))

							}
						}
					}
				}
			}
		}
		c.Next()

		ctx := c.Request.Context()
		statusCode := c.Writer.Status()
		duration := time.Since(start)
		logging.WithContext(ctx).Info(fmt.Sprintf("[HTTP] - %3d - %v - %s - %s - %s", statusCode, duration, c.ClientIP(), method, path),
			fields...)
	}
}
