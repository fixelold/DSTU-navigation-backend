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

	path, err = d.get(coordinates, plus, axis)
	if err != nil {
		return err
	}

	check, err := d.Repository.checkBorderAud(path)
	if err != nil {
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
