package handler

import "floweys_app/features/orders"

type OrderRequest struct {
	Category string `json:"category"`
	Item     string `json:"item"`
	Qty      uint   `json:"qty"`
	TaskID   uint   `json:"taskID"`
}

func OrderRequestToCore(order OrderRequest) orders.OrderCore {
	return orders.OrderCore{
		Category: order.Category,
		Item:     order.Item,
		Qty:      order.Qty,
		TaskID:   order.TaskID,
	}
}
