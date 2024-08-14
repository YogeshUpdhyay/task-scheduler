package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"task-scheduler/constants"
	"testing"
)

var dc = DataCenter{
	DataCenterId: uuid.NewString(),
	Location:     "ap-south-1",
}

func TestDataCenter_AddResource(t *testing.T) {
	dc.AddResource(&Resource{
		ResourceId:   uuid.NewString(),
		ResourceType: constants.ServerInstanceResourceType,
		Price:        1000,
		CPUConfig:    16,
	})

	assert.Equal(t, 1, len(dc.Resources), "one resource should be present in the data center now.")
}

func TestDataCenter_RemoveResource(t *testing.T) {
	resourceId := uuid.NewString()

	dc.AddResource(&Resource{
		ResourceId:   resourceId,
		ResourceType: constants.ServerInstanceResourceType,
		Price:        1000,
		CPUConfig:    16,
	})

	assert.Equal(t, 1, len(dc.Resources), "one resource should be present in the data center now.")

	err := dc.RemoveResource(resourceId)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(dc.Resources), "no resource should be present.")
}
