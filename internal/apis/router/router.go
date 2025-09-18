package router

import (
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler"
	"github.com/0xjasoncao/gin-scaffold/pkg/middleware"
	"github.com/0xjasoncao/gin-scaffold/pkg/token"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(
	config *config.Config,
	handler *handler.Handler,
	tokenSrc token.Service,
	db *gorm.DB,
) *gin.Engine {

	//set gin run mode
	gin.SetMode(config.App.RunMode)
	// create gin engine
	app := gin.New()
	//Gzip
	gzipConf := config.Http.Gzip

	if gzipConf.Enable {
		app.Use(gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedExtensions(gzipConf.ExcludedExtensions),
			gzip.WithExcludedPaths(gzipConf.ExcludedPath)))
	}

	//Trace
	app.Use(middleware.Trace())

	//Auth
	app.Use(middleware.Auth(tokenSrc))
	//CopyBody
	app.Use(middleware.CopyBodyMiddleware(config.Http))
	//Logger
	app.Use(middleware.LoggerMiddleware())
	// 配置CORS
	if config.Cors.Enable {
		app.Use(middleware.CORSMiddleware(config.Cors))

	}

	// v1 route
	v1 := app.Group("api/v1")
	{
		group := v1.Group("/user")
		{
			group.GET("/info", handler.V1.User.Info)
		}
	}

	return app
}
