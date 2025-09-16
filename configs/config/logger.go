package config

// Logger 日志总配置
type Logger struct {
	// Outputs 多种输出的配置集合（key: 输出名称，如 "console"、"file"）
	Outputs map[string]Output `yaml:"outputs"`
}

// Output 单个输出的配置（如控制台输出、文件输出）
type Output struct {
	Enabled       bool          `yaml:"enabled"`        // 是否启用该输出
	Level         []string      `yaml:"level"`          // 日志级别（如 ["info", "error"]）
	Format        string        `yaml:"format"`         // 日志格式（如 "json"、"text"）
	Path          string        `yaml:"path"`           // 日志文件路径（仅文件输出有效）
	MaxSize       int           `yaml:"max-size"`       // 单个日志文件最大大小（MB）
	MaxBackup     int           `yaml:"max-backup"`     // 最大备份文件数
	MaxAge        int           `yaml:"max-age"`        // 日志文件最大保留天数
	Compress      bool          `yaml:"compress"`       // 是否压缩备份文件
	EncoderConfig EncoderConfig `yaml:"encoder-config"` // 该输出的编码器配置
}

// EncoderConfig 日志编码器配置（定义日志字段的格式）
type EncoderConfig struct {
	TimeKey         string `yaml:"time-key"`         // 时间字段名
	LevelKey        string `yaml:"level-key"`        // 级别字段名
	NameKey         string `yaml:"name-key"`         // 日志器名称字段名
	CallerKey       string `yaml:"caller-key"`       // 调用者字段名
	MessageKey      string `yaml:"message-key"`      // 消息字段名
	StacktraceKey   string `yaml:"stacktrace-key"`   // 堆栈跟踪字段名
	LineEnding      string `yaml:"line-ending"`      // 行结束符
	LevelEncoder    string `yaml:"level-encoder"`    // 级别编码器（如 "lowercase"）
	TimeLayout      string `yaml:"time-layout"`      // 时间编码格式（如 "2006-01-02 15:04:05"）
	DurationEncoder string `yaml:"duration-encoder"` //  duration 编码器
	CallerEncoder   string `yaml:"caller-encoder"`   // 调用者编码器
}
