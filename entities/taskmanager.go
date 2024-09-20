package entities

// there can be multiple data centers
// TODO we can implement multiple dcs for the different AZ and Region logic later

// task manager
type TaskManager struct {
	DataCenter *DataCenter
	Tasks      []*Task
}

// add tasks to the task manager
func (tm *TaskManager) AddTask(task *Task) {
	tm.Tasks = append(tm.Tasks, task)
}

// allocate a available resource of the data center to the task manager
func (tm *TaskManager) AllocateResourceToTask(task *Task) *Resource {
	return nil
}
