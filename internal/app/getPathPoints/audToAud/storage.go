package audToAud

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	checkBorderAud(coordinates models.Coordinates, audNumber string) (bool, appError.AppError)
	checkBorderAud2(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError)
}
