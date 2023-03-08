package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

const ()

func (d *Path) getPoints(axis int) error {
	var err error
	var path models.Coordinates
	coordinates := prepare(d.AudienceBorderPoint, axis)

	path, err = d.get(coordinates, plus, axis)
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
		path, err = d.get(coordinates, minus, axis)
		if err != nil {
			return err
		}

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

func (d *Path) get(coordinates models.Coordinates, sign, axis int) (models.Coordinates, error) {
	var path models.Coordinates
	var err error

	switch axis {
	case AxisX:
		switch sign {
		case plus:
			path.X = coordinates.X
			path.Y = coordinates.Y
			path.Widht = coordinates.Widht
			path.Height = coordinates.Height
	
		case minus:
			path.X = coordinates.X
			path.Y = coordinates.Y
			path.Widht = -coordinates.Widht
			path.Height = coordinates.Height
	
		default:
			d.logger.Errorln("Function get. Error switch default")
			err = User000004
		}
	case AxisY:
		switch sign {
		case plus:
			path.X = coordinates.X
			path.Y = coordinates.Y
			path.Widht = coordinates.Widht
			path.Height = coordinates.Height
	
		case minus:
			path.X = coordinates.X
			path.Y = coordinates.Y
			path.Widht = coordinates.Widht
			path.Height = -coordinates.Height
	
		default:
			d.logger.Errorln("Function get. Error switch default")
			err = User000004
		}
	}

	return path, err
}

func prepare(borderPoint models.Coordinates, axis int) models.Coordinates {
	var coordinates models.Coordinates

	XX := borderPoint.X + 1
	YX := (borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2

	XY := (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2
	YY := borderPoint.Y + 1

	if axis == AxisX {
		coordinates.X = XX
		coordinates.Y = YX
		coordinates.Widht = WidhtX
		coordinates.Height = HeightX

	} else if axis == AxisY {
		coordinates.X = XY
		coordinates.Y = YY
		coordinates.Widht = WidhtY
		coordinates.Height = HeightY
	}

	return coordinates
}
