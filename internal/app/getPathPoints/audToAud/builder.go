package audToAud

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
func (s *audToAudController) pathBuilding(points models.Coordinates, axis, sign int) (models.Coordinates, appError.AppError) {
	var path models.Coordinates
	switchAxisError.Wrap("setPointsAudStart")
	switchSignError.Wrap("setPointsAudStart")
	switchAxisError.Err = errors.New(fmt.Sprintf("no suitable value, expected: %d or %d received: %d", s.constData.axisX, s.constData.axisY, axis))
	switchSignError.Err = errors.New(fmt.Sprintf("no suitable value, expected: %d or %d received: %d", s.constData.positiveCoordinate, s.constData.negativeCoordinate, sign))
	switch axis {
	case s.constData.axisX:
		switch sign {
		case s.constData.positiveCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case s.constData.negativeCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = -points.Widht
			path.Height = points.Height

		default:
			return path, *switchSignError
		}
	case s.constData.axisY:
		switch sign {
		case s.constData.positiveCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case s.constData.negativeCoordinate:
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