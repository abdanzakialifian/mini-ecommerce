package model

type UpdateUser struct {
	ID          int
	Name        *string
	Email       *string
	OldPassword *string
	NewPassword *string
}
