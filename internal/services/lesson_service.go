package service

import (
	"rest-project/internal/models"
	"rest-project/internal/repository"
)

type LessonService struct {
	repo *repository.LessonRepository
}

func NewLessonService(repo *repository.LessonRepository) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) GetAllLessons() ([]models.Lesson, error) {
	return s.repo.GetAll()
}

func (s *LessonService) GetLessonByID(id uint) (*models.Lesson, error) {
	return s.repo.GetByID(id)
}

func (s *LessonService) GetLessonsByCourseID(courseID uint) ([]models.Lesson, error) {
	return s.repo.GetByCourseID(courseID)
}

func (s *LessonService) CreateLesson(lesson *models.Lesson) error {
	return s.repo.Create(lesson)
}

func (s *LessonService) UpdateLesson(lesson *models.Lesson) error {
	return s.repo.Update(lesson)
}

func (s *LessonService) DeleteLesson(id uint) error {
	return s.repo.Delete(id)
}
