package sectorToSector

import (
	"navigation/internal/models"
)

func (m *sectorToSectorController) downLeftX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.X < points.X {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - (points.X + points.Widht),
            Height: -5,
		}
	} else if borderPoint.X > points.X {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - points.X,
            Height: -5,
		}
	}

	return result
}

func (m *sectorToSectorController) downRightX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.X < points.X {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - points.X,
            Height: -5,
		}
	} else if borderPoint.X > points.X {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - (points.X + points.Widht),
            Height: -5,
		}
	}

	return result
}

func (m *sectorToSectorController) upLeftX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.X < points.X {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - (points.X + points.Widht),
            Height: 5,
		}
	} else if borderPoint.X > points.X {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - points.X,
            Height: 5,
		}
	}

	return result
}

func (m *sectorToSectorController) upRightX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.X < points.X {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - points.X,
            Height: 5,
		}
	} else if borderPoint.X > points.X {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - (points.X + points.Widht),
            Height: 5,
		}
	}

	return result
}