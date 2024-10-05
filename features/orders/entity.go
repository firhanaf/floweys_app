package orders

type OrderCore struct {
	ID       uint
	Category string
	Item     string
	Qty      uint
	TaskID   uint
}

type OrderDataInterface interface {
	Create(input OrderCore) error
	Get(id uint) (OrderCore, error)
	Update(id uint, input OrderCore) error
	Delete(id uint) error
}

type OrderServiceInterface interface {
	Add(input OrderCore) error
	Read(id uint) (OrderCore, error)
	Edit(id uint, input OrderCore) error
	Remove(id uint) error
}
