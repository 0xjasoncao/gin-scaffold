package v1

import (
	"gin-scaffold/internal/apis/system/request"
	"gin-scaffold/internal/apis/system/response"
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/api"
	"gin-scaffold/pkg/errorsx"
	"gin-scaffold/pkg/router"
	"gin-scaffold/pkg/token"
	"gin-scaffold/pkg/utils/structureutil"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSrv  system.UserService
	TokenSrv token.Service
}

func NewUserHandler(loginSrv system.UserService, tokenSrv token.Service) *UserHandler {
	user := &UserHandler{
		UserSrv:  loginSrv,
		TokenSrv: tokenSrv,
	}
	router.AddRoute(func(group *gin.RouterGroup) {
		g := group.Group("/system/v1/user")
		g.POST("login", user.Login)
		g.POST("logout", user.Logout)
	})

	return user
}

// Login 登录
//
//	@Summary		login
//	@Description	login by mobile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.LoginRequest	true	"登录参数"
//	@Success		200		{object}	api.Response{data=response.LoginResponse}
//	@Router			/system/v1/user/login/   [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var in request.LoginRequest
	err := api.ParseJSON(ctx, &in)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	user, err := h.UserSrv.Login(ctx.Request.Context(), system.UserQueryParam{
		Mobile:   in.Mobile,
		Password: in.Password,
	})
	if err != nil {
		api.ResError(ctx, err)
		return
	}

	res := response.LoginResponse{}
	structureutil.Copy(&user, &res)

	issuedTokenInfo, err := h.TokenSrv.IssuingToken(ctx, token.Payload{UserID: user.ID})
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	res.TokenInfo = issuedTokenInfo

	api.ResData(ctx, res)

}

// Logout 登出
//
//	@Summary		logout
//	@Description	login by mobile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"access_token"
//	@Success		200				{object}	api.Response
//	@Router			/system/v1/user/logout [post]
func (h *UserHandler) Logout(ctx *gin.Context) {
	accessToken := api.GetToken(ctx)
	if accessToken == "" {
		api.ResError(ctx, errorsx.NewUnauthorized("用户未登录"))
		return
	}
	if err := h.TokenSrv.DestroyToken(ctx, accessToken); err != nil {
		api.ResError(ctx, err)
		return
	}
	api.ResWithMessage(ctx, "登出成功")

}
