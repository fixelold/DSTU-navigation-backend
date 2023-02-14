package drawPath

import "navigation/internal/models"

type Repository interface {
	getAuditoryPosition(audNumber string) (*models.AuditoryPosition, error)
	getBorderPoint(number string) (*models.BorderPoint, error)
}