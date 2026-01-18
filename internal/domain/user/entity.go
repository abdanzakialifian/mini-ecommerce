package user

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type UpdateUser struct {
	ID          int
	Name        *string
	Email       *string
	OldPassword *string
	NewPassword *string
}

type LoginUser struct {
	Email    string
	Password string
}
