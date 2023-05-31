package audToAud

import (
	"navigation/internal/models"
)

func (m *middleController) firstFinalX(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates
	var factoHeight int

	if m.endPoints.Widht == 10 {
		factoHeight = 1
	} else {
		factoHeight = 0
	}

	result = models.Coordinates{
		X: points.X,
		Y: points.Y + (points.Height * factoHeight),
		Widht: borderPoint.X - points.X,
		Height: 5,
	}
	return result
}

func (m *middleController) firstFinalY(borderPoint, points models.Coordinates, factor int) models.Coordinates {
	var result models.Coordinates

	result = models.Coordinates{
		X: points.X,
		Y: points.Y,
		Widht: 5,
		Height: (borderPoint.Y - points.Y),
	}

	return result
}