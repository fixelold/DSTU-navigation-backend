package audToAud

import (
	"fmt"

	"navigation/internal/models"
)

func (m *middleController) pointsLeftX(borderPoint, points, endPoints models.Coordinates, factor int, final bool) models.Coordinates {
	var result models.Coordinates
	// if points.Widht == 5 {result = m.final(borderPoint, points, factor, axis)
	// } else if points.Height == 5 {
	// 	if endPoints.Widht == 10 {result = m.endLeftX(borderPoint, points, factor)	
	// 	} else if endPoints.Widht == -10 {result = m.endRightX(borderPoint, points, factor)}
	// }

	if final {result = m.final(borderPoint, points, factor)
		} else {
			if endPoints.Height == 10 || endPoints.Height == -10 { result = m.final(borderPoint, points, factor)}
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
		Widht: (borderPoint.X - (points.X + points.Widht)),
		Height: 5 * factor,
	}
	
	fmt.Println("result: ", result)
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