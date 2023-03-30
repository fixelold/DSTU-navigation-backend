package getPathPoints

import (
	"errors"
	"navigation/internal/appError"
	"navigation/internal/models"
)

var (
	pointsError = appError.NewAppError("can't set points")
)

// занесение точек начального пути
func (d *data) setAudStartPoints() appError.AppError {
	var err appError.AppError

	axis := d.defenitionAxis(d.audBorderPoints.Widht, d.audBorderPoints.Height)

	axis = d.changeAxis(axis)

	err = d.audStartPoints(axis)
	if err.Err != nil {
		err.Wrap("setAudStartPoints")
		return err
	}

	return appError.AppError{}
}

// функция расчета начального пути от границы аудитории.
func (d *data) audStartPoints(axis int) appError.AppError {
	var err appError.AppError
	err.Wrap("audStartPoints")
	var path models.Coordinates

	// подготовка точек исходя из оси, типа и границ аудитории.
	coordinates := d.preparePoints(audStartPoints, axis, d.audBorderPoints, models.Coordinates{})

	// получение точек для начального пути.
	path, err = d.setPointsAudStart(coordinates, axis, plus)
	if err.Err != nil {
		return err
	}

	// проверка, чтобы точки пути не находились в пределах аудиториию
	check, err := d.repository.checkBorderAud(path)
	if err.Err != nil {
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
		return appError.AppError{}
	} else {
		path, err = d.setPointsAudStart(coordinates, axis, plus)
		if err.Err != nil {
			return err
		}

		check, err = d.repository.checkBorderAud(path)
		if err.Err != nil {
			return err
		}

		if check {
			d.points = append(d.points, path)
			return appError.AppError{}
		} else {
			pointsError.Wrap("audStartPoints")
			pointsError.Err = errors.New("the dots are in the audience area")
			return *pointsError
		}
	}
}
