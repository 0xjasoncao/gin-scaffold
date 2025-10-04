package router

import (
	"github.com/0xjasoncao/gin-scaffold/internal/apis/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(handler *handler.Handler, g *gin.RouterGroup) {

	// v1 route
	v1Handler := handler.V1
	v1 := g.Group("v1")
	{
		loginGroup := v1.Group("login")
		{
			loginGroup.POST("", v1Handler.Login.Login)
			loginGroup.POST("exit", v1Handler.Login.Logout)
		}
	}

}
