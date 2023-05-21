package middle

import "navigation/internal/models"

func (m *middleController) preparationLeftY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: 5,
            Height: ((borderPoint.Y + borderPoint.Height) - (points.Y + points.Height)) - 9,
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y,
			Widht: 5,
			Height: borderPoint.Y - points.Y + 9,
		}
	}

	return result
}

func (m *middleController) preparationRightY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates

	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: -5,
            Height: ((borderPoint.Y + borderPoint.Height) - (points.Y + points.Height)) - 9,
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y,
			Widht: -5,
			Height: borderPoint.Y - points.Y + 9,
		}
	}

	return result
}