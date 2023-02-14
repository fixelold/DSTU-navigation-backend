package drawPath

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

const (
	AxisX = 1
	AxisY = 2

	WidhtX = 130
	HeightX = 30

	WidhtY = 30
	HeightY = 130
	
	plus = 0
	minus = 1
)

var (
	User000004 = appError.NewError("drawPath", "GetSelector", "Input does not match desired length", "-", "US-000004")
)

func DrawPathAuditory(borderPoints, auditory *models.Reactangle) ([]int, error) {
	var path models.Reactangle
	var err error
	var points []int
	axis := defenitionAxis(borderPoints)

	switch axis {

	case AxisX:
		path.X = borderPoints.Widht / 2
		path.Y = borderPoints.Y
		path.Widht = path.X + WidhtX
		path.Height = path.Y + HeightX

		if checkBorder(&path, auditory) {
			points = append(points, path.X, path.Y, path.Widht, path.Height)
		} else {
			path.Widht = path.X - WidhtX
			path.Height = path.Y - HeightX
		}
	case AxisY:
		//....
	default:
		err = User000004
	}
	// 3 начать рисовать

	// 4 возвратить данные

	return points, err
}

func drawAxisX(borderPoints *models.Reactangle, sign int) (*models.Reactangle, error) {
	var path models.Reactangle
	var err error
	switch sign {
	
	case plus:
		path.X = borderPoints.Widht / 2
		path.Y = borderPoints.Y
		path.Widht = path.X + WidhtX
		path.Height = path.Y + HeightX

	case minus:
		path.X = borderPoints.Widht / 2
		path.Y = borderPoints.Y
		path.Widht = path.X - WidhtX
		path.Height = path.Y - HeightX
	
	default:
		err = User000004
	}

	return &path, err
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
