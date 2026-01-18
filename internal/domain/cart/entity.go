package cart

type Cart struct {
	ID     int
	UserID int
}

type CartItem struct {
	ID        int
	CartID    int
	ProductID string
	Quantity  int
}

type UpdateCartItem struct {
	ID       int
	Quantity int
}
