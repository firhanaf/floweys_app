package service

import (
	"errors"
	"floweys_app/features/orders"
)

type OrderService struct {
	orderData orders.OrderDataInterface
}

func (service *OrderService) Remove(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (service *OrderService) Add(input orders.OrderCore) error {
	if len(input.Category) < 3 {
		return errors.New("category is too short, min 3 characters")
	}
	if len(input.Item) < 3 {
		return errors.New("item is too short, min 3 characters")
	}
	errInsert := service.orderData.Create(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (service *OrderService) Read(id uint) (orders.OrderCore, error) {
	//TODO implement me
	panic("implement me")
}

func (service *OrderService) Edit(id uint, input orders.OrderCore) error {
	//TODO implement me
	panic("implement me")
}

func New(repo orders.OrderDataInterface) orders.OrderServiceInterface {
	return &OrderService{
		orderData: repo,
	}
}
