package business

import (
	"context"
	"task-scheduler/entities"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const initializeTaskManagerFuncName = "InitializeTaskManager"

func InitializeTaskManager(ctx context.Context, resources []*entities.Resource) (*entities.TaskManager, error) {
	// creating datacenter
	log.Debug().Ctx(ctx).Msgf("%s: creating one dc for the tasks", initializeTaskManagerFuncName)
	dataCenterId := uuid.NewString()
	dc := entities.DataCenter{
		DataCenterId: dataCenterId,
		Location:     "ap-south-1",
		Resources:    []entities.Resource{},
	}
	log.Debug().Ctx(ctx).Msgf("%s: dc creation complete", initializeTaskManagerFuncName)

	// adding resources to data center
	log.Debug().Ctx(ctx).Msgf("%s: adding resources to the dc", initializeTaskManagerFuncName)
	for _, resource := range resources {
		dc.AddResource(resource)
	}
	log.Debug().Ctx(ctx).Msgf("%s: resoruce addition completed", initializeTaskManagerFuncName)

	return &entities.TaskManager{DataCenter: &dc, Tasks: []*entities.Task{}}, nil
}
