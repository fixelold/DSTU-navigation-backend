package drawPath

import (
	"database/sql/driver"
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
	var err error
	var points []int
	axis := defenitionAxis(borderPoints)

	switch axis {

	case AxisX:
		path, err := draw(WidhtX, HeightX, borderPoints, auditory)
		if err != nil {
			return nil, err
		}

		points = append(points, path.X, path.Y, path.Widht, path.Height)
	case AxisY:
		path, err := draw(WidhtY, HeightY, borderPoints, auditory)
		if err != nil {
			return nil, err
		}

		points = append(points, path.X, path.Y, path.Widht, path.Height)
	default:
		err = User000004
	}

	return points, err
}

func draw(widht, height int, borderPoints, auditory *models.Reactangle) (models.Reactangle, error) {
	var err error
	path, err := drawAxisX(widht, height, borderPoints, plus)
	if err != nil {
		return path, err
	}

	if checkBorder(&path, auditory) {
		return path, nil
	} else {
		path, err = drawAxisX(widht, height, borderPoints, minus)
		if err != nil {
			return path, err
		}

		if checkBorder(&path, auditory) {
			return path, nil
		} else {
			return path, User000004
		}
	}
}

func drawAxisX(widht, height int, borderPoints *models.Reactangle, sign int) (models.Reactangle, error) {
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

	return path, err
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
