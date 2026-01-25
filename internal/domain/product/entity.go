package product

type Data struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Price       float64
	Stock       int
}

type Update struct {
	ID          string
	CategoryID  *string
	Name        *string
	Description *string
	Price       *float64
	Stock       *int
}
