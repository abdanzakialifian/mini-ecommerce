package category

type CreateRequest struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
}

type UpdateRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,min=3,max=50"`
}
