package sectorToSector

import "navigation/internal/models"

func (s *sectorToSectorController) finalX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if points.Height == -5 {
		result = models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y ,
			Widht: borderPoint.X - (points.X + points.Widht),
			Height: -5,
		}
	} else if points.Height == 5 {
		result = models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y ,
			Widht: borderPoint.X - (points.X + points.Widht),
			Height: 5,
		}
	}

	return result
}