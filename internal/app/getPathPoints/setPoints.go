package getPathPoints

import (
	"errors"
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/models"
)

var (
	switchAxisError = appError.NewAppError("switch error")
	switchSignError = appError.NewAppError("switch error")
)

// для начального пути от границ сектора
func (d *data) setPointsAudStart(points models.Coordinates, axis, sign int) (models.Coordinates, appError.AppError) {
	var path models.Coordinates
	switchAxisError.Wrap("setPointsAudStart")
	switchSignError.Wrap("setPointsAudStart")
	switchAxisError.Err = errors.New(fmt.Sprintf("no suitable value, expected: %d or %d received: %d", AxisX, AxisY, axis))
	switchSignError.Err = errors.New(fmt.Sprintf("no suitable value, expected: %d or %d received: %d", plus, minus, sign))
	switch axis {
	case AxisX:
		switch sign {
		case plus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case minus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = -points.Widht
			path.Height = points.Height

		default:
			return path, *switchSignError
		}
	case AxisY:
		switch sign {
		case plus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case minus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = -points.Height

		default:
			return path, *switchSignError
		}
	default:
		return path, *switchAxisError
	}

	return path, appError.AppError{}
}

// точки от начала пути до вхождение в пределы сектора
func (d *data) setPointsPath2Sector(borderPoints, points, lastPathPoint models.Coordinates, axis int) (models.Coordinates) {
	p := models.Coordinates{
		X: (points.X),
		Y: (points.Y)}
	p.Widht = points.Widht
	p.Height = points.Height
	return p
}
