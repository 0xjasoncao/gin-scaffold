package response

import (
	"gin-scaffold/internal/domain/system"
	"gin-scaffold/pkg/utils/structureutil"
	"time"
)

type RoleQueryResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Comment   string    `json:"comment"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"` // CreatedAt
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt
}

func RolesFromDomain(roles system.Roles) []*RoleQueryResponse {
	list := make([]*RoleQueryResponse, len(roles))
	for i, item := range roles {
		list[i] = &RoleQueryResponse{}
		structureutil.Copy(item, list[i])
	}
	return list
}
