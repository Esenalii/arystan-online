package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type CourseRepository struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) GetAll() ([]models.Course, error) {
	var courses []models.Course
	err := r.DB.Find(&courses).Error
	return courses, err
}

func (r *CourseRepository) Create(course *models.Course) error {
	return r.DB.Create(course).Error
}
