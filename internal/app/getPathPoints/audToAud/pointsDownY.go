package audToAud

import (
	"navigation/internal/models"
)

func (m *middleController) pointsDownY(borderPoint, points, endPoints models.Coordinates, factor, axis int) models.Coordinates {
	var result models.Coordinates

	if points.Height == 5 || points.Height == -5  {result = m.pointsDownFinal(borderPoint, points, factor, axis)
	} else if points.Widht == -5 { // added -5 for 1-309 to 1-340
		if endPoints.Height == 10 {result = m.endDownUpY(borderPoint, points, factor)	
		} else if endPoints.Height == -10 {result = m.endDownDownY(borderPoint, points, factor)}
	}
	
	return result
}

func (m *middleController) pointsDownFinal(borderPoint, points models.Coordinates, factor, axis int) models.Coordinates {
	var result models.Coordinates
	var heightFactor int
	if borderPoint.Widht == 1 && axis == m.constData.axisX {
		heightFactor = 1
	} else {
		heightFactor = 0
	}

	factor = 1
	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y,
		Widht: 5 * factor,
		Height: (borderPoint.Y - points.Y) + (10 * heightFactor),
	}

	return result
}

func (m *middleController) endDownUpY(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y,
		Widht: 5 * factor,
		Height: (borderPoint.Y - points.Y) + 10,
	}

	return result
}

func (m *middleController) endDownDownY(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y,
		Widht: 5 * factor,
		Height: (borderPoint.Y - points.Y) - 15,
	}

	return result
}