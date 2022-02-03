package dto


type LoginDTO struct {
	Email string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,max=20,min=6"`
}