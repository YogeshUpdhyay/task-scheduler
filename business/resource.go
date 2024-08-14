package business

import (
	"task-scheduler/entities"

	"github.com/rs/zerolog/log"
)

const addFuncName = "AddResourceToDataCenter"
const deleteFuncName = "RemoveResourceFromDataCenter"
const listFuncName = "ListResources"

func AddResourceToDataCenter(_ *entities.Resource) error {
	log.Debug().Msgf("%s: adding resource to the datacenters", addFuncName)
	log.Debug().Msgf("%s: completed adding resource to datacenters", addFuncName)
	return nil
}

func RemoveResourceFromDataCenter(_ string) error {
	log.Debug().Msgf("%s: removing resource from the datacenter", deleteFuncName)
	log.Debug().Msgf("%s: completed removing resource from datacenter", deleteFuncName)
	return nil
}

func ListResources() ([]*entities.Resource, error) {
	log.Debug().Msgf("%s: list resources", listFuncName)
	log.Debug().Msgf("%s: list resource completed", listFuncName)
	return nil, nil
}
