package main

import (
	"context"
	"task-scheduler/constants"
	"task-scheduler/utils/configs"
	"task-scheduler/utils/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	logger.InitLogger()
	ctx := context.Background()

	applicationName := configs.Get(ctx, constants.ApplicationConfig).GetString(constants.ApplicationNameKey)
	ctx = context.WithValue(ctx, constants.ContextApplicationNameKey, applicationName)
	log.Debug().Ctx(ctx).Msg("hi")
}
