package role

import (
	"time"

	"github.com/0xjasoncao/gin-scaffold/internal/domain/user/rolemenu"
)

type Role struct {
	ID        string
	Name      string
	Sequence  int
	Memo      *string
	Status    int
	Creator   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	RoleMenus rolemenu.RoleMenus
}

type Roles []*Role
