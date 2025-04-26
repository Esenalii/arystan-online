package service

import (
	"rest-project/internal/models"
	"rest-project/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(name, email, password, role string) (*models.User, error) {
	return s.repo.Create(name, email, password, role)
}

func (s *UserService) UpdateUser(id int, userEdit *models.UserEdit) (*models.User, error) {
	return s.repo.Update(id, userEdit)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
