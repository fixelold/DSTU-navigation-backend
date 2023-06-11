package user

import "navigation/internal/models"

type Repository interface {
	Create(user models.User) (models.User, error)
	Update(newData models.User) error
	FindRoot() (models.User, error)
}