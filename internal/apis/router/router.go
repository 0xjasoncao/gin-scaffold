package router

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/docs"

	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"github.com/0xjasoncao/gin-scaffold/pkg/middleware"
	"github.com/0xjasoncao/gin-scaffold/pkg/token"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func NewRouter(
	ctx context.Context,
	config *config.Config,
	handler *handler.Handler,
	tokenSrv token.Service,
	db *gorm.DB,
) *gin.Engine {

	//set gin run mode
	gin.SetMode(config.App.RunMode)
	// create gin engine
	app := gin.New()
	app.HandleMethodNotAllowed = true
	app.NoRoute(middleware.NoRoute())
	app.NoMethod(middleware.NoMethod())

	//gzip
	gzipConf := config.Http.Gzip
	if gzipConf.Enable {
		app.Use(gzip.Gzip(gzip.DefaultCompression,
			gzip.WithExcludedExtensions(gzipConf.ExcludedExtensions),
			gzip.WithExcludedPaths(gzipConf.ExcludedPath)))
	}

	apiPrefix := "/api"

	//trace
	app.Use(middleware.Trace(middleware.NotSkippedPathPrefix(apiPrefix)))
	//copyBody
	app.Use(middleware.CopyBodyMiddleware(config.Http, middleware.NotSkippedPathPrefix(apiPrefix)))
	//logger
	app.Use(middleware.LoggerMiddleware())
	//recovery
	app.Use(middleware.RecoveryMiddleware())
	//auth
	app.Use(middleware.Auth(tokenSrv, middleware.SkippedPathPrefix("/api/v1/login", "/swagger"), middleware.NotSkippedPathPrefix(apiPrefix)))
	//cors
	if config.Cors.Enable {
		app.Use(middleware.CORSMiddleware(config.Cors))
	}

	api := app.Group(apiPrefix)
	//register  router
	RegisterUserRouter(handler, api)
	//swagger
	docs.SwaggerInfo.Version = "1.0"
	app.GET("/swagger/*any", gin.BasicAuth(gin.Accounts{"admin": "123456"}), ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes := app.Routes()
	logging.WithContext(ctx).Sugar().Infof("Registered %d routes:", len(routes))
	for _, r := range routes {
		logging.WithContext(ctx).Sugar().Infof("[Route] %-5s-> %s (Handler: %s)", r.Method, r.Path, r.Handler)
	}

	return app
}
