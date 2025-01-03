package handler

import (
	"floweys_app/features/orders"
	"floweys_app/features/orders/data"
	"floweys_app/features/tasks"
	"time"
)

type TaskRequest struct {
	Name        string       `json:"name"`
	Address     string       `json:"address"`
	Phone       string       `json:"phone"`
	Orders      []data.Order `json:"orders"`
	CreatedBy   string       `json:"created_by"`
	UpdatedBy   string       `json:"updated_by"`
	Status      string       `json:"status"`
	Description string       `json:"description"`
	Filename    string       `json:"filename"`
}

func TaskRequestToCore(task TaskRequest) tasks.TaskCore {
	return tasks.TaskCore{
		Name:        task.Name,
		Address:     task.Address,
		Phone:       task.Phone,
		Orders:      []orders.OrderCore{},
		CreatedBy:   task.CreatedBy,
		UpdatedBy:   task.UpdatedBy,
		Status:      task.Status,
		Description: task.Description,
		Filename:    task.Filename,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
}
