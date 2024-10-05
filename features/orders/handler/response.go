package handler

type OrderResponse struct {
	ID       uint   `json:"id"`
	Category string `json:"category"`
	Item     string `json:"item"`
	Qty      uint   `json:"qty"`
	TaskID   uint   `json:"taskID"`
}
