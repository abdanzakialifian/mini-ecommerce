package cart

type Data struct {
	ID     int
	UserID int
}

type Item struct {
	ID        int
	CartID    int
	ProductID string
	Quantity  int
}

type UpdateItem struct {
	ID       int
	Quantity int
}
