package main

import (
	"context"
	"task-scheduler/business"
	"task-scheduler/constants"
	"task-scheduler/entities"
	"task-scheduler/utils/configs"
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

	log.Info().Ctx(ctx).Msg("intializing task manager")

	// fetching the avialable resources from the config
	availableResources, err := getAvailableResource(ctx)
	if err != nil {
		log.Error().Ctx(ctx).Stack().Err(err).Msg("error getting available resources")
		return
	}

	// initializing task manager
	taskManager, err := business.InitializeTaskManager(ctx, availableResources)
	if err != nil {
		log.Info().Ctx(ctx).Msg("error initializing the task manager")
		return
	}

	log.Info().Ctx(ctx).Msg("task manager initialization complete")

	// prepare tasks based on the users input

	// assign tasks to the task manager
	taskManager.AddTask(&entities.Task{TaskID: "1", RequestedConfiguration: entities.TaskConfiguration{}})
}

// fetches avaiable resources from config
func getAvailableResource(ctx context.Context) ([]*entities.Resource, error) {
	availableResources := []*entities.Resource{}
	resources := configs.Get(ctx, constants.ApplicationConfig).Get(constants.DataCenterResourcesKey).([]interface{})

	for _, resource := range resources {
		availableResources = append(availableResources, &entities.Resource{
			ResourceId:   resource.(map[string]interface{})[constants.ResourceIdKey].(string),
			ResourceType: resource.(map[string]interface{})[constants.ResourceTypeKey].(string),
			Price:        resource.(map[string]interface{})[constants.ResourcePriceKey].(int),
			CPUConfig:    resource.(map[string]interface{})[constants.ResourceCPUConfigKey].(int),
		})
	}

	return availableResources, nil
}
