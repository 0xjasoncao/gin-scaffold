package config

// App app config
type App struct {
	RunMode     string  `mapstructure:"run-mode"`
	Name        string  `mapstructure:"name"`
	PrintConfig bool    `mapstructure:"print-config"`
	Router      Router  `mapstructure:"router"`
	Swagger     Swagger `mapstructure:"swagger"`
}
type Router struct {

	// GlobalPrefix 全局路由前缀
	GlobalPrefix string `mapstructure:"global-prefix"`
	// PrintWithStart 启动时打印所有路由
	PrintWithStart bool `mapstructure:"print-with-start"`
}
type Swagger struct {
	Enable bool `mapstructure:"enable"`
	Auth   struct {
		Enable   bool   `mapstructure:"enable"`
		Account  string `mapstructure:"account"`
		Password string `mapstructure:"password"`
	} `mapstructure:"auth"`
}
