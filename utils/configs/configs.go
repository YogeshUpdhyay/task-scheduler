package configs

import (
	"context"
	"task-scheduler/constants"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Get(ctx context.Context, configName string) *viper.Viper {
	configClient := viper.New()

	configClient.SetConfigType(constants.ConfigType)
	configClient.AddConfigPath(constants.ConfigDir)
	configClient.SetConfigName(configName)

	if err := configClient.ReadInConfig(); err != nil {
		log.Ctx(ctx).Info().Msg("initializing the application")
	}

	return configClient
}
