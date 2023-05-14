package middle

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	var factor int // необходимо для рпавильного расчета высоты
	switch axis {
	case m.constData.axisX:
		if points.Height > 0 {
			factor = -1
		} else {
			factor = 1
		}
		return models.Coordinates{
			X:      points.X,
			Y:      points.Y + points.Height,
			Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X,
			Height: m.constData.heightX * factor,
		}, appError.AppError{}


	case m.constData.axisY:
		if points.Widht > 0 {
			factor = -1
		} else {
			factor = 1
		}
		return models.Coordinates{
			X:      points.X + points.Widht,
			Y:      points.Y + points.Height,
			Widht:  m.constData.widhtY * factor,
			Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height),
		}, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}