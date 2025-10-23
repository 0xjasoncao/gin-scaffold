package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsConfig struct {
	Enable           bool     `mapstructure:"enable"`
	AllowOrigins     []string `mapstructure:"allow-origins"`
	AllowMethods     []string `mapstructure:"allow-methods"`
	AllowHeaders     []string `mapstructure:"allow-headers"`
	AllowCredentials bool     `mapstructure:"allow-credentials"`
	MaxAge           int      `mapstructure:"max-age"`
}

func CORSMiddleware(config *CorsConfig) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     config.AllowOrigins,
		AllowMethods:     config.AllowMethods,
		AllowHeaders:     config.AllowHeaders,
		AllowCredentials: config.AllowCredentials,
		MaxAge:           time.Second * time.Duration(config.MaxAge),
	})
}
