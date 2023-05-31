package sectorToSector

import "navigation/internal/models"

func (s *sectorToSectorController) finalY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if points.Widht == -5 {
		result = models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: -5,
			Height: borderPoint.Y - (points.Y + points.Height),
		}
	} else if points.Widht == 5 {
		result = models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: 5,
			Height: borderPoint.Y - (points.Y + points.Height),
		}
	}

	return result
}