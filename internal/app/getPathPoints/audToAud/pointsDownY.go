package audToAud

import (
	"navigation/internal/models"
)

func (m *middleController) pointsDownY(borderPoint, points, endPoints models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	if endPoints.Height == 5 {result = m.pointsDownFinal(borderPoint, points, factor)
	} else if endPoints.Widht == 5 {
		if endPoints.Height == 10 {result = m.endDownUpY(borderPoint, points, factor)	
		} else if endPoints.Height == -10 {result = m.endDownDownY(borderPoint, points, factor)}
	}
	
	return result
}

func (m *middleController) pointsDownFinal(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	factor = 1
	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y + points.Height, // added points.Height for end 1-33а
		Widht: 5 * factor,
		Height: (borderPoint.Y - (points.Y + points.Height)) + 10,  // added points.Height for end 1-33а
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