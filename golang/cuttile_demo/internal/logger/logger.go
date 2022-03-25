package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var levelMap = map[int32]zapcore.Level{
	0: zapcore.DebugLevel,
	1: zapcore.InfoLevel,
	2: zapcore.WarnLevel,
	3: zapcore.ErrorLevel,
	4: zapcore.PanicLevel,
	5: zapcore.FatalLevel,
}

func GetLoggerLevel(lvl int32) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func LoggerInit(lvl int32, name string) (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		CallerKey:  "caller",
		MessageKey: "msg",
		//StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
	}

	level := GetLoggerLevel(lvl)
	atom := zap.NewAtomicLevelAt(level)

	config := zap.Config{
		Level:            atom,                                        // 日志级别
		Development:      false,                                       // 开发模式，堆栈跟踪
		Encoding:         "json",                                      // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                               // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": name}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout"},                          // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}

	// 构建日志
	log, err := config.Build()
	if err != nil {
		return nil, err
	}
	return log, nil
}

var Logger *zap.Logger
