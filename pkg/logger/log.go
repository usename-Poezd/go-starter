package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"sync"
	"time"
)

const logPath string = "./logs"

var consoleWriter io.Writer = zerolog.ConsoleWriter{
	Out:        os.Stdout,
	TimeFormat: time.RFC3339,
}

func NewConsole() zerolog.Logger {
	return buildLogger(true)
}

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		log = buildLogger(false)
	})

	return log
}

func buildLogger(consoleOnly bool) zerolog.Logger {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	output := consoleWriter

	if !consoleOnly {
		fileLogger := &lumberjack.Logger{
			Filename:   logPath + "/app.log",
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		}

		output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
	}

	return zerolog.New(output).
		With().
		Timestamp().
		Logger()
}
