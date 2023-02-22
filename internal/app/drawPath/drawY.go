package drawPath

import (
	"fmt"
	"navigation/internal/logging"
	"navigation/internal/models"
)

func (d *drawPathAud2Sector) drawAudY() error {
	fmt.Println("Work - drawAudY")
	var err error
	var path models.Coordinates

	path, err = drawAxisY(d.AudienceBorderPoint, plus)
	if err != nil {
		logging.GetLogger().Errorln("draw drawAxis. Error - ", err)
		return err
	}

	check, err := d.Repository.checkBorderSector(path)
	if err != nil {
		logging.GetLogger().Errorln("checkBorderSectro db error - ", err)
		return err
	}

	if check {
		d.Path = append(d.Path, path)
		return nil
	} else {
		path, err = drawAxisY(d.AudienceCoordinates, minus)

		if err != nil {
			logging.GetLogger().Errorln("draw else. Error - ", err)
			return err
		}

		check, err = d.Repository.checkBorderSector(path)
		if err != nil {
			logging.GetLogger().Errorln("checkBorderSectro db error - ", err)
			return err
		}

		if check {
			d.Path = append(d.Path, path)
			return nil
		} else {
			err = User000004
			logging.GetLogger().Errorln("draw else 2. Error - ", err)
			return User000004
		}
	}
}

func drawAxisY(borderPoints models.Coordinates, sign int) (models.Coordinates, error) {
	var path models.Coordinates
	var err error

	switch sign {
	case plus:
		fmt.Println("Work plus")
		path.X = (borderPoints.X + (borderPoints.Widht + borderPoints.X)) / 2
		path.Y = borderPoints.Y + 1
		path.Widht = WidhtY
		path.Height = HeightY

	case minus:
		fmt.Println("Work minus")
		path.X = (borderPoints.X + (borderPoints.Widht + borderPoints.X)) / 2
		path.Y = borderPoints.Y + 1
		path.Widht = -WidhtY
		path.Height = -HeightY

	default:
		err = User000004
	}

	return path, err
}
