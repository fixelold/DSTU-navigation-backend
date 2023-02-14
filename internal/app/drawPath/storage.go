package drawPath

import "navigation/internal/models"

type Repository interface {
	getAuditoryPosition(audNumber string) (*models.AuditoryPosition, error)
}