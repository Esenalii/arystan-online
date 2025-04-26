package repository

import (
	"gorm.io/gorm"
	"rest-project/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(name, email, password, role string) (*models.User, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}
	err := r.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id int, userEdit *models.UserEdit) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	user.Name = userEdit.Name
	user.Email = userEdit.Email
	user.Role = userEdit.Role

	if userEdit.Password != "" {
		user.Password = userEdit.Password
	}

	err = r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) DeleteUser(id int) error {
	return r.DB.Delete(&models.User{}, id).Error
}
