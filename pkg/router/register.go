package router

import (
	"gin-scaffold/pkg/logging"
	"github.com/gin-gonic/gin"
)

// routes 用于存储路由组信息
var routes = make([]*route, 0)

// route 路由
type route struct {
	setupFunc func(group *gin.RouterGroup)
	subPath   []string
}

// AddRoute 添加路由
func AddRoute(setupFunc func(group *gin.RouterGroup)) {
	g := &route{
		setupFunc: setupFunc,
	}
	routes = append(routes, g)
}

type Options struct {

	// GlobalPrefix 全局路由前缀
	GlobalPrefix string

	// PrintWithStart 启动时打印所有路由
	PrintWithStart bool
}

// ApplyRoutes 根据Options提供的参数进行路由初始化
func ApplyRoutes(engine *gin.Engine, opt *Options) {
	for _, router := range routes {
		router.setupFunc(engine.Group(opt.GlobalPrefix))
	}

	if opt.PrintWithStart {
		routes := engine.Routes()
		logging.Logger().Sugar().Infof("[Route] - Registered %d routes:", len(routes))
		for _, r := range routes {
			logging.Logger().Sugar().Infof("%-5s-> %s (Handler: %s)", r.Method, r.Path, r.Handler)
		}
	}
}
