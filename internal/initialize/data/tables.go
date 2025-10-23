package data

import "gin-scaffold/internal/domain/system"

// Models 用于gorm.AutoMigrate 自动创建数据库表结构
func Models() []interface{} {
	register(&system.User{})
	register(&system.Role{})
	register(&system.RoleMenu{})
	register(&system.Menu{})
	return models
}

var models []interface{}

func register(model interface{}) {
	models = append(models, model)
}
