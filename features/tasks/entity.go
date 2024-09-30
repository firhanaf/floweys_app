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
	Delete(input TaskCore) error
	GetAll() ([]TaskCore, error)
}

type TaskServiceInterface interface {
	Add(input TaskCore) error
	Edit(input TaskCore) error
	Remove(input TaskCore) error
	ReadAll() ([]TaskCore, error)
}
