package domain

import (
	"github.com/0xjasoncao/gin-scaffold/internal/domain/user"
	"github.com/0xjasoncao/gin-scaffold/internal/domain/user/role"
	"github.com/0xjasoncao/gin-scaffold/internal/domain/user/rolemenu"
)

func Models() []interface{} {
	register(&user.User{})
	register(&role.Role{})
	register(&rolemenu.RoleMenu{})
	return models
}

var models []interface{}

func register(model interface{}) {
	models = append(models, model)
}
