package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type GzipConfig struct {
	Enable             bool     `mapstructure:"enable"`
	ExcludedExtensions []string `mapstructure:"excluded-extensions"`
	ExcludedPath       []string `mapstructure:"excluded-path"`
}

func GzipMiddleware(config *GzipConfig) gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression,
		gzip.WithExcludedExtensions(config.ExcludedExtensions),
		gzip.WithExcludedPaths(config.ExcludedPath))
}
