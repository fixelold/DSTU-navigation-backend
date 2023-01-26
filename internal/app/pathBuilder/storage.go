package pathBuilder

import "navigation/internal/models"

type Repository interface {
	GetSectorLink() ([]models.SectorLink, error)
}
