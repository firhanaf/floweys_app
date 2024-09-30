package orders

type OrderCore struct {
	ID       uint
	Category string
	Item     string
	Qty      uint
	TaskID   uint
}
