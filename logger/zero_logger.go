package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type _ZeroLogger struct {
	logger zerolog.Logger
}

const defaultLogLevel = zerolog.DebugLevel

func NewZeroLogger(appName string, c LogConfig) ILogger {
	log.Logger = log.With().
		Str("module", appName).
		Logger()

	if c.IsPretty {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	// Set correct log level
	if l := c.Level; l != "" {
		level, err := zerolog.ParseLevel(l)
		if err != nil {
			fmt.Printf("invalid log level '%s': %s\n", l, err.Error())
			os.Exit(1)
		}
		log.Info().Msgf("Set global log level to %s", level)
		zerolog.SetGlobalLevel(level)
	} else {
		log.Info().Msgf("Global log level set to default: %s", defaultLogLevel)
		zerolog.SetGlobalLevel(defaultLogLevel)
	}

	return &_ZeroLogger{
		logger: log.Logger,
	}
}

func (l *_ZeroLogger) With(k, v string) ILogger {
	return &_ZeroLogger{
		logger: l.logger.With().Str(k, v).Logger(),
	}
}

func (l *_ZeroLogger) Debugf(f string, v ...interface{}) {
	l.logger.Debug().Msgf(f, v...)
}

func (l *_ZeroLogger) Infof(f string, v ...interface{}) {
	l.logger.Info().Msgf(f, v...)
}

func (l *_ZeroLogger) Warnf(f string, v ...interface{}) {
	l.logger.Warn().Msgf(f, v...)
}

func (l *_ZeroLogger) Errorf(f string, v ...interface{}) {
	l.logger.Error().Msgf(f, v...)
}
