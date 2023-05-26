package audToAud

import (
	"navigation/internal/models"
)

func (m *middleController) pointsLeftX(borderPoint, points, endPoints models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	if endPoints.Widht == 5 {result = m.final(borderPoint, points, factor)
	} else if endPoints.Height == 5 {
		if endPoints.Widht == 10 {result = m.endLeftX(borderPoint, points, factor)	
		} else if endPoints.Widht == -10 {result = m.endRightX(borderPoint, points, factor)}
	}
	
	return result
}

func (m *middleController) final(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height,
		Widht: (borderPoint.X - (points.X + points.Widht)) + 10,
		Height: 5 * factor,
	}

	return result
}

func (m *middleController) endLeftX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height,
		Widht: (borderPoint.X - (points.X + points.Widht)) + 10,
		Height: -5,
	}

	return result
}

func (m *middleController) endRightX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height,
		Widht: (borderPoint.X - (points.X + points.Widht)) - 15,
		Height: -5,
	}

	return result
}