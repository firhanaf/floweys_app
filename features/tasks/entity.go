package tasks

import (
	"floweys_app/features/orders"
	"time"
)

type TaskCore struct {
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

type TaskDataInterface interface {
	Create(input TaskCore) error
	Update(input TaskCore) error
	Delete(id uint) error
	GetAll() ([]TaskCore, error)
}

type TaskServiceInterface interface {
	Add(input TaskCore) error
	Edit(id uint, input TaskCore) error
	Remove(id uint) error
	ReadAll() ([]TaskCore, error)
}
