package main

import (
	"context"
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
		DataCenterId:     uuid.NewString(),
		Location:         "ap-south-1",
		Resources:        []*datacenter.Resource{},
		Tasks:            []*datacenter.Task{},
		ExecutionSummary: []*datacenter.ExecutionSummaryLog{},
	}

	// starting the data center this will make the data center start listening for the tasks
	//  we will listening for the tasks on a different thread and the commands for adding and removing
	// resource on a different thread hence we are using a wait group here to run them

	var wg sync.WaitGroup

	// the data center runs till the po
	wg.Add(1)
	go dataCenter.Start(ctx, &wg)
	wg.Wait()

	log.Info().Ctx(ctx).Interface("logs", dataCenter.ExecutionSummary).Msg("execution summary")
}
