package drawPath

import "navigation/internal/models"

type Repository interface {
	getAuditoryPosition(audNumber string) (*models.Coordinates, error)
	getAudBorderPoint(number string) (*models.Coordinates, error)
	getSectorBorderPoint(number int) (*models.Coordinates, error)
	checkBorderAud(coordinates models.Coordinates) (bool, error)
	checkBorderSector(coordinates models.Coordinates) (bool, error)
}