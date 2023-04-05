package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	getAudPoints(audNumber string) (models.Coordinates, appError.AppError)
	getAudBorderPoint(number string) (models.Coordinates, appError.AppError)
	getSectorBorderPoint(entry, exit int) (models.Coordinates, appError.AppError)
	checkBorderAud(coordinates models.Coordinates) (bool, appError.AppError)
	checkBorderSector(coordinates models.Coordinates) (bool, appError.AppError)
	getTransitionSectorBorderPoint(start, exit int) (models.Coordinates, appError.AppError)
}
