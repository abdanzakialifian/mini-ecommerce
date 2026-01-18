package user

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=4"`
}

type UpdateUserRequest struct {
	Name        *string `json:"name" binding:"omitempty,min=3,max=50"`
	Email       *string `json:"email" binding:"omitempty,email,max=50"`
	OldPassword *string `json:"old_password" binding:"omitempty,min=4"`
	NewPassword *string `json:"new_password" binding:"omitempty,min=4"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=4"`
}
