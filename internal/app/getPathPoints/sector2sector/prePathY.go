package sectorToSector

import "navigation/internal/models"

func (s *sectorToSectorController) prePathUpY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if points.Widht == -5 {
		result = models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: -5,
			Height: ((borderPoint.Y + borderPoint.Height) - (points.Y + points.Height)) - 10,
		}
	} else if points.Widht == 5 {
		result = models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: 5,
			Height: ((borderPoint.Y + borderPoint.Height) - (points.Y + points.Height)) - 10,
		}
	}

	return result
}

func (s *sectorToSectorController) prePathDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if points.Widht == -5 {
		result = models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: -5,
			Height: (borderPoint.Y - (points.Y + points.Height)) + 10,
		}
	} else if points.Widht == 5 {
		result = models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: 5,
			Height: (borderPoint.Y - (points.Y + points.Height)) + 10,
		}
	}

	return result
}