package datacenter

import (
	"context"
	"sync"
	"time"
)

type DataCenter struct {
	DataCenterId string `json:"dataCenterId"`
	Location     string `json:"location"`

	ResourcesLock sync.RWMutex
	Resources     []*Resource `json:"resources"`

	Tasks            []*Task
	ExecutionSummary []*ExecutionSummaryLog
}

type Resource struct {
	ResourceId   string `json:"resourceId"`
	ResourceType string `json:"resourceType"`
	Price        int    `json:"price"`
	CPUConfig    int    `json:"cpuConfig"`
	IsAllocated  bool   `json:"isAllocated"`
}

type Task struct {
	TaskId       string
	ResourceType string
	CpuConfig    int
}

type ExecutionSummaryLog struct {
	TaskStatus    string
	TaskDuration  int
	TaskStartTime *time.Time
	TaskEndTime   *time.Time
	TaskId        string
}

func (dc *DataCenter) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	dc.ExecutionSummary = append(dc.ExecutionSummary, &ExecutionSummaryLog{TaskStatus: "Completed"})
}
