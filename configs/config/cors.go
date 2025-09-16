package config

type Cors struct {
	Enable           bool     `yaml:"enable"`
	AllowOrigins     []string `yaml:"allow-origins"`
	AllowMethods     []string `yaml:"allow-methods"`
	AllowHeaders     []string `yaml:"allow-headers"`
	AllowCredentials bool     `yaml:"allow-credentials"`
	MaxAge           int      `yaml:"max-age"`
}
