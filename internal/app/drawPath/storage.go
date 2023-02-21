package drawPath

import "navigation/internal/models"

type Repository interface {
	getAuditoryPosition(audNumber string) (*models.Coordinates, error)
	getBorderPoint(number string) (*models.Coordinates, error)
	checkBorderAud(coordinates models.Coordinates) (bool, error)
	checkBorderSector(coordinates models.Coordinates) (bool, error)
}