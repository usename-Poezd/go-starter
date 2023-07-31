package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const logPath string = "./logs"

var consoleWriter io.Writer = zerolog.ConsoleWriter{
	Out:        os.Stdout,
	TimeFormat: time.RFC3339,
}

func NewConsole() zerolog.Logger {
	return zerolog.New(consoleWriter).
		With().
		Timestamp().
		Logger()
}

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {

		fileLogger := &lumberjack.Logger{
			Filename: logPath + "/app.log",
			MaxAge:   5,
		}

		output := zerolog.MultiLevelWriter(consoleWriter, fileLogger)

		log = zerolog.New(output).
			With().
			Timestamp().
			Logger()
	})

	return log
}
