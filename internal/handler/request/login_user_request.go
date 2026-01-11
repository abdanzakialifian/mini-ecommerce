package request

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=4"`
}
