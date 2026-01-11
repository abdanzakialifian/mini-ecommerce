package request

type UpdateUserRequest struct {
	Name        *string `json:"name" binding:"omitempty,min=3,max=50"`
	Email       *string `json:"email" binding:"omitempty,email,max=50"`
	OldPassword *string `json:"old_password" binding:"omitempty,min=4"`
	NewPassword *string `json:"new_password" binding:"omitempty,min=4"`
}
