package data

import (
	"errors"
	"floweys_app/features/orders"
	"gorm.io/gorm"
)

type OrderQuery struct {
	db *gorm.DB
}

func (repo *OrderQuery) GetAll() ([]orders.OrderCore, error) {
	var data []Order
	tx := repo.db.Find(&data)
	if tx.Error != nil {
		return []orders.OrderCore{}, tx.Error
	}
	var orderCore []orders.OrderCore
	for _, value := range data {
		orderCore = append(orderCore, orders.OrderCore{
			ID:       value.ID,
			Category: value.Category,
			Item:     value.Item,
			Qty:      value.Qty,
			TaskID:   value.TaskID,
		})
	}
	return orderCore, nil
}

func (repo *OrderQuery) Delete(id uint) error {
	var data Order
	tx := repo.db.Where(id).Delete(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *OrderQuery) Create(input orders.OrderCore) error {
	var orderModel = OrderCoreToModel(input)

	tx := repo.db.Create(&orderModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *OrderQuery) Get(id uint) (orders.OrderCore, error) {
	var result Order
	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return orders.OrderCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return orders.OrderCore{}, errors.New("data not found")
	}
	resultCore := OrderModelToCOre(result)
	return resultCore, nil
}

func (repo *OrderQuery) Update(id uint, input orders.OrderCore) error {
	var data Order
	tx := repo.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		return tx.Error
	}
	updatedOrder := OrderCoreToModel(input)
	tx = repo.db.Model(&data).Updates(updatedOrder)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) *OrderQuery {
	return &OrderQuery{
		db: db,
	}
}
