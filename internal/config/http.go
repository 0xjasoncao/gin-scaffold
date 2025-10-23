package config

// Http Http config
type Http struct {
	Host             string `mapstructure:"host"`
	Port             int    `mapstructure:"port"`
	MaxContentLength int64  `mapstructure:"max-content-length"`
	ShutdownTimeout  int    `mapstructure:"shutdown-timeout"`
	ReadTimeout      int64  `mapstructure:"read-timeout"`
	WriteTimeout     int64  `mapstructure:"write-timeout"`
	IdleTimeout      int64  `mapstructure:"idle-timeout"`
	MaxHeaderBytes   int    `mapstructure:"max-header-bytes"`
}
