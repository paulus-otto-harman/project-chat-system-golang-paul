package domain

type Login struct {
	Email    string `json:"email" binding:"required,email" example:"user1@mail.com"`
	Password string `json:"password" binding:"required,min=5" example:"user1"`
}
