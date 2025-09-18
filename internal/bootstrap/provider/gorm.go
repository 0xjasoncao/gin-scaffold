package provider

import (
	"context"
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/internal/model"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"github.com/0xjasoncao/gin-scaffold/pkg/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

func InitGorm(ctx context.Context, config *config.Config) (*gorm.DB, func(), error) {

	var dialect gorm.Dialector
	c := config.Gorm
	switch strings.ToLower(c.Use) {
	case "mysql":
		dialect = mysql.Open(config.Mysql.DSN())
	default:
		return nil, func() {}, errors.Errorf("unsupported database type: %s", c.Use)
	}
	ormCfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:                                   logger.Discard,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if c.Debug {
		ormCfg.Logger = logger.Default
	}

	db, err := gorm.Open(dialect, ormCfg)
	if err != nil {
		return nil, func() {}, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, func() {}, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, func() {}, err
	}
	sqlDB.SetMaxOpenConns(c.MaxOpen)
	sqlDB.SetMaxIdleConns(c.MaxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	if c.EnableAutoMigrate {
		if strings.ToLower(c.Use) == "mysql" {
			db.Set("gorm:table_options", "ENGINE=InnoDB")
		}
		err := db.AutoMigrate(model.Models()...)
		if err != nil {
			logging.WithContext(ctx).Sugar().Errorf("AutoMigrate failed, error:%v", err)
		}
	}

	return db, func() {}, nil
}
