package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email" gorm:"unique"`
}
type UserRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
