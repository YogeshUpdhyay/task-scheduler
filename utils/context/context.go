package context

import (
	"context"
	"task-scheduler/constants"
	"task-scheduler/utils/configs"
)

func IntializeContext(ctx context.Context) context.Context {
	// setting application name in config
	applicationName := configs.Get(ctx, constants.ApplicationConfig).GetString(constants.ApplicationNameKey)
	ctx = context.WithValue(ctx, constants.ContextApplicationNameKey, applicationName)

	return ctx
}
