package handler

import "floweys_app/features/orders"

type OrderResponse struct {
	ID       uint   `json:"id"`
	Category string `json:"category"`
	Item     string `json:"item"`
	Qty      uint   `json:"qty"`
	TaskID   uint   `json:"taskID"`
}

func OrderCoreToResponse(input orders.OrderCore) OrderResponse {
	return OrderResponse{
		ID:       input.ID,
		Category: input.Category,
		Item:     input.Item,
		Qty:      input.Qty,
		TaskID:   input.TaskID,
	}
}
