package drawPath

import "navigation/internal/models"

type Repository interface {
	getAuditoryPosition(audNumber string) (*models.Coordinates, error)
	getBorderPoint(number string) (*models.Coordinates, error)
}