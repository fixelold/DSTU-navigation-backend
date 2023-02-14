package drawPath

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

const (
	AxisX = 1
	AxisY = 2
)

var (
	User000004 = appError.NewError("drawPath", "GetSelector", "Input does not match desired length", "-", "US-000004")
)

func DrawPathAuditory(position, auditory *models.Reactangle) ([]int, error) {
	var err error
	var points []int
	axis := defenitionAxis(position)

	switch axis {

	case AxisX:
		//....
	case AxisY:
		//....
	default:
		err = User000004
	}
	// 3 начать рисовать

	// 4 возвратить данные

	return points, err
}

func defenitionAxis(position *models.Reactangle) int {
	if position.Widht == 1 {
		return AxisX
	} else if position.Height == 1 {
		return AxisY
	} else {
		return 0
	}
}
