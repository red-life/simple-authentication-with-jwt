package dto


type LoginDTO struct {
	email string `json:"email" binding:"required" validate:"required,email"`
	password string `json:"password" binding:"required" validate:"required,max=20,min=6"`
}