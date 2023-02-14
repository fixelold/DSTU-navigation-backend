package drawPath

import "navigation/internal/models"

const (
	AxisX = 0
	AxisY = 1
)

func DrawPathAuditory(position *models.BorderPoint) ([]int, error) {

	// 2 определить ось рисования

	// 3 начать рисовать

	// 4 возвратить данные
}

func defenitionAxis(position *models.BorderPoint) int {
	if position.Widht == 1 {
		return AxisX
	} else if position.Height == 1 {
		return AxisY
	} else {
		return 0
	}
}