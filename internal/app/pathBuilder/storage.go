package pathBuilder

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

type Repository interface {
	GetSectorLink() ([]models.SectorLink, appError.AppError)
	GetSector(number string, building uint) (int, appError.AppError)
	GetTransitionSector(sectorNumber, t int) (int, appError.AppError)
}
