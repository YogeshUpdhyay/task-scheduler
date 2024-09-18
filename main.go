package main

import (
	"context"
	appConfig "task-scheduler/utils/context"
	"task-scheduler/utils/logger"

	"github.com/rs/zerolog/log"
)

func main() {
	// intialize logger
	logger.InitLogger()

	// prepare application config
	ctx := context.Background()
	ctx = appConfig.IntializeContext(ctx)

	log.Debug().Ctx(ctx).Msg("hi")
}
