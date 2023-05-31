package middle

import (
	"navigation/internal/models"
)

func (m *middleController) preparationUpX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if borderPoint.X < points.X {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: ((borderPoint.X + borderPoint.Widht) - (points.X + points.Widht)) - 10,
            Height: 5,
		}
	} else if borderPoint.X > points.X {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - points.X + 10,
            Height: 5,
		}
	}

	return result
}

func (m *middleController) preparationDownX(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.X < points.X {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: ((borderPoint.X + borderPoint.Widht) - (points.X + points.Widht)) - 10,
            Height: -5,
		}
	} else if borderPoint.X > points.X {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y + points.Height,
            Widht: borderPoint.X - points.X + 10,
            Height: -5,
		}
	}

	return result
}