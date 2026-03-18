package repository

import "go-clean-architecture/internal/models"

type UserRepository interface {
	GetAll() ([]models.User, error)
	Create(user models.User) (models.User, error)
}