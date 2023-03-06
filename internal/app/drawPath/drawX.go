package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

// func (d *Path) drawAudX() error {
// 	var err error
// 	var path models.Coordinates

// 	path, err = d.drawAxisX(d.AudienceBorderPoint, plus)
// 	if err != nil {
// 		logging.GetLogger().Errorln("draw drawAxis. Error - ", err)
// 		return err
// 	}

// 	check, err := d.Repository.checkBorderAud(path)
// 	if err != nil {
// 		logging.GetLogger().Errorln("checkBorderSectro db error - ", err)
// 		return err
// 	}

// 	if check {
// 		d.Path = append(d.Path, path)
// 		return nil
// 	} else {
// 		path, err = d.drawAxisX(d.AudienceBorderPoint, minus)
// 		if err != nil {
// 			return err
// 		}

// 		check, err = d.Repository.checkBorderAud(path)
// 		if err != nil {
// 			return err
// 		}

// 		if check {
// 			d.Path = append(d.Path, path)
// 			return nil
// 		} else {
// 			err = User000004
// 			return User000004
// 		}
// 	}
// }

// func (d *Path) drawAxisX(borderPoints models.Coordinates, sign int) (models.Coordinates, error) {
// 	var path models.Coordinates
// 	var err error

// 	switch sign {
// 	case plus:
// 		path.X = borderPoints.X + 1
// 		path.Y = (borderPoints.Y + (borderPoints.Height + borderPoints.Y)) / 2
// 		path.Widht = WidhtX
// 		path.Height = HeightX

// 	case minus:
// 		path.X = borderPoints.X + 1
// 		path.Y = (borderPoints.Y + (borderPoints.Height + borderPoints.Y)) / 2
// 		path.Widht = -WidhtX
// 		path.Height = -HeightX

// 	default:
// 		d.logger.Errorln("draw axis X default")
// 		err = User000004
// 	}

// 	return path, err
// }

func (d *Path) getPoints(axis int) error {
	var err error
	var path models.Coordinates
	var x int
	var y int

	XX := d.AudienceBorderPoint.X + 1
	YX := (d.AudienceBorderPoint.Y + (d.AudienceBorderPoint.Height + d.AudienceBorderPoint.Y)) / 2

	XY := (d.AudienceBorderPoint.X + (d.AudienceBorderPoint.Widht + d.AudienceBorderPoint.X)) / 2
	YY := d.AudienceBorderPoint.Y + 1

	if axis == AxisX {
		x = XX
		y = YX
	} else if axis == AxisY {
		x = XY
		y = YY
	}
	path, err = d.get(x, y, plus)
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
		path, err = d.get(x, y, minus)
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

func (d *Path) get(x, y, sign int) (models.Coordinates, error) {
	var path models.Coordinates
	var err error

	switch sign {
	case plus:
		path.X = x
		path.Y = y
		path.Widht = WidhtX
		path.Height = HeightX

	case minus:
		path.X = x
		path.Y = y
		path.Widht = -WidhtX
		path.Height = -HeightX

	default:
		d.logger.Errorln("draw axis X default")
		err = User000004
	}

	return path, err
}