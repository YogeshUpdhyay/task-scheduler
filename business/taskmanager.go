package business

import (
	"task-scheduler/entities"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

const allocateResourceFuncName = "AllocateResourceToTask"

func AllocateResourceToTask() error {
	log.Debug().Msgf("%s: allocate resource to task", allocateResourceFuncName)
	log.Debug().Msgf("%s: completed allocating resource to task", allocateResourceFuncName)
	return nil
}

func InitializeTaskManager() error {
	dataCenterId := uuid.NewString()

	dc := entities.DataCenter{
		DataCenterId: dataCenterId,
		Location:     "ap-south-1",
		Resources:    []entities.Resource{},
	}

	return nil
}
