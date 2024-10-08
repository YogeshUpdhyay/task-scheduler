package datacenter

import (
	"context"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type DataCenter struct {
	DataCenterId string `json:"dataCenterId"`
	Location     string `json:"location"`

	ResourcesLock sync.RWMutex
	Resources     []*Resource `json:"resources"`

	TasksLock        sync.RWMutex
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

func (r *Resource) FromCommandArgString(ctx context.Context, commandString string) {
	parts := strings.Split(commandString, " ")

	price, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal().Ctx(ctx).Msg("error parsing price to int")
	}

	cpuConfig, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal().Ctx(ctx).Msg("error parsing cpu config to int")
	}

	r.ResourceType, r.Price, r.CPUConfig, r.IsAllocated = parts[0], price, cpuConfig, false
}

type Task struct {
	TaskId       string
	ResourceType string
	CpuConfig    int
}

// creating a task from the command string
func (t *Task) FromCommandArgString(ctx context.Context, commandString string) {
	parts := strings.Split(commandString, " ")

	// parse cpu config
	cpuConfig, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal().Ctx(ctx).Msg("error invalid cpu config")
	}
	t.TaskId, t.ResourceType, t.CpuConfig = parts[0], parts[1], cpuConfig
}

type ExecutionSummaryLog struct {
	TaskStatus    string
	TaskDuration  int
	TaskStartTime *time.Time
	TaskEndTime   *time.Time
	TaskId        string
}

// add a resource to the data center
func (dc *DataCenter) AddResource(ctx context.Context, resource *Resource) string {
	// generate an id for the resource and add that resource to the data center
	resourceId := uuid.NewString()
	resource.ResourceId = resourceId

	dc.ResourcesLock.Lock()
	dc.Resources = append(dc.Resources, resource)
	dc.ResourcesLock.Unlock()

	log.Debug().Ctx(ctx).Msg("resource added successfully")

	return resourceId
}

// delete a resource from the datacenter by id
func (dc *DataCenter) DeleteResource(ctx context.Context, resourceId string) bool {
	// find the resource using the resource id and delete from the resource
	// the resource should be deleted only if it is not allocated to the task
	// if allocated then wait till the resource becomes available and the delete it

	dc.ResourcesLock.Lock()
	// find if the resource with resource id exists or not and if it can be deleted or not
	for i := 0; i < len(dc.Resources); i++ {
		if dc.Resources[i].ResourceId == resourceId {
			log.Debug().Ctx(ctx).Msg("resource found for deletion")

			if dc.Resources[i].IsAllocated {
				log.Debug().Ctx(ctx).Msg("the resource can not be deleted at the moment")
				return false
			}

			dc.Resources = append(dc.Resources[0:i], dc.Resources[i+1:]...)
			log.Debug().Ctx(ctx).Msg("resource deleted successfully")
			return true
		}
	}
	dc.ResourcesLock.Unlock()

	// no resource found with this resource so default to true
	log.Debug().Ctx(ctx).Msg("no resource found for deletion")

	return true

}

func (dc *DataCenter) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	dc.ExecutionSummary = append(dc.ExecutionSummary, &ExecutionSummaryLog{TaskStatus: "Completed"})
}

func (dc *DataCenter) AreAllTasksExecuted(ctx context.Context) bool {
	// checks if all the tasks are executed
	return false
}

func (dc *DataCenter) AddTask(task *Task) {
	// get lock on the tasks and update the task list
	dc.TasksLock.Lock()
	dc.Tasks = append(dc.Tasks, task)
	dc.TasksLock.Unlock()
}
