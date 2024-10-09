package main

import (
	"bufio"
	"context"
	"os"
	"strings"
	"sync"
	"task-scheduler/constants"
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

	log.Info().Ctx(ctx).Msg("task manager started")

	for {
		// if all tasks are executed then we can exit the command listening
		if dataCenter.AreAllTasksExecuted(ctx) {
			log.Info().Ctx(ctx).Msg("all tasks executed successfully")
			break
		}

		// command for the data center to process
		reader := bufio.NewReader(os.Stdin)
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal().Ctx(ctx).Msgf("error scanning for the command %s", err.Error())
		}

		parts := strings.SplitN(command, " ", 2)
		commandType := parts[0]
		commandArgs := parts[1]

		switch commandType {
		case constants.AddTask:
			// get task from the command args
			task := datacenter.Task{}
			task.FromCommandArgString(ctx, commandArgs)

			// add task to the dc
			dataCenter.AddTask(&task)
			log.Info().Ctx(ctx).Msg("task added to the dc")
		case constants.AddResource:
			resource := datacenter.Resource{}
			resource.FromCommandArgString(ctx, commandArgs)

			// add resource to the dc
			resourceId := dataCenter.AddResource(ctx, &resource)
			log.Info().Ctx(ctx).Msgf("resource added to the dc: %s", resourceId)
		case constants.DeleteResource:
			// delete resource from id
			resourceId := commandArgs

			if ok := dataCenter.DeleteResource(ctx, resourceId); !ok {
				log.Debug().Ctx(ctx).Msgf("error deleting the resource with id %s", resourceId)
			}
			log.Info().Ctx(ctx).Msg("resource deleted successfully.")
		case constants.ListResources:
			tasks := datacenter.Task{}
			log.Info().Ctx(ctx).Interface("tasks", tasks).Msg("tasks available")
		default:
			log.Fatal().Ctx(ctx).Msg("error invalid task")
		}

	}

	wg.Wait()

	log.Info().Ctx(ctx).Interface("logs", dataCenter.ExecutionSummary).Msg("execution summary")
}
