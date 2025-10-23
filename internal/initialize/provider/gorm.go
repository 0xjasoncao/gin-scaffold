package provider

import (
	"context"
	"fmt"
	"gin-scaffold/internal/initialize/data"
	"strings"
	"time"

	"gin-scaffold/internal/config"
	"gin-scaffold/pkg/logging"

	"gin-scaffold/pkg/errorsx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitGorm(ctx context.Context, config *config.Config) (*gorm.DB, func(), error) {
	logging.WithContext(ctx).Sugar().Infof("[Gorm] - Initializing Gorm...")

	var dialect gorm.Dialector
	c := config.Gorm

	switch strings.ToLower(c.Use) {
	case "mysql":
		c.Mysql.DSN()
		dialect = mysql.Open(c.Mysql.DSN())
		logging.WithContext(ctx).Sugar().Infof("[GORM] - Using MySQL, Addr: %s", fmt.Sprintf("%s:%d", c.Mysql.Host, c.Mysql.Port))
	default:
		return nil, func() {}, errorsx.Errorf("[GORM] - Unsupported database type: %s", c.Use)
	}
	ormCfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction:                   true,
		Logger:                                   logger.Discard,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if c.Debug {
		ormCfg.Logger = logger.Default
	}

	db, err := gorm.Open(dialect, ormCfg)
	if err != nil {
		return nil, func() {}, errorsx.WithStack(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, func() {}, errorsx.WithStack(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		return nil, func() {}, errorsx.WithStack(err)
	}
	sqlDB.SetMaxOpenConns(c.MaxOpen)
	sqlDB.SetMaxIdleConns(c.MaxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	if c.EnableAutoMigrate {
		if strings.ToLower(c.Use) == "mysql" {
			db.Set("gorm:table_options", "ENGINE=InnoDB")
		}
		err := db.AutoMigrate(data.Models()...)
		if err != nil {
			logging.WithContext(ctx).Sugar().Errorf("[GORM] - Database auto-migration failed: %v", err)
		} else {
			logging.WithContext(ctx).Sugar().Info("[GORM] - Database auto-migration completed successfully")
		}
	}

	return db, func() {}, nil
}
