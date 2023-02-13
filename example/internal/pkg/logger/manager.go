package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Manager struct {
	logger *zap.Logger
}

func New() *Manager {
	hook := lumberjack.Logger{
		Filename:   C.Output, // 日志文件路径
		MaxSize:    10,       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 7,        // 日志文件最多保存多少个备份
		MaxAge:     10,       // 文件最多保存多少天
		Compress:   true,     // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	var writes []zapcore.WriteSyncer
	if C.Output == "" {
		writes = []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}
	} else {
		writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	}

	// 新建一个ZapCore
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),  // 编码器配置
		zapcore.NewMultiWriteSyncer(writes...), // 打印到控制台和文件
		atomicLevel,                            // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("appName", C.ProjectName))
	// 构造日志
	// logger = zap.New(core, caller, development)
	logger := zap.New(core, caller, development, filed)

	logger.Info("success to create logger")

	return &Manager{
		logger: logger,
	}
}
