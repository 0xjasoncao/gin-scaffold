package config

// App app config
type App struct {
	RunMode     string `yaml:"run-mode"`
	Name        string `yaml:"name"`
	PrintConfig bool   `yaml:"print-config"`
}
