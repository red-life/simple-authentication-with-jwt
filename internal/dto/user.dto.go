package dto


type LoginDTO struct {
	Email string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,max=20,min=6"`
}

type RegisterDTO struct {
	Firstname string `json:"firstname" binding:"required" validate:"required"`
	Lastname  string `json:"lastname" binding:"required" validate:"required"`
	Email     string `json:"email" binding:"required" validate:"required"`
	Username  string `json:"username" binding:"required" validate:"required"`
	Password  string `json:"password" binding:"required" validate:"required,max=64,min=6"`
}