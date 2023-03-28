package getPathPoints

import "navigation/internal/models"

func (d *data) getAudStartPoints() error {
	var err error

	axis := d.defenitionAxis(d.audBorderPoints.Widht, d.audBorderPoints.Height)

	err = d.audStartPoints(axis)
	if err != nil {
		return err
	}

	return nil
}

func (d *data) audStartPoints(axis int) error {
	var err error
	var path models.Coordinates
	coordinates := d.preparePoints(audStartPoints, axis, d.audBorderPoints)

	path, err = d.setPoints(audStartPoints, plus, axis, coordinates)
	if err != nil {
		return err
	}

	check, err := d.repository.checkBorderAud(path)
	if err != nil {
		return err
	}

	if check {
		d.points = append(d.points, path)
		return nil
	} else {
		path, err = d.setPoints(audStartPoints, minus, axis, coordinates)
		if err != nil {
			return err
		}

		check, err = d.repository.checkBorderAud(path)
		if err != nil {
			return err
		}

		if check {
			d.points = append(d.points, path)
			return nil
		} else {
			return nil
		}
	}
}
