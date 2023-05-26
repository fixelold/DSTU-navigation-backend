package audToAud

import (
	"navigation/internal/models"
)

func (m *middleController) pointsUpY(borderPoint, points, endPoints models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	if endPoints.Widht == 5 {result = m.upFinal(borderPoint, points, factor)
	} else if endPoints.Height == 5 {
		if endPoints.Widht == 10 {result = m.endUpUpX(borderPoint, points, factor)	
		} else if endPoints.Widht == -10 {result = m.endUpDownX(borderPoint, points, factor)}
	}
	
	return result
}

func (m *middleController) upFinal(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height,
		Widht: 5 * factor,
		Height: (borderPoint.Y - (points.Y + points.Height)) + 10,
	}

	return result
}

func (m *middleController) endUpUpX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	factor = -1
	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height,
		Widht: 5 * factor,
		Height: (borderPoint.Y - (points.Y + points.Height)) + 15,
	}

	return result
}

func (m *middleController) endUpDownX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height,
		Widht: 5 * factor,
		Height: (borderPoint.Y - (points.Y + points.Height)) - 10,
	}

	return result
}