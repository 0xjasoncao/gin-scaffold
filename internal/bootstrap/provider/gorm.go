package provider

import (
	"github.com/0xjasoncao/gin-scaffold/configs/config"
	"github.com/0xjasoncao/gin-scaffold/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

func InitGorm(config config.Config) (*gorm.DB, error) {

	var dialect gorm.Dialector
	c := config.Gorm
	switch strings.ToLower(c.Use) {
	case "mysql":
		dialect = mysql.Open(config.Mysql.DSN())
	default:
		return nil, errors.Errorf("unsupported database type: %s", c.Use)
	}
	ormCfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Discard,
	}

	if c.Debug {
		ormCfg.Logger = logger.Default
	}

	db, err := gorm.Open(dialect, ormCfg)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(c.MaxOpen)
	sqlDB.SetMaxIdleConns(c.MaxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	return db, nil
}
