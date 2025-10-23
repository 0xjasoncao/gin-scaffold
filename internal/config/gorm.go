package config

import "fmt"

type Gorm struct {
	Use               string `mapstructure:"use"`
	Debug             bool   `mapstructure:"debug"`
	MaxLifetime       int    `mapstructure:"max-lifetime"`
	MaxOpen           int    `mapstructure:"max-open"`
	MaxIdle           int    `mapstructure:"max-idle"`
	EnableAutoMigrate bool   `mapstructure:"enable-auto-migrate"`
	Mysql             MySQL  `mapstructure:"mysql"`
}

type MySQL struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	DBName     string `mapstructure:"db-name"`
	Parameters string `mapstructure:"parameters"`
}

func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}
