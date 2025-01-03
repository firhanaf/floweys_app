package handler

import (
	"floweys_app/features/orders"
	"floweys_app/features/tasks"
	"time"
)

type TaskResponse struct {
	ID          uint
	Name        string
	Address     string
	Phone       string
	Orders      []orders.OrderCore
	CreatedBy   string
	UpdatedBy   string
	Status      string
	Description string
	Filename    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func TaskCoreToResponse(input tasks.TaskCore) TaskResponse {
	return TaskResponse{
		ID:          input.ID,
		Name:        input.Name,
		Address:     input.Address,
		Phone:       input.Phone,
		Orders:      input.Orders,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
		Status:      input.Status,
		Description: input.Description,
		Filename:    input.Filename,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		DeletedAt:   input.DeletedAt,
	}
}
