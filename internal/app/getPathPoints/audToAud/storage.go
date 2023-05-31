package audToAud

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	checkBorderSector(coordinates models.Coordinates) (bool, appError.AppError)
	checkBorderAud(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError)
	checkBorderAudY(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError)
	checkBorderAud2(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError)
	checkBorderAud3(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError)
	checkBorderAud4(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError)
}
