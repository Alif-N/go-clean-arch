package service

import (
	"fmt"
	"go-clean-architecture/internal/models"
	"go-clean-architecture/internal/repository"
)

type UserService interface {
	GetAll() ([]models.User, error)
	Create(name string) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) Create(name string) (models.User, error) {
	if name == "" {
		return models.User{}, fmt.Errorf("user name cannot be empty")
	}

	newUser := models.User{
		Name: name,
	}

	return s.repo.Create(newUser)
}