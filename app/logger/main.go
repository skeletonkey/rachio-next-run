package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/skeletonkey/rachio-next-run/app/config"
)

var log *zerolog.Logger
var logCfg *logger
var lock = &sync.Mutex{}

func (l *logger) Initialize() {
	lock.Lock()
	defer lock.Unlock()

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano

	logLevel, err := zerolog.ParseLevel(l.LogLevel)
	if err != nil {
		panic(fmt.Errorf("configured loglevel (%s) is not recognized by zerolog.ParseLevel", l.LogLevel))
	}

	outputs := make([]io.Writer, 0)

	if l.LogToStdout {
		outputs = append(outputs, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		})
	}

	if l.LogToFile {
		outputs = append(outputs, &lumberjack.Logger{
			Filename:   l.LogFile,
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		})
	}

	if l.LogToStderr {
		outputs = append(outputs, os.Stderr)
	}

	var output io.Writer = zerolog.MultiLevelWriter(outputs...)

	tempLog := zerolog.New(output).
		Level(logLevel).
		With().
		Timestamp().
		Logger()
	// TODO: This doesn't work - it's setting log to a knew memory location
	//   instead of putting the new zerlog into the existing memory location
	log = &tempLog
}

var once sync.Once

// Get a reference to the zerolog.Logger with the appropriate configured settings.
func Get() *zerolog.Logger {
	logCfg = getConfig()
	once.Do(func() {
		config.RegisterInitializer("logger", logCfg)
	})
	return log
}
