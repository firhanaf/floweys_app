package data

import (
	"floweys_app/features/orders"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Category string
	Item     string
	Qty      uint
	TaskID   uint
}

func OrderModelToCOre(input Order) orders.OrderCore {
	return orders.OrderCore{
		Category: input.Category,
		Item:     input.Item,
		Qty:      input.Qty,
		TaskID:   input.TaskID,
	}
}

func OrderCoreToModel(input orders.OrderCore) Order {
	return Order{
		Category: input.Category,
		Item:     input.Item,
		Qty:      input.Qty,
		TaskID:   input.TaskID,
	}
}
