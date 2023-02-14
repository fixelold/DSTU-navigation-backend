package drawPath

import "navigation/internal/models"

type Repository interface {
	getAuditoryPosition(audNumber string) (*models.Reactangle, error)
	getBorderPoint(number string) (*models.Reactangle, error)
}