package dto

type UserDTO struct {
	Name  string `json:"name" binding:"required"`
	Phone string `json:"phone"`
}
