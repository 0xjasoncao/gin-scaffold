package provider

import (
	"context"
	"gin-scaffold/internal/apis"
	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/middleware"
	"gin-scaffold/pkg/redisx"
	"gin-scaffold/pkg/router"
	"gin-scaffold/pkg/token"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	ctx context.Context,
	config *config.Config,
	tokenSrv token.Service,
	//占位 使wire进行生成代码
	_ *apis.RouterHandlers,
	redisFactory *redisx.Factory,
) *gin.Engine {
	logging.WithContext(ctx).Sugar().Infof("[Gin] - Initializing gin engine...")

	//set gin run mode
	gin.SetMode(config.App.RunMode)
	// create gin engine
	app := gin.New()
	app.HandleMethodNotAllowed = true

	// middleware
	{
		app.NoRoute(middleware.NoRoute())
		app.NoMethod(middleware.NoMethod())

		//gzip
		gzipConf := config.Middleware.Gzip
		if gzipConf.Enable {
			app.Use(middleware.GzipMiddleware(&gzipConf))
		}

		routerCfg := config.App.Router

		//trace
		if config.Middleware.TraceId.Enable {
			app.Use(middleware.Trace(&config.Middleware.TraceId))
		}
		//copyBody
		if config.Middleware.CopyBody.Enable {
			app.Use(middleware.CopyBodyMiddleware(&config.Middleware.CopyBody))

		}
		//logger
		if config.Middleware.Logger.Enable {
			app.Use(middleware.LoggerMiddleware())
		} else {
			app.Use(gin.Logger())
		}
		//recovery
		app.Use(middleware.RecoveryMiddleware())
		//auth
		if config.Middleware.Auth.Enable {
			app.Use(middleware.Auth(tokenSrv, &config.Middleware.Auth))
		}
		//cors
		if config.Middleware.Cors.Enable {
			app.Use(middleware.CORSMiddleware(&config.Middleware.Cors))
		}
		//rate limit
		if config.Middleware.RateLimit.Enable {
			//这里默认使用db0存储
			app.Use(middleware.RateLimitMiddleware(redisFactory.GetDefault(), &config.Middleware.RateLimit))
		}

		router.ApplyRoutes(app,
			&router.Options{
				GlobalPrefix:   routerCfg.GlobalPrefix,
				PrintWithStart: routerCfg.PrintWithStart,
			})
	}
	return app
}
