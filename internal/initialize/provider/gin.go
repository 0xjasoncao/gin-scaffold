package provider

import (
	"context"
	"gin-scaffold/internal/apis"
	"gin-scaffold/internal/config"
	middleware2 "gin-scaffold/pkg/core/middleware"
	"gin-scaffold/pkg/core/token"
	"gin-scaffold/pkg/logging"
	"gin-scaffold/pkg/redisx"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	ctx context.Context,
	config *config.Config,
	tokenSrv token.Service,
	rhs *apis.RouterHandlers,
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
		app.NoRoute(middleware2.NoRoute())
		app.NoMethod(middleware2.NoMethod())

		//gzip
		gzipConf := config.Middleware.Gzip
		if gzipConf.Enable {
			app.Use(middleware2.GzipMiddleware(&gzipConf))
		}

		routerCfg := config.App.Router

		//trace
		if config.Middleware.TraceId.Enable {
			app.Use(middleware2.Trace(&config.Middleware.TraceId))
		}
		//copyBody
		if config.Middleware.CopyBody.Enable {
			app.Use(middleware2.CopyBodyMiddleware(&config.Middleware.CopyBody))

		}
		//logger
		if config.Middleware.Logger.Enable {
			app.Use(middleware2.LoggerMiddleware())
		} else {
			app.Use(gin.Logger())
		}
		//recovery
		app.Use(middleware2.RecoveryMiddleware())
		//auth
		if config.Middleware.Auth.Enable {
			app.Use(middleware2.Auth(tokenSrv, &config.Middleware.Auth))
		}
		//cors
		if config.Middleware.Cors.Enable {
			app.Use(middleware2.CORSMiddleware(&config.Middleware.Cors))
		}
		//rate limit
		if config.Middleware.RateLimit.Enable {
			//这里默认使用db0存储,实际情况自行选择
			app.Use(middleware2.RateLimitMiddleware(redisFactory.GetDefault(), &config.Middleware.RateLimit))
		}
		//register all routes
		rhs.RegisterRoutes(app.Group(routerCfg.GlobalPrefix))

		if routerCfg.PrintWithStart {
			routes := app.Routes()
			logging.Logger().Sugar().Infof("[Route] - Registered %d routes:", len(routes))
			for _, r := range routes {
				logging.Logger().Sugar().Infof("%-5s-> %s (Handler: %s)", r.Method, r.Path, r.Handler)
			}
		}
	}
	return app
}
