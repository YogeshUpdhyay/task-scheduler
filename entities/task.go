package entities

type Task struct {
	TaskID                 string            `json:"taskId"`
	RequestedConfiguration TaskConfiguration `json:"taskConfiguration"`
}

type TaskConfiguration struct {
	ResourceType string `json:"resourceType"`
	MinCpu       int    `json:"minCpu"`
}
