package models

type UserEdit struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role" binding:"required"`
}
