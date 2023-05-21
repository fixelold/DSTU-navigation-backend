package audToAud

import (
	"errors"

	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)

var (
	pointsError = appError.NewAppError("can't set points")
)

func (a *audToAudController) getStartPoints() appError.AppError {

	//отрисовка для стартовой аудитории
	axis := axes.DefenitionAxis(a.startAudBorderPoint.Widht, a.startAudBorderPoint.Height, a.constData.axisX, a.constData.axisY)
	axis = axes.ChangeAxis(axis, a.constData.axisX, a.constData.axisY)

	err := a.start(axis, a.startAudBorderPoint, a.startAud)
	if err.Err != nil {
		err.Wrap("startPath")
		return err
	}

	//отрисовка для конечной аудитории
	axis = axes.DefenitionAxis(a.endAudBorderPoint.Widht, a.endAudBorderPoint.Height, a.constData.axisX, a.constData.axisY)
	axis = axes.ChangeAxis(axis, a.constData.axisX, a.constData.axisY)

	err = a.start(axis, a.endAudBorderPoint, a.endAud)
	if err.Err != nil {
		err.Wrap("startPath")
		return err
	}

	return appError.AppError{}
}

func (a *audToAudController) start(axis int, borderPoint models.Coordinates, audNumber string) appError.AppError {
	var err appError.AppError
	var path models.Coordinates
	repository := NewRepository(a.client)
 
	// подготовка точек исходя из оси, типа и границ аудитории.
	coordinates := a.preparation(axis, borderPoint)
	// получение точек для начального пути.
	path, err = a.pathBuilding(coordinates, axis, a.constData.positiveCoordinate)
	if err.Err != nil {
		err.Wrap("audStartPoints")
		return err
	}

	// проверка, чтобы точки пути не находились в пределах аудиториию
	check, err := repository.checkBorderAud(path, audNumber)
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
		a.points = append(a.points, path)
		return appError.AppError{}
	} else {
		path, err = a.pathBuilding(coordinates, axis, a.constData.negativeCoordinate)
		if err.Err != nil {
			err.Wrap("audStartPoints")
			return err
		}

		check, err = repository.checkBorderAud(path, audNumber)
		if err.Err != nil {
			err.Wrap("audStartPoints")
			return err
		}
		
		if check {
			a.points = append(a.points, path)
			return appError.AppError{}
		} else {
			pointsError.Wrap("audStartPoints")
			pointsError.Err = errors.New("the dots are in the audience area")
			return *pointsError
		}
	}
	// return appError.AppError{}
}