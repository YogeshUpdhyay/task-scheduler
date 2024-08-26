package business

import (
	"task-scheduler/entities"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const allocateResourceFuncName = "AllocateResourceToTask"
const initializeTaskManagerFuncName = "InitializeTaskManager"

func AllocateResourceToTask() error {
	log.Debug().Msgf("%s: allocate resource to task", allocateResourceFuncName)
	log.Debug().Msgf("%s: completed allocating resource to task", allocateResourceFuncName)
	return nil
}

func InitializeTaskManager(resources []*entities.Resource) (*entities.DataCenter, error) {
	// creating datacenter
	log.Debug().Msgf("%s: creating one dc for the tasks", initializeTaskManagerFuncName)
	dataCenterId := uuid.NewString()
	dc := entities.DataCenter{
		DataCenterId: dataCenterId,
		Location:     "ap-south-1",
		Resources:    []entities.Resource{},
	}
	log.Debug().Msgf("%s: dc creation complete", initializeTaskManagerFuncName)

	// adding resources to data center
	log.Debug().Msgf("%s: adding resources to the dc", initializeTaskManagerFuncName)
	for _, resource := range resources {
		dc.AddResource(resource)
	}
	log.Debug().Msgf("%s: resoruce addition completed", initializeTaskManagerFuncName)

	return &dc, nil
}
