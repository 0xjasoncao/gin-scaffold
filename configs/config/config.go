package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/0xjasoncao/gin-scaffold/pkg/config/loader"
	"github.com/0xjasoncao/gin-scaffold/pkg/config/parser"
)

var (
	C = new(Config)
)

type Config struct {
	App    App    `yaml:"app"`
	Logger Logger `yaml:"logger"`
	Jwt    JWT    `yaml:"jwt"`
	Http   Http   `yaml:"http"`
	Gorm   Gorm   `yaml:"gorm"`
	Cors   Cors   `yaml:"cors"`
	Mysql  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	Cache  Cache  `yaml:"cache"`
}

// Load load config
func Load(configDir string) error {

	entries, err := os.ReadDir(configDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(configDir, entry.Name())
		fl := loader.NewLocalConfigLoader(fullPath)
		data, err := fl.Load()
		if err != nil {
			return err
		}
		index := strings.LastIndex(fullPath, ".")
		getParser, err := parser.GetParser(fullPath[index+1:])
		if err != nil {
			return err
		}
		if err = getParser.Parse(data, C); err != nil {
			return err
		}
	}

	return nil
}

func (c *Config) PrintWithJSON() {
	if !c.App.PrintConfig {
		return
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		fmt.Println("[CONFIG] Failed to marshal config to JSON:")
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println(" ======================================")
	fmt.Println("[CONFIG] Current configuration:")
	fmt.Println("--------------------------------------")
	fmt.Println(string(b))
	fmt.Println("======================================")
}
