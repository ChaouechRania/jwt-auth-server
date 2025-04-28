// internal/services/userService.go
package services

import (
	"jwt-auth-server/internal/api"
	"jwt-auth-server/internal/models"
	"jwt-auth-server/internal/repositories"
)

type UserService interface {
	CreateUser(user *models.User) error
	Authenticate(login *api.Login) (string, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) Authenticate(login *api.Login) (string, error) {
	// implementation for Authenticate
	return s.repo.Authenticate(login)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
