package docs

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

type SwaggerHandler struct {
}

func NewSwaggerHandler(config *config.Config) *SwaggerHandler {
	SwaggerInfo.Title = "Gin Scaffold API"
	SwaggerInfo.Description = "This is a sample server caller server."
	SwaggerInfo.Version = "1.0"
	SwaggerInfo.Host = "localhost:8080"
	SwaggerInfo.BasePath = "/api/"
	SwaggerInfo.InfoInstanceName = "v1"
	swag.Register(SwaggerInfo.InfoInstanceName, SwaggerInfo)
	swaggerConf := config.App.Swagger
	if swaggerConf.Enable {
		auth := func(c *gin.Context) {}
		if config.App.Swagger.Auth.Enable {
			auth = gin.BasicAuth(gin.Accounts{swaggerConf.Auth.Account: swaggerConf.Auth.Password})
		}
		router.AddRoute(func(group *gin.RouterGroup) {
			group.GET("/swagger/*any", auth, ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.InstanceName(SwaggerInfo.InfoInstanceName)))
		})
	}
	return &SwaggerHandler{}
}
