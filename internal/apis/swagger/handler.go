package swagger

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

type (
	Handler struct {
		Config *config.Swagger
	}
)

// RegisterRoutes 注册swagger路由
func (s *Handler) RegisterRoutes(g *gin.RouterGroup) {
	docs.SwaggerInfo.Title = "Gin Scaffold API"
	docs.SwaggerInfo.Description = "This is a sample server caller server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.InfoInstanceName = "v1"
	swag.Register(docs.SwaggerInfo.InfoInstanceName, docs.SwaggerInfo)
	swaggerConf := s.Config
	if swaggerConf.Enable {
		auth := func(c *gin.Context) {}
		if swaggerConf.Enable {
			auth = gin.BasicAuth(gin.Accounts{swaggerConf.Auth.Account: swaggerConf.Auth.Password})
		}
		g.GET("/swagger/*any", auth, ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName(docs.SwaggerInfo.InfoInstanceName)))
	}
}

func NewHandler(config *config.Config) *Handler {
	return &Handler{
		Config: &config.App.Swagger,
	}
}
