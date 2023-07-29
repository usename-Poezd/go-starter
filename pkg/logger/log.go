package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logPath string = "./logs"

func NewConsole() *zap.Logger {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	core := zapcore.NewCore(
		// use NewConsoleEncoder for human readable output
		zapcore.NewConsoleEncoder(cfg),

		// write to stdout as well as log files
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)
	logger := zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync() // flushes buffer, if any

	return logger
}

func New() *zap.SugaredLogger {

	fileLogger := &lumberjack.Logger{
		Filename: logPath + "/app.log",
		MaxAge:   5,
	}

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	core := zapcore.NewTee(
		zapcore.NewCore(
			// use NewConsoleEncoder for human readable output
			zapcore.NewJSONEncoder(cfg),

			// write to stdout as well as log files
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(fileLogger)),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		),
		zapcore.NewCore(
			// use NewConsoleEncoder for human readable output
			zapcore.NewConsoleEncoder(cfg),

			// write to stdout as well as log files
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		),
	)

	logger := zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	return sugar
}
