package dto

type UserCreateDTO struct {
	Username string `json:"userName" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
