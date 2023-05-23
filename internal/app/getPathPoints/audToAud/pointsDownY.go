package audToAud

import "navigation/internal/models"

func (m *middleController) pointsDownY(borderPoint, points, endPoints models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	if endPoints.Height == 5 {result = m.rightFinal(borderPoint, points, factor)
	} else if endPoints.Height == 5 {
		if endPoints.Height == 10 {result = m.endRightLeftX(borderPoint, points, factor)	
		} else if endPoints.Height == -10 {result = m.endRightRightX(borderPoint, points, factor)}
	}
	
	return result
}

// func (m *middleController) downFinal(borderPoint, points models.Coordinates, factor int) models.Coordinates {
// 	var result models.Coordinates

// 	result = models.Coordinates{
// 		X: points.X + points.Widht,
// 		Y: points.Y,
// 		Widht: 5 * factor,
// 		Height: (borderPoint.Y - points.Y) + 10,
// 	}

// 	return result
// }

// func (m *middleController) endDownUpX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
// 	var result models.Coordinates

// 	result = models.Coordinates{
// 		X: points.X + points.Widht,
// 		Y: points.Y,
// 		Widht: 5 * factor,
// 		Height: (borderPoint.Y - points.Y) + 10,
// 	}

// 	return result
// }

func (m *middleController) endDownDownX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X + points.Widht,
		Y: points.Y,
		Widht: 5 * factor,
		Height: (borderPoint.Y - points.Y) - 15,
	}

	return result
}