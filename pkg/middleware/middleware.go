package middleware

import "github.com/gin-gonic/gin"

type Skipper struct {
	SkipPathPrefix    []string
	NotSkipPathPrefix []string
}

type SkipFunc func(opt *Skipper)

func SkippedPathPrefix(prefix ...string) SkipFunc {
	return func(opt *Skipper) {
		opt.SkipPathPrefix = prefix
	}
}

func NotSkippedPathPrefix(prefix ...string) SkipFunc {
	return func(opt *Skipper) {
		opt.NotSkipPathPrefix = prefix
	}
}

func NeedSkip(c *gin.Context, optionsFunc ...SkipFunc) bool {
	var o Skipper
	for _, opt := range optionsFunc {
		opt(&o)
	}
	path := c.Request.URL.Path
	pathLen := len(path)

	for _, p := range o.SkipPathPrefix {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return true
		}
	}
	for _, p := range o.NotSkipPathPrefix {
		if pl := len(p); pathLen >= pl && path[:pl] == p {
			return false
		}
	}
	return false

}
