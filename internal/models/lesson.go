package models

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Video    string `json:"video"`               // YouTube сілтемесі
	CourseID uint   `json:"course_id"`           // Қай курсқа тиесілі
	Course   Course `gorm:"foreignKey:CourseID"` // Курспен байланыс
}
