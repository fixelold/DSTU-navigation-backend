package importantPlaces

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	Create(places models.ImportantPlaces) (models.ImportantPlaces, appError.AppError)
	Read(id int) (models.ImportantPlaces, error)
	Update(oldpPlaces models.ImportantPlaces, newPlaces models.ImportantPlaces) (models.ImportantPlaces, error)
	Delete(id int) (appError.AppError)
	List(numberBuild models.ImportantPlaces) ([]models.ImportantPlaces, appError.AppError)
}