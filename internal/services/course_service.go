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

func (s *CourseService) GetCourseWithLessons(id uint) (*models.Course, error) {
	return s.Repo.GetByIDWithLessons(id)
}

func (s *CourseService) GetLessonsByCourseID(courseID uint) ([]models.Lesson, error) {
	return s.Repo.GetLessonsByCourseID(courseID)
}

func (s *CourseService) DeleteCourse(id uint) error {
	return s.Repo.Delete(id)
}
