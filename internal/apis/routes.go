package apis

import (
	"gin-scaffold/internal/apis/swagger"
	"gin-scaffold/internal/apis/system"
	"github.com/gin-gonic/gin"
)

type Router interface {
	RegisterRoutes(g *gin.RouterGroup)
}

// RouterHandlers 模块Handler的组合
type RouterHandlers struct {
	Swagger *swagger.Handler
	System  *system.Handlers
}

// RegisterRoutes 注册所有模块路由
func (r *RouterHandlers) RegisterRoutes(g *gin.RouterGroup) {
	r.System.RegisterRoutes(g)
	r.Swagger.RegisterRoutes(g)
}
