package data

import (
	"floweys_app/features/orders"
	"floweys_app/features/orders/data"
	"floweys_app/features/tasks"
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Name        string
	Address     string
	Phone       string
	Orders      []data.Order
	CreatedBy   string
	UpdatedBy   string
	Status      string
	Description string
	Filename    string
}

func TaskModelToCore(input Task) tasks.TaskCore {
	return tasks.TaskCore{
		Name:        input.Name,
		Address:     input.Address,
		Phone:       input.Phone,
		Orders:      []orders.OrderCore{},
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
		Status:      input.Status,
		Description: input.Description,
		Filename:    input.Filename,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   time.Time{},
	}
}

func TaskCoreToModel(input tasks.TaskCore) Task {
	return Task{
		Model:       gorm.Model{},
		Name:        input.Name,
		Address:     input.Address,
		Phone:       input.Phone,
		Orders:      []data.Order{},
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
		Status:      input.Status,
		Description: input.Description,
		Filename:    input.Filename,
	}
}
