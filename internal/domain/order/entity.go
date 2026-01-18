package order

type Status string

const (
	StatusPending   Status = "pending"
	StatusPaid      Status = "paid"
	StatusCancelled Status = "cancelled"
)

type Order struct {
	ID         int
	UserID     int
	TotalPrice float64
	Status     Status
}

type UpdateOrder struct {
	ID         int
	TotalPrice *float64
	Status     *Status
}

type OrderItem struct {
	ID        int
	OrderID   int
	ProductID string
	Price     float64
	Quantity  int
}
