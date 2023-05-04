package start

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	checkBorderAud(coordinates models.Coordinates, audNumber string) (bool, appError.AppError)
}
