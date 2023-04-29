package start

import (
	"errors"

	"navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/models"
)

var (
	pointsError = appError.NewAppError("can't set points")
)

type constData struct {
	positiveCoordinate int
	negativeCoordinate int
	axisX int
	axisY int
	widhtX int
	heightX int
	widhtY int
	heightY int
}

type startController struct {
	points []models.Coordinates
	audienceBoundaryPoints models.Coordinates
	constData constData
	client postgresql.Client
}

func NewStartController(
	audienceBoundaryPoints models.Coordinates,
	client postgresql.Client,
	positiveCoordinate, negativeCoordinate int,
	axisX, axisY, widhtX, heightX, widhtY, heightY int) *startController {
		return &startController{
			audienceBoundaryPoints: audienceBoundaryPoints,
			client: client,
			constData: constData{
				positiveCoordinate: positiveCoordinate,
				negativeCoordinate: negativeCoordinate,
				axisX: axisX,
				axisY: axisY,
				widhtX: widhtX,
				heightX: heightX,
				widhtY: widhtY,
				heightY: heightY,
			},
		}
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
func (s *startController) audStartPoints(axis int) appError.AppError {
	var err appError.AppError
	var path models.Coordinates

	// подготовка точек исходя из оси, типа и границ аудитории.
	coordinates := s.preparation(axis, s.audienceBoundaryPoints)

	// получение точек для начального пути.
	path, err = s.pathBuilding(coordinates, axis, s.constData.positiveCoordinate)
	if err.Err != nil {
		err.Wrap("audStartPoints")
		return err
	}

	// проверка, чтобы точки пути не находились в пределах аудиториию
	repository := NewRepository(s.client)
	check, err := repository.checkBorderAud(path)
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
		s.points = append(s.points, path)
		return appError.AppError{}
	} else {
		path, err = s.pathBuilding(coordinates, axis, s.constData.negativeCoordinate)
		if err.Err != nil {
			err.Wrap("audStartPoints")
			return err
		}

		check, err = repository.checkBorderAud(path)
		if err.Err != nil {
			err.Wrap("audStartPoints")
			return err
		}
		
		if check {
			s.points = append(s.points, path)
			return appError.AppError{}
		} else {
			pointsError.Wrap("audStartPoints")
			pointsError.Err = errors.New("the dots are in the audience area")
			return *pointsError
		}
	}
}