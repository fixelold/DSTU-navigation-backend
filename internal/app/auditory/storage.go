package auditory

import "navigation/internal/models"

type Repository interface {
	Update(description, number string) error
	Read(number string) (models.AuditoryDescription, error)
}