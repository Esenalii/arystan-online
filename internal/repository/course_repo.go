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

func (r *CourseRepository) GetByIDWithLessons(id uint) (*models.Course, error) {
	var course models.Course
	if err := r.DB.
		Preload("Lessons").
		First(&course, id).
		Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *CourseRepository) GetLessonsByCourseID(courseID uint) ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := r.DB.
		Where("course_id = ?", courseID).
		Find(&lessons).
		Error; err != nil {
		return nil, err
	}
	return lessons, nil
}

func (r *CourseRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Course{}, id).Error
}
