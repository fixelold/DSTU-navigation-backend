package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	getAudPoints(audNumber string) (models.Coordinates, appError.AppError)
	getTransitionPoints(number int) (models.Coordinates, appError.AppError)
}
