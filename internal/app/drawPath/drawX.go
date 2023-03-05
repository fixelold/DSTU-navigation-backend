package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

func (d *Path) drawAudX() error {
	var err error
	var path models.Coordinates

	path, err = d.drawAxisX(d.AudienceBorderPoint, plus)
	if err != nil {
		logging.GetLogger().Errorln("draw drawAxis. Error - ", err)
		return err
	}

	check, err := d.Repository.checkBorderAud(path)
	if err != nil {
		logging.GetLogger().Errorln("checkBorderSectro db error - ", err)
		return err
	}

	if check {
		d.Path = append(d.Path, path)
		return nil
	} else {
		d.logger.Infoln("draw aud X => draw axis x - minus")
		path, err = d.drawAxisX(d.AudienceBorderPoint, minus)
		if err != nil {
			return err
		}

		d.logger.Infoln("draw aud X => check border aud")
		check, err = d.Repository.checkBorderAud(path)
		if err != nil {
			return err
		}

		if check {
			d.Path = append(d.Path, path)
			return nil
		} else {
			err = User000004
			return User000004
		}
	}
}

func (d *Path) drawAxisX(borderPoints models.Coordinates, sign int) (models.Coordinates, error) {
	var path models.Coordinates
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
		d.logger.Errorln("draw axis X default")
		err = User000004
	}

	return path, err
}
