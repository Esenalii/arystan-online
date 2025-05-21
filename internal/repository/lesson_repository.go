package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type LessonRepository struct {
	DB *gorm.DB
}

func NewLessonRepository(db *gorm.DB) *LessonRepository {
	return &LessonRepository{DB: db}
}

func (r *LessonRepository) GetAll() ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := r.DB.Find(&lessons).Error
	return lessons, err
}

func (r *LessonRepository) GetByID(id uint) (*models.Lesson, error) {
	var lesson models.Lesson
	err := r.DB.First(&lesson, id).Error
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (r *LessonRepository) GetByCourseID(courseID uint) ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := r.DB.Where("course_id = ?", courseID).Find(&lessons).Error
	return lessons, err
}

func (r *LessonRepository) Create(lesson *models.Lesson) error {
	return r.DB.Create(lesson).Error
}

func (r *LessonRepository) Update(lesson *models.Lesson) error {
	return r.DB.Save(lesson).Error
}

func (r *LessonRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Lesson{}, id).Error
}
