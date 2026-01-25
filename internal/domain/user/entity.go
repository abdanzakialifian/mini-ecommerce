package user

type Data struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Update struct {
	ID          int
	Name        *string
	Email       *string
	OldPassword *string
	NewPassword *string
}

type Login struct {
	Email    string
	Password string
}
