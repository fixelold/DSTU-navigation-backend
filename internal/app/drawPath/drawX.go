package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

func drawX(borderPoints, auditory *models.Reactangle) (models.Reactangle, error) {
	var err error
	var path models.Reactangle

	path, err = drawAxisX(borderPoints, plus)
	if err != nil {
		logging.GetLogger().Errorln("draw drawAxis. Error - ", err)
		return path, err
	}

	if checkBorder(AxisX, &path, auditory) {
		return path, nil
	} else {
		path, err = drawAxisX(borderPoints, minus)

		if err != nil {
			logging.GetLogger().Errorln("draw else. Error - ", err)
			return path, err
		}

		if checkBorder(AxisX, &path, auditory) {
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
