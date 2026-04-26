package models

type Task struct {
	ID             int    `json:"id"`
	TaskName       string `json:"TaskName"`
	PriorityOfTask string `json:"priorityOfTask"`
	Status         string `json:"Status"`
}

type New_Task struct {
	OldTaskName string `json:"OldTaskName"`
	NewTaskName string `json:"NewTaskName"`
}

type Update_Status struct {
	TaskName string `json:"TaskName"`
	Status   string `json:"Status"`
}
