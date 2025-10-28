package system

import (
	"gin-scaffold/internal/apis/handler/system/v1"
	"github.com/gin-gonic/gin"
)

// V1 V1版本的handler
type V1 struct {
	User *v1.UserHandler
	Role *v1.RoleHandler
}

// Handlers 该模块的不同版本Handler组合
type Handlers struct {
	V1 *V1
}

// RegisterRoutes 注册该模块下所有路由
func (hs *Handlers) RegisterRoutes(g *gin.RouterGroup) {
	// v1
	{
		h := hs.V1
		system := g.Group("v1/system")
		//user
		{
			system.Group("user").
				POST("login", h.User.Login).
				POST("logout", h.User.Logout)
		}
		//role
		{
			system.Group("role").
				POST("create", h.Role.Create).
				POST("delete", h.Role.Delete)

		}

	}

}
