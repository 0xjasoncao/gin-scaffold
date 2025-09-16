package middleware

import "github.com/gin-gonic/gin"

type Options struct {
	AllowedPathPrefix    []string
	NotAllowedPathPrefix []string
}

type OptionsFunc func(opt *Options)

func AllowPathPrefix(prefix ...string) OptionsFunc {
	return func(opt *Options) {
		opt.AllowedPathPrefix = prefix
	}
}

func NotAllowPathPrefix(prefix ...string) OptionsFunc {
	return func(opt *Options) {
		opt.NotAllowedPathPrefix = prefix
	}
}

func NeedSkip(c *gin.Context, optionsFunc ...OptionsFunc) bool {
	var o Options
	for _, opt := range optionsFunc {
		opt(&o)
	}
	path := c.Request.URL.Path
	pathLen := len(path)

	for _, p := range o.AllowedPathPrefix {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return true
		}
	}
	for _, p := range o.NotAllowedPathPrefix {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return false
		}
	}
	return false

}
