package entities

type TaskManager struct {
	TaskManagerId string       `json:"taskManagerId"`
	DataCenters   []DataCenter `json:"dataCenters"`
}
