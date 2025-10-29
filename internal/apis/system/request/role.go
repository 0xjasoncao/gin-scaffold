package request

import (
	"gin-scaffold/pkg/core"
)

type RoleCreateRequest struct {
	Name    string `json:"name" binding:"required"`
	Comment string `json:"comment"`
}

type RoleDeleteRequest struct {
	IDs []uint64 `json:"ids" binding:"required,min=1"`
}

type RoleUpdateRequest struct {
	ID      uint64 `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Comment string `json:"comment"`
	Status  int    `json:"status"`
}

type RoleQueryRequest struct {
	core.PageParam
}
