package v1

import (
	"gin-scaffold/internal/apis/handler/system/request"
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/api"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	RoleSrv system.RoleService
}

func NewRoleHandler(roleService system.RoleService) *RoleHandler {
	return &RoleHandler{RoleSrv: roleService}
}

func (r *RoleHandler) Create(ctx *gin.Context) {
	var in request.RoleCreateRequest
	err := api.ParseJSON(ctx, &in)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	err = r.RoleSrv.Create(ctx.Request.Context(), system.Role{
		Name:    in.Name,
		Comment: in.Comment,
	})
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	api.ResOKWithMessage(ctx, "创建成功")

}

func (r *RoleHandler) Delete(ctx *gin.Context) {
	var in request.RoleDeleteRequest
	err := api.ParseJSON(ctx, &in)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	err = r.RoleSrv.Delete(ctx.Request.Context(), in.IDs)
	if err != nil {
		api.ResError(ctx, err)
		return
	}
	api.ResOKWithMessage(ctx, "删除成功")

}
