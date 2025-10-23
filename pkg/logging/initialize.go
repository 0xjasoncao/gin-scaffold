package logging

import (
	"context"
	"github.com/natefinch/lumberjack"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// InitLogger initializes the global logger
func InitLogger(ctx context.Context, outputs ...Output) error {

	if len(outputs) == 0 {
		defaultLogger, _ := zap.NewDevelopment(zap.AddStacktrace(zapcore.ErrorLevel))
		zap.ReplaceGlobals(defaultLogger)
		Logger().Warn("Logging configuration is not specified, default logging configuration will be used")
		return nil
	}

	// Collect cores for all output targets
	var cores []zapcore.Core

	// Create Core for each output target
	for _, output := range outputs {
		if !output.Enabled {
			continue
		}
		// Create log writer
		writeSyncer, err := createWriteSyncer(output)
		if err != nil {
			return errors.Wrapf(err, "Failed to create output target [%s]", output.Name)
		}

		// Create log encoder
		encoder := createEncoder(output)

		// Parse output level
		levelEnabler := createLevelEnabler(output.Level)

		// Create Core and add to list
		core := zapcore.NewCore(encoder, writeSyncer, levelEnabler)
		cores = append(cores, core)
	}

	// Create multi-Core Logger
	core := zapcore.NewTee(cores...)

	// Add caller information and stack trace
	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	zap.ReplaceGlobals(logger)
	WithContext(ctx).Sugar().Info(" Logger initialized successfully.")

	return nil
}

// 创建日志编码器
func createEncoder(output Output) zapcore.Encoder {
	// 转换自定义EncoderConfig为zap的EncoderConfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:          output.EncoderConfig.TimeKey,
		LevelKey:         output.EncoderConfig.LevelKey,
		NameKey:          output.EncoderConfig.NameKey,
		CallerKey:        output.EncoderConfig.CallerKey,
		MessageKey:       output.EncoderConfig.MessageKey,
		StacktraceKey:    output.EncoderConfig.StacktraceKey,
		LineEnding:       output.EncoderConfig.LineEnding,
		EncodeLevel:      getLevelEncoder(output.EncoderConfig.LevelEncoder),
		EncodeTime:       getTimeEncoder(output.EncoderConfig.TimeLayout),
		EncodeDuration:   getDurationEncoder(output.EncoderConfig.DurationEncoder),
		EncodeCaller:     getCallerEncoder(output.EncoderConfig.CallerEncoder),
		ConsoleSeparator: " - ",
	}

	// 根据格式选择编码器
	if output.Format == "json" {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}
func createLevelEnabler(levels []string) zapcore.LevelEnabler {
	if levels == nil || len(levels) == 0 {
		return zapcore.InfoLevel
	}
	return zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		for _, lvlStr := range levels {
			lvl, err := parseLevel(lvlStr)
			if err != nil {
				continue
			}
			if level == lvl {
				return true
			}
		}
		return false
	})
}

// 创建日志写入器
func createWriteSyncer(output Output) (zapcore.WriteSyncer, error) {
	if output.Path == "" || output.Path == "stdout" {
		return zapcore.AddSync(os.Stdout), nil
	}

	if output.Path == "stderr" {
		return zapcore.AddSync(os.Stderr), nil
	}

	// 创建日志目录
	dir := filepath.Dir(output.Path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// 使用lumberjack实现日志轮转
	lumberjackLogger := &lumberjack.Logger{
		Filename:   output.Path,
		MaxSize:    output.MaxSize,
		MaxBackups: output.MaxBackup,
		MaxAge:     output.MaxAge,
		Compress:   output.Compress,
	}

	return zapcore.AddSync(lumberjackLogger), nil
}

// 解析级别字符串为zapcore.Level
func parseLevel(levelStr string) (zapcore.Level, error) {
	switch strings.ToLower(levelStr) {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn", "warning":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	case "dpanic":
		return zapcore.DPanicLevel, nil
	case "panic":
		return zapcore.PanicLevel, nil
	case "fatal":
		return zapcore.FatalLevel, nil
	default:
		return zapcore.InfoLevel, errors.Errorf("不支持的日志级别: %s", levelStr)
	}
}

// 获取级别编码器
func getLevelEncoder(encoder string) zapcore.LevelEncoder {
	switch encoder {
	case "capital":
		return zapcore.CapitalLevelEncoder
	case "capitalColor":
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// 获取持续时间编码器
func getDurationEncoder(encoder string) zapcore.DurationEncoder {
	switch encoder {
	case "string":
		return zapcore.StringDurationEncoder
	default:
		return zapcore.SecondsDurationEncoder
	}
}

// 获取调用者编码器
func getCallerEncoder(encoder string) zapcore.CallerEncoder {
	switch encoder {
	case "full":
		return zapcore.FullCallerEncoder
	default:
		return zapcore.ShortCallerEncoder
	}
}

// 获取时间编码器
func getTimeEncoder(layout string) zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(layout))
	}
}
