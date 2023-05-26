package audToAud

import (
	"fmt"

	"navigation/internal/models"
)

func (m *middleController) pointsRightX(borderPoint, points, endPoints models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	
	// Если ширина 5, то значит, что конечная аудитория находится на одной оси с начальной
	if points.Widht == 5 {result = m.rightFinal(borderPoint, points, factor)
	} else if endPoints.Height == 5 {
		if endPoints.Widht == 10 {result = m.endRightLeftX(borderPoint, points, factor)	
		} else if endPoints.Widht == -10 {result = m.endRightRightX(borderPoint, points, factor)}
	}

	return result
}

func (m *middleController) rightFinal(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	result = models.Coordinates{
		X: points.X,
		Y: points.Y + points.Height,
		Widht: (borderPoint.X - points.X), // delete +10 for 1-344 to 1-340
		Height: 5 * factor,
	}
	return result
}

func (m *middleController) endRightLeftX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X,
		Y: points.Y + points.Height,
		Widht: (borderPoint.X - points.X) + 15,
		Height: 5 * factor,
	}

	return result
}

func (m *middleController) endRightRightX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X,
		Y: points.Y + points.Height,
		Widht: (borderPoint.X - points.X) - 10,
	   Height: 5 * factor,
	}

	fmt.Println("result: ", result)
	return result
}