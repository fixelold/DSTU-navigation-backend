package getPathPoints

import "navigation/internal/models"

type Repository interface {
	getAudPoints(audNumber string) (models.Coordinates, error)
	getAudBorderPoint(number string) (models.Coordinates, error)
	getSectorBorderPoint(entry, exit int) (models.Coordinates, error)
	checkBorderAud(coordinates models.Coordinates) (bool, error)
	// checkBorderSector(coordinates models.Coordinates) (bool, error)
}
