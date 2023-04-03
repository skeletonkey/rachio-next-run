package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var once sync.Once

var log zerolog.Logger

func Get() zerolog.Logger {
	once.Do(func() {
		config := getConfig()
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := zerolog.ParseLevel(config.LogLevel)
		if err != nil {
			panic(fmt.Errorf("configured loglevel (%s) is not recognized by zerolog.ParseLevel", config.LogLevel))
		}

		outputs := make([]io.Writer, 0)

		if config.LogToStdout {
			outputs = append(outputs, zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: time.RFC3339,
			})
		}

		if config.LogToFile {
			outputs = append(outputs, &lumberjack.Logger{
				Filename:   config.LogFile,
				MaxSize:    5,
				MaxBackups: 10,
				MaxAge:     14,
				Compress:   true,
			})
		}

		if config.LogToStderr {
			outputs = append(outputs, os.Stderr)
		}

		var output io.Writer = zerolog.MultiLevelWriter(outputs...)

		var gitRevision string

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			for _, v := range buildInfo.Settings {
				if v.Key == "vcs.revision" {
					gitRevision = v.Value
					break
				}
			}
		}

		log = zerolog.New(output).
			Level(logLevel).
			With().
			Timestamp().
			Str("git_revision", gitRevision).
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return log
}
