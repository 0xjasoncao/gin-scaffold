package config

import (
	"encoding/json"
	"fmt"
	"gin-scaffold/pkg/redisx"
)

var (
	C = new(Config)
)

type Config struct {
	App        App           `mapstructure:"app"`
	Logger     Logger        `mapstructure:"logger"`
	Jwt        JWT           `mapstructure:"jwt"`
	Http       Http          `mapstructure:"http"`
	Gorm       Gorm          `mapstructure:"gorm"`
	Middleware Middleware    `mapstructure:"middleware"`
	Redis      redisx.Config `mapstructure:"redis"`
	Dir        string
}

func (c *Config) PrintWithJSON() {
	if !c.App.PrintConfig {
		return
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("[CONFIG] - Failed to marshal config to JSON:")
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(" ======================================")
	fmt.Println("[CONFIG] Current configuration:")
	fmt.Println("--------------------------------------")
	fmt.Println(string(b))
	fmt.Println("======================================")
}
