package logging

// Output 单个输出的配置（如控制台输出、文件输出）
type Output struct {
	Name          string        `mapstructure:"name"`           //名称
	Enabled       bool          `mapstructure:"enabled"`        // 是否启用该输出
	Level         []string      `mapstructure:"level"`          // 日志级别（如 ["info", "error"]）
	Format        string        `mapstructure:"format"`         // 日志格式（如 "json"、"text"）
	Path          string        `mapstructure:"path"`           // 日志文件路径（仅文件输出有效）
	MaxSize       int           `mapstructure:"max-size"`       // 单个日志文件最大大小（MB）
	MaxBackup     int           `mapstructure:"max-backup"`     // 最大备份文件数
	MaxAge        int           `mapstructure:"max-age"`        // 日志文件最大保留天数
	Compress      bool          `mapstructure:"compress"`       // 是否压缩备份文件
	EncoderConfig EncoderConfig `mapstructure:"encoder-config"` // 该输出的编码器配置
}

// EncoderConfig 日志编码器配置（定义日志字段的格式）
type EncoderConfig struct {
	TimeKey         string `mapstructure:"time-key"`         // 时间字段名
	LevelKey        string `mapstructure:"level-key"`        // 级别字段名
	NameKey         string `mapstructure:"name-key"`         // 日志器名称字段名
	CallerKey       string `mapstructure:"caller-key"`       // 调用者字段名
	MessageKey      string `mapstructure:"message-key"`      // 消息字段名
	StacktraceKey   string `mapstructure:"stacktrace-key"`   // 堆栈跟踪字段名
	LineEnding      string `mapstructure:"line-ending"`      // 行结束符
	LevelEncoder    string `mapstructure:"level-encoder"`    // 级别编码器（如 "lowercase"）
	TimeLayout      string `mapstructure:"time-layout"`      // 时间编码格式（如 "2006-01-02 15:04:05"）
	DurationEncoder string `mapstructure:"duration-encoder"` //  duration 编码器
	CallerEncoder   string `mapstructure:"caller-encoder"`   // 调用者编码器
}
