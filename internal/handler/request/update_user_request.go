package request

type UpdateUserRequest struct {
	ID       int     `json:"id" binding:"required,omitempty"`
	Name     *string `json:"name" binding:"omitempty,min=3,max=50"`
	Email    *string `json:"email" binding:"omitempty,email,max=50"`
	Password *string `json:"password" binding:"omitempty,min=4"`
}
