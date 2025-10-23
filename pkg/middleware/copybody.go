package middleware

import (
	"bytes"
	"gin-scaffold/pkg/api"
	"io"
	"net/http"

	"gin-scaffold/pkg/errorsx"

	"github.com/gin-gonic/gin"
)

const (
	ReqBodyKey = "request_body"
)

type CopyBodyConfig struct {
	MaxContentLength  int64    `mapstructure:"max-content-length"`
	Enable            bool     `mapstructure:"enable"`
	SkipPathPrefix    []string `mapstructure:"skip-path-prefix"`
	NotSkipPathPrefix []string `mapstructure:"not-skip-path-prefix"`
}

func CopyBodyMiddleware(conf *CopyBodyConfig) gin.HandlerFunc {

	var maxMemory int64 = 10 << 20 // 10 MB
	if v := conf.MaxContentLength; v > 0 {
		maxMemory = v << 20
	}

	return func(c *gin.Context) {

		if NeedSkip(c, SkippedPathPrefix(conf.SkipPathPrefix...), NotSkippedPathPrefix(conf.NotSkipPathPrefix...)) {
			c.Next()
			return
		}

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxMemory)

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			var e *http.MaxBytesError
			if errorsx.As(err, &e) {
				err = errorsx.NewRequestEntityTooLarge(err.Error())
			}
			api.ResError(c, err)
			return
		}
		_ = c.Request.Body.Close()
		c.Set(ReqBodyKey, body)
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		c.Next()
	}
}
