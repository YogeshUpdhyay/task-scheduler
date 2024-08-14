package entities

type Resource struct {
	ResourceId   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Price        int    `json:"price"`
	CPUConfig    int    `json:"cpuConfig"`
	IsAllocated  bool   `json:"isAllocated"`
}
