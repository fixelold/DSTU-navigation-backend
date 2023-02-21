package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

func (d *drawPathAud2Sector) drawX() error {
	var err error
	var path models.Coordinates

	path, err = drawAxisX(d.AudienceBorderPoint, plus)
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
		d.Path = append(d.Path, path.X, path.Y, path.Widht, path.Height)
		return nil
	} else {
		path, err = drawAxisX(d.AudienceBorderPoint, minus)
		if err != nil {
			logging.GetLogger().Errorln("draw else. Error - ", err)
			return err
		}

		check, err = d.Repository.checkBorderAud(path)
		if err != nil {
			logging.GetLogger().Errorln("checkBorderSectro db error - ", err)
			return err
		}

		if check {
			d.Path = append(d.Path, path.X, path.Y, path.Widht, path.Height)
			return nil
		} else {
			err = User000004
			logging.GetLogger().Errorln("draw else 2. Error - ", err)
			return User000004
		}
	}
}

func drawAxisX(borderPoints models.Coordinates, sign int) (models.Coordinates, error) {
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
		err = User000004
	}

	return path, err
}
