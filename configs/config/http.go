package config

// Http Http config
type Http struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	MaxContentLength int64  `yaml:"max-content-length"`
	ShutdownTimeout  int    `yaml:"shutdown-timeout"`
	ReadTimeout      int64  `yaml:"read-timeout"`
	WriteTimeout     int64  `yaml:"write-timeout"`
	IdleTimeout      int64  `yaml:"idle-timeout"`
	MaxHeaderBytes   int    `yaml:"max-header-bytes"`
	Gzip             Gzip   `yaml:"gzip"`
}
