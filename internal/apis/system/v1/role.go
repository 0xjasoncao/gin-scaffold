package v1

import (
	"gin-scaffold/internal/apis/system/request"
	"gin-scaffold/internal/apis/system/response"
	"gin-scaffold/internal/domain/shared"
	"gin-scaffold/internal/domain/system"
	ginutil "gin-scaffold/pkg/core/ginutil"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	RoleSrv system.RoleService
}

func NewRoleHandler(roleService system.RoleService) *RoleHandler {
	return &RoleHandler{RoleSrv: roleService}
}

// Create 创建角色
//
//	@Summary		创建角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.RoleCreateRequest	true	"创建参数"
//	@Success		200		{object}	ginutil.Response "创建成功"
//	@Router			/system/role/create   [post]
func (r *RoleHandler) Create(ctx *gin.Context) {
	var in request.RoleCreateRequest
	err := ginutil.ParseJSON(ctx, &in)
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	err = r.RoleSrv.Create(ctx.Request.Context(), system.Role{
		Name:    in.Name,
		Comment: in.Comment,
	})
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	ginutil.ResOKWithMessage(ctx, "创建成功")

}

// Delete 批量删除角色
//
//	@Summary		删除角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.RoleDeleteRequest	true	"删除参数"
//	@Success		200		{object}	ginutil.Response "删除成功"
//	@Router			/system/role/delete   [post]
func (r *RoleHandler) Delete(ctx *gin.Context) {
	var in request.RoleDeleteRequest
	err := ginutil.ParseJSON(ctx, &in)
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	err = r.RoleSrv.Delete(ctx.Request.Context(), in.IDs)
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	ginutil.ResOKWithMessage(ctx, "删除成功")

}

// Update 编辑角色
//
//	@Summary		编辑角色
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.RoleUpdateRequest	true	"编辑参数"
//	@Success		200		{object}	ginutil.Response "编辑成功"
//	@Router			/system/role/update   [post]
func (r *RoleHandler) Update(ctx *gin.Context) {
	var in request.RoleUpdateRequest
	err := ginutil.ParseJSON(ctx, &in)
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	err = r.RoleSrv.Update(ctx.Request.Context(), system.Role{
		BasicInfo: shared.BasicInfo{
			ID: in.ID,
		},
		Name:    in.Name,
		Comment: in.Comment,
		Status:  in.Status,
	})
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	ginutil.ResOKWithMessage(ctx, "编辑成功")
}

// Query 查询角色列表
//
//	@Summary		查询角色列表
//	@Tags			角色
//	@Accept			json
//	@Produce		json
//	@Param			request	body		request.RoleQueryRequest	true	"查询参数"
//	@Success		200		{object}	ginutil.Response{Meta=core.Pagination,Data=[]response.RoleQueryResponse}
//	@Router			/system/role/query   [get]
func (r *RoleHandler) Query(ctx *gin.Context) {
	var in request.RoleQueryRequest
	err := ginutil.ParseQuery(ctx, &in)
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	roles, pagination, err := r.RoleSrv.Query(ctx.Request.Context(), system.RoleQueryParam{
		PageParam: in.PageParam,
	})
	if err != nil {
		ginutil.ResError(ctx, err)
		return
	}
	ginutil.ResPage(ctx, response.RolesFromDomain(roles), pagination)
}
