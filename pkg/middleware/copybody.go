package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/utils/api"

	"github.com/gin-gonic/gin"
)

func CopyBodyMiddleware(cfg config.Http, skipFunc ...SkipFunc) gin.HandlerFunc {

	var maxMemory int64 = 10 << 20 // 10 MB
	if v := cfg.MaxContentLength; v > 0 {
		maxMemory = v << 20
	}

	return func(c *gin.Context) {

		if NeedSkip(c, skipFunc...) {
			c.Next()
			return
		}

		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxMemory)

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			var e *http.MaxBytesError
			if errors.As(err, &e) {
				err = errors.NewRequestEntityTooLarge(err.Error())
			}
			api.ResError(c, err)
			return
		}
		c.Request.Body.Close()
		c.Set(api.ReqBodyKey, body)
		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		c.Next()
	}
}
