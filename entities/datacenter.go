package entities

import (
	"fmt"
)

type DataCenter struct {
	DataCenterId string     `json:"dataCenterId"`
	Resources    []Resource `json:"resources"`
	Location     string     `json:"location"`
}

func (dc *DataCenter) AddResource(resource *Resource) {
	// add resources to the data center
	dc.Resources = append(dc.Resources, *resource)
}

func removeResourceById(resources []Resource, resourceId string) ([]Resource, error) {
	// returns the filtered slice of resources by removing the resource with specific id
	for i, resource := range resources {
		if resource.ResourceId == resourceId {
			return append(resources[:i], resources[i+1:]...), nil
		}
	}
	return nil, fmt.Errorf("no resource find with id %s", resourceId)
}

func (dc *DataCenter) RemoveResource(resourceId string) error {
	// remove a resource from the data center
	filteredResources, err := removeResourceById(dc.Resources, resourceId)
	if err != nil {
		return err
	}

	dc.Resources = filteredResources

	return nil
}

func (dc *DataCenter) GetResources(resourceType *string, minCpu *int, isAllocated *bool) []Resource {
	filteredResources := []Resource{}

	for _, resource := range dc.Resources {
		resourceFitsFilter := false

		if resourceType != nil && resource.ResourceType == *resourceType {
			resourceFitsFilter = true
		}

		if minCpu != nil && resource.CPUConfig >= *minCpu {
			resourceFitsFilter = true
		}

		if isAllocated != nil && resource.IsAllocated {
			resourceFitsFilter = true
		}

		if resourceFitsFilter {
			filteredResources = append(filteredResources, resource)
		}
	}

	return filteredResources
}
