package middle

import (
	"navigation/internal/models"
)

func (m *middleController) leftX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: 5,
            Height: borderPoint.Y - (points.Y + points.Height),
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y,
			Widht: 5,
			Height: borderPoint.Y - points.Y,
		}
	}

	return result
}

func (m *middleController) rightX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: -5,
            Height: borderPoint.Y - (points.Y + points.Height),
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y,
			Widht: -5,
			Height: borderPoint.Y - points.Y,
		}
	}

	return result
}

func (m *middleController) upAndDownX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	var lenght = len(m.Points)

	m.Points[lenght-1].X = borderPoint.X + 10

	result = models.Coordinates{
		X: points.X,
		Y: points.Y + points.Height,
		Widht: 5,
		Height: borderPoint.Y - (points.Y + points.Height),
	}

	return result
}