package product

type Product struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Price       float64
	Stock       int
}

type UpdateProduct struct {
	ID          string
	CategoryID  *string
	Name        *string
	Description *string
	Price       *float64
	Stock       *int
}
