package config

import "gin-scaffold/pkg/logging"

// Logger 日志总配置
type Logger struct {
	// Outputs 多种输出的配置集合（key: 输出名称，如 "console"、"file"）
	Outputs []logging.Output `mapstructure:"outputs"`
}
