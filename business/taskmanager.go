package business

import "github.com/rs/zerolog/log"

const allocateResourceFuncName = "AllocateResource"

func AllocateResource() error {
	log.Debug().Msgf("%s: allocate resource to task", allocateResourceFuncName)
	log.Debug().Msgf("%s: completed allocating resource to task", allocateResourceFuncName)
	return nil
}
