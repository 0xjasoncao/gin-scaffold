package config

// Gzip gzip config
type Gzip struct {
	Enable             bool     `yaml:"enable"`
	ExcludedExtensions []string `yaml:"excluded-extensions"`
	ExcludedPath       []string `yaml:"excluded-path"`
}
