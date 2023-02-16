package drawPath

import (
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

const (
	AxisX = 1
	AxisY = 2

	WidhtX  = 130
	HeightX = 30

	WidhtY  = 30
	HeightY = 130

	plus  = 0
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
		logging.GetLogger().Info("AxisX - work!")
		path, err := draw(axis, borderPoints, auditory)
		if err != nil {
			logging.GetLogger().Errorln("DrawPathAuditory case AxisX. Error - ", err)
			return nil, err
		}

		points = append(points, path.X, path.Y, path.Widht, path.Height)

	case AxisY:
		logging.GetLogger().Info("AxisY - work!")
		path, err := draw(axis, borderPoints, auditory)
		if err != nil {
			logging.GetLogger().Errorln("DrawPathAuditory case AxisY. Error - ", err.Error())
			return nil, err
		}

		points = append(points, path.X, path.Y, path.Widht, path.Height)
	default:
		logging.GetLogger().Errorln("DrawPathAuditory case default. Error - ", err)
		err = User000004
	}

	return points, err
}

func draw(axis int, borderPoints, auditory *models.Reactangle) (models.Reactangle, error) {
	var err error
	var path models.Reactangle

	if axis == AxisX {
		path, err = drawAxisX(borderPoints, plus)
	} else {
		fmt.Println("Work")
		path, err = drawAxisY(borderPoints, plus)
	}
	if err != nil {
		logging.GetLogger().Errorln("draw drawAxis. Error - ", err)
		return path, err
	}

	if checkBorder(axis, &path, auditory) {
		return path, nil
	} else {
		if axis == AxisX {
			path, err = drawAxisX(borderPoints, minus)
		} else {
			fmt.Println("Work 2")
			path, err = drawAxisY(borderPoints, minus)
		}
		if err != nil {
			logging.GetLogger().Errorln("draw else. Error - ", err)
			return path, err
		}

		if checkBorder(axis, &path, auditory) {
			return path, nil
		} else {
			err = User000004
			logging.GetLogger().Errorln("draw else 2. Error - ", err)
			return path, User000004
		}
	}
}

func drawAxisX(borderPoints *models.Reactangle, sign int) (models.Reactangle, error) {
	var path models.Reactangle
	var err error

	switch sign {
	case plus:
		path.X = borderPoints.X + 1
		path.Y = (borderPoints.Y + (borderPoints.Height + borderPoints.Y)) / 2
		path.Widht = WidhtX
		path.Height = HeightX

	case minus:
		path.X = borderPoints.X + 1
		path.Y = (borderPoints.Y + (borderPoints.Height + borderPoints.Y)) / 2
		path.Widht = -WidhtX
		path.Height = -HeightX

	default:
		err = User000004
	}

	return path, err
}

func drawAxisY(borderPoints *models.Reactangle, sign int) (models.Reactangle, error) {
	var path models.Reactangle
	var err error

	switch sign {
	case plus:
		path.X = (borderPoints.X + (borderPoints.Widht + borderPoints.X)) / 2
		path.Y = borderPoints.Y + 1
		path.Widht = WidhtY
		path.Height = HeightY

	case minus:
		path.X = (borderPoints.X + (borderPoints.Widht + borderPoints.X)) / 2
		path.Y = borderPoints.Y + 1
		path.Widht = -WidhtY
		path.Height = -HeightY

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
