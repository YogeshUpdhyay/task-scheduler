package logger

import (
	"os"
	"task-scheduler/constants"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type TracingHook struct{}

func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	applicationName := ctx.Value(constants.ContextApplicationNameKey).(string)
	e.Str(constants.ContextApplicationNameKey, applicationName)
}

// InitLogger initializes the global logger with the specified level
func InitLogger() {
	logger := zerolog.New(os.Stderr).
		Level(zerolog.DebugLevel).
		With().Caller().Timestamp().Logger()

	logger = logger.Hook(TracingHook{})
	log.Logger = logger
}
