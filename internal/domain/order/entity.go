package order

type Status string

const (
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusCancelled Status = "cancelled"
)

type Data struct {
	ID         int
	UserID     int
	TotalPrice float64
	Status     Status
}

type Update struct {
	ID         int
	TotalPrice *float64
	Status     *Status
}

type Item struct {
	ID        int
	OrderID   int
	ProductID string
	Price     float64
	Quantity  int
}

type Detail struct {
	Data  Data
	Items []Item
}

type NewItem struct {
	ProductID string
	Quantity  int
}
