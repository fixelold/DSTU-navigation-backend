package getPathPoints

import "navigation/internal/models"

// занесение точек начального пути
func (d *data) setAudStartPoints() error {
	var err error

	axis := d.defenitionAxis(d.audBorderPoints.Widht, d.audBorderPoints.Height)

	err = d.audStartPoints(axis)
	if err != nil {
		return err
	}

	return nil
}

// функция расчета начального пути от границы аудитории.
func (d *data) audStartPoints(axis int) error {
	var err error
	var path models.Coordinates

	// подготовка точек исходя из оси, типа и границ аудитории.
	coordinates := d.preparePoints(audStartPoints, axis, d.audBorderPoints)

	// получение точек для начального пути.
	path, err = d.setPoints(audStartPoints, plus, axis, coordinates)
	if err != nil {
		return err
	}

	// проверка, чтобы точки пути не находились в пределах аудиториию
	check, err := d.repository.checkBorderAud(path)
	if err != nil {
		return err
	}

	/*
		если пересечения нет, то точки пути заносятся в главный массив всех точек.
		если пересечение есть, то меняется знак на противополоный.
		например:
			на оси 'x' знак '+' будет означать, что путь будет отрисоваться слева на право
			на оси 'x' знак '-' будет означать, что путь будет отрисовываться справа на лево.
	*/
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
