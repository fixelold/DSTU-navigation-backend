package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

func drawY(borderPoints, auditory *models.Reactangle) (models.Reactangle, error) {
	var err error
	var path models.Reactangle

	path, err = drawAxisY(borderPoints, plus)
	if err != nil {
		logging.GetLogger().Errorln("draw drawAxis. Error - ", err)
		return path, err
	}

	if checkBorder(AxisY, &path, auditory) {
		return path, nil
	} else {
		path, err = drawAxisY(borderPoints, minus)

		if err != nil {
			logging.GetLogger().Errorln("draw else. Error - ", err)
			return path, err
		}

		if checkBorder(AxisY, &path, auditory) {
			return path, nil
		} else {
			err = User000004
			logging.GetLogger().Errorln("draw else 2. Error - ", err)
			return path, User000004
		}
	}
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
