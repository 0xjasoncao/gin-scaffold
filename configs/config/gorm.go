package config

type Gorm struct {
	Use         string `yaml:"use"`
	Debug       bool   `yaml:"debug"`
	MaxLifetime int    `yaml:"max-lifetime"`
	MaxOpen     int    `yaml:"max-open"`
	MaxIdle     int    `yaml:"max-idle"`
}
