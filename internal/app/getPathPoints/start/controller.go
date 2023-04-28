package start

import (
	"errors"

	"navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)

var (
	pointsError = appError.NewAppError("can't set points")
)

type startController struct {
	points []int
	audienceBoundaryPoints models.Coordinates
}

// занесение точек начального пути
func (s *startController) setAudStartPoints() appError.AppError {
	var err appError.AppError

	a := axis.DefenitionAxis(s.audienceBoundaryPoints.Widht, s.audienceBoundaryPoints.Height)
	a = axis.ChangeAxis(a)

	// err = d.audStartPoints(axis)
	if err.Err != nil {
		err.Wrap("setAudStartPoints")
		return err
	}

	return appError.AppError{}
}

// функция расчета начального пути от границы аудитории.
func (d *data) audStartPoints(axis int) appError.AppError {
	var err appError.AppError
	var path models.Coordinates

	// подготовка точек исходя из оси, типа и границ аудитории.
	coordinates := d.preparePoints(audStartPoints, axis, d.audBorderPoints, models.Coordinates{})

	// получение точек для начального пути.
	path, err = d.setPointsAudStart(coordinates, axis, plus)
	if err.Err != nil {
		err.Wrap("audStartPoints")
		return err
	}

	// проверка, чтобы точки пути не находились в пределах аудиториию
	check, err := d.repository.checkBorderAud(path)
	if err.Err != nil {
		err.Wrap("audStartPoints")
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
		path, err = d.setPointsAudStart(coordinates, axis, minus)
		if err.Err != nil {
			err.Wrap("audStartPoints")
			return err
		}

		check, err = d.repository.checkBorderAud(path)
		if err.Err != nil {
			err.Wrap("audStartPoints")
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