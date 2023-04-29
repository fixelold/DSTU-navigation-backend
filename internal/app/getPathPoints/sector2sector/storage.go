package sectorToSector

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	checkBorderSector(coordinates models.Coordinates) (bool, appError.AppError)
	checkBorderAud(coordinates models.Coordinates) (bool, appError.AppError)
}
