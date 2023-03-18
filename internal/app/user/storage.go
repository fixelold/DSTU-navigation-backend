package user

import "navigation/internal/models"

type Repository interface {
	Create(user models.User) (models.User, error)
}