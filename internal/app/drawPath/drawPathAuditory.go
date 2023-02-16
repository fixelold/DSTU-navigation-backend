package drawPath

import (
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
		path, err := drawX(borderPoints, auditory)
		if err != nil {
			logging.GetLogger().Errorln("DrawPathAuditory case AxisX. Error - ", err)
			return nil, err
		}

		points = append(points, path.X, path.Y, path.Widht, path.Height)

	case AxisY:
		logging.GetLogger().Info("AxisY - work!")
		path, err := drawY(borderPoints, auditory)
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

func defenitionAxis(position *models.Reactangle) int {
	if position.Widht == 1 {
		return AxisX
	} else if position.Height == 1 {
		return AxisY
	} else {
		return 0
	}
}
