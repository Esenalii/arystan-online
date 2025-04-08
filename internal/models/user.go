// internal/model/user.go

package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role"` // "admin", "teacher", "student"
	CreatedAt time.Time
	UpdatedAt time.Time
}
