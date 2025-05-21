package models

type Course struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Lessons     []Lesson `json:"lessons" gorm:"foreignKey:CourseID"`
}
