package data

import (
	"errors"
	"floweys_app/features/orders"
	"gorm.io/gorm"
)

type OrderQuery struct {
	db *gorm.DB
}

func (repo *OrderQuery) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (repo *OrderQuery) Update(id uint, input orders.OrderCore) error {
	//TODO implement me
	panic("implement me")
}

func New(db *gorm.DB) orders.OrderDataInterface {
	return &OrderQuery{
		db: db,
	}
}
