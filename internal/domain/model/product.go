package model

import "time"

type Product struct {
	ID          string
	CategoryID  string
	Name        string
	Description string
	Price       float64
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
