package middle

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	switch axis {
	case m.constData.axisX:
		var factorX int
		var factorBorderX int
		var factorY int
		var initFactorY = 1

		if borderPoint.X > points.X {
			factorX = 0 
			factorBorderX = 10
		} else {
			factorX = 1
			factorBorderX = 10 // тут было -10
		}

		if points.Height == m.constData.heightX || points.Height == -m.constData.heightX { // возможно тут над будет изменить
			factorX = 1
			initFactorY = 0
			factorBorderX = -15
		}

		if points.Widht == -m.constData.widhtY {
			factorX = 0
		}

		if points.Height > 0 {
			factorY=1
		} else {
			factorY=-1
		}

		result := models.Coordinates{
			X: points.X + (points.Widht * factorX),
			Y: points.Y + (points.Height * initFactorY),
			Widht: (borderPoint.X - points.X + (points.Widht * factorX)) + factorBorderX,
			Height: m.constData.heightX * factorY,
		}

		if borderPoint.X < points.X && result.X + result.Widht != ((borderPoint.X + borderPoint.Widht) - 10) {
			result.Widht = result.Widht - 5
		}

		return result, appError.AppError{}
	case m.constData.axisY:
		var factorY int
		var factorX int
		var factorBorderY int
		var factorFinal int

		if borderPoint.Y > points.Y {
			factorY = 0
			factorBorderY = 10
		} else {
			factorBorderY = -10 // тут было -10
			factorY = 1
		}

		if points.Widht > 0 {factorX=1} else {factorX=-1}

		result := models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y + (points.Height * factorY * factorFinal),
			Widht: m.constData.widhtY * factorX,
			Height: (borderPoint.Y - points.Y + (points.Height * factorY)) + factorBorderY,
		}

		// TODO: надо посмотреть про высоты отрисовки.
		if borderPoint.Y < points.Y && result.Y + result.Height != ((borderPoint.Y + borderPoint.Height) - 10) {
			result.Height = result.Height + 20
		}

		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}

func (m *middleController) finalPreparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	if axis == m.constData.axisX {
		if len(m.Points) == 1 {
			path, err := m.leftAndRightX(borderPoint, points)
			if err.Err != nil {
				
			}
		}
}