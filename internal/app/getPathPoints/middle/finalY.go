package middle

import (
	"navigation/internal/models"
)

// расчетт пути елси конечная аудитория находится слева
func (m *middleController) leftY(borderPoint, points models.Coordinates) models.Coordinates {
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

// расчетт пути елси конечная аудитория находится справа
func (m *middleController) rightY(borderPoint, points models.Coordinates) models.Coordinates {
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

// расчетт пути елси конечная аудитория находится внизу или вверху
func (m *middleController) upAndDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	var lenght = len(m.Points)

	m.Points[lenght-1].X = borderPoint.X + 10
	points.X = borderPoint.X + 10

	result = models.Coordinates{
		X: points.X,
		Y: points.Y + points.Height,
		Widht: 5,
		Height: borderPoint.Y - (points.Y + points.Height),
	}

	return result
}