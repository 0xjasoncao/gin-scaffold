package model

import "github.com/0xjasoncao/gin-scaffold/internal/repository/model/user"

func Models() []interface{} {
	var models []interface{}
	models = append(models, &user.User{})
	return models
}
