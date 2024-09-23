package main

import (
	"context"
	"fmt"
	"sync"
	"task-scheduler/internal/datacenter"
	appConfig "task-scheduler/utils/context"
	"task-scheduler/utils/logger"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func main() {
	// intialize logger
	logger.InitLogger()

	// prepare application config
	ctx := context.Background()
	ctx = appConfig.IntializeContext(ctx)

	// bulding the data center
	// TODO:
	// to build the data center we require the data center which will house the resources for the process executions
	// and the tasks that are to be executed on the data center

	// building the data center
	dataCenter := datacenter.DataCenter{
		DataCenterId: uuid.NewString(),
		Location:     "ap-south-1",
		Resources:    []*datacenter.Resource{},
		Tasks:        []*datacenter.Task{},
	}

	// starting the data center
	// this will make the data center start listening for the commands and processing the tasks
	// we will listening for the commands on the main thread and process the tasks on a different thread

	var wg sync.WaitGroup

	// starting the task processing of the data center
	wg.Add(1)
	go dataCenter.Start(ctx, &wg)

	for {
		// if all tasks are executed then we can exit the command listening
		if dataCenter.AreAllTasksExecuted(ctx) {
			break
		}

		var command string
		_, err := fmt.Scanln(&command)

		if err != nil {
			log.Fatal().Ctx(ctx).Msg("error scanning for the command")
		}

	}

	wg.Wait()
	log.Info().Ctx(ctx).Interface("logs", dataCenter.ExecutionSummary).Msg("execution summary")
}
