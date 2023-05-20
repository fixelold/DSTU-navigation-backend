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
			result.Height = result.Height + 10
		}
		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}

func (m *middleController) finalPreparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	switch axis {
	case m.constData.axisX:
		var factorX int
		var factorY int
		var initFactorY = 1

		if borderPoint.X > points.X {
			factorX = 0 
		} else {
			factorX = 1
		}

		if points.Height > 0 {
			factorY=1
		} else {
			factorY=-1
		}

		if points.Height == m.constData.heightX || points.Height == -m.constData.heightX { // возможно тут над будет изменить
			factorX = 1
			initFactorY = 0
		}

		if points.Widht == -m.constData.widhtY {
			factorX = 0
		}

		result := models.Coordinates{
			X: points.X + (points.Widht * factorX),
			Y: points.Y + (points.Height * initFactorY),
			Widht: borderPoint.X - (points.X + (points.Widht * factorX)),
			Height: m.constData.heightX * factorY,
		}

		if borderPoint.X < points.X && result.X + result.Widht != ((borderPoint.X + borderPoint.Widht) - 10) {
			result.Widht = result.Widht - 5
		}
		return result, appError.AppError{}
	case m.constData.axisY:
		var factorLenPath int
		var faactorX = 1
		var factorXWidht int

		if len(m.Points) == 1 {
			if points.Height != m.constData.heightX && points.Height != -m.constData.heightX {
				faactorX = 0

				} else {
					faactorX = 1
				}
			factorLenPath = 1
		} else if points.Height > 0 || borderPoint.Y > points.Y {
			factorLenPath = 1
		} else {
			factorLenPath = 0
		}

		if len(m.Points) != 1 && m.Points[0].Height == m.constData.heightX || m.Points[0].Height == -m.constData.heightX { // возможно тут над будет изменить
			if m.Points[1].Height == m.constData.heightX || m.Points[1].Height == -m.constData.heightX {
				factorLenPath = 0
			}
		}

		if points.Widht > 0 {factorXWidht=1} else {factorXWidht=-1}
		result := models.Coordinates{
			X: points.X + (points.Widht * faactorX),
			Y: points.Y + (points.Height * factorLenPath),
			Widht: m.constData.widhtY * factorXWidht,
			Height: borderPoint.Y - (points.Y + (points.Height * factorLenPath)),
		}
		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}