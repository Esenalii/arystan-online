package service

import (
	"rest-project/internal/models"
	"rest-project/internal/repository"
)

type CourseService struct {
	Repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{Repo: repo}
}

func (s *CourseService) GetAllCourses() ([]models.Course, error) {
	return s.Repo.GetAll()
}

func (s *CourseService) CreateCourse(course *models.Course) error {
	return s.Repo.Create(course)
}
