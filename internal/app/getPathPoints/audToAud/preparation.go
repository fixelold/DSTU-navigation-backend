package audToAud

import (
	"fmt"

	"navigation/internal/appError"
	"navigation/internal/models"
)

func (s *audToAudController) startPreparation(axis int, borderPoint models.Coordinates) models.Coordinates {
	var coordinates models.Coordinates

	XX := borderPoint.X
	YX := (borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2

	XY := (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2
	YY := borderPoint.Y

	if axis == s.constData.axisX {
		coordinates.X = XX
		coordinates.Y = YX
		coordinates.Widht = s.constData.widhtX
		coordinates.Height = s.constData.heightX

	} else if axis == s.constData.axisY {
		coordinates.X = XY
		coordinates.Y = YY
		coordinates.Widht = s.constData.widhtY
		coordinates.Height = s.constData.heightY
	}

	return coordinates
}


func (a *audToAudController) middlePreparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	switch axis {
	case a.constData.axisX:
		var factorX int
		var factorBorderX int
		var factorY int
		var initFactorY = 1

		if borderPoint.X > points.X {
			factorX = 0 
			factorBorderX = 10
		} else {
			factorX = 1
			factorBorderX = borderPoint.Widht - (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2 // тут было -10
		}

		fmt.Println("Work: ", (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2, borderPoint)
		if points.Height == a.constData.heightX || points.Height == -a.constData.heightX { // возможно тут над будет изменить
			factorX = 1
			initFactorY = 0
			factorBorderX = -15
		}

		if points.Widht == -a.constData.widhtY {
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
			Height: a.constData.heightX * factorY,
		}

		if borderPoint.X < points.X && result.X + result.Widht != ((borderPoint.X + borderPoint.Widht) - 10) {
			result.Widht = result.Widht - 5
		}

		return result, appError.AppError{}
	case a.constData.axisY:
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

		if borderPoint.Widht == 1 {
			factorBorderY = 10
		}

		result := models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y + (points.Height * factorY * factorFinal),
			Widht: a.constData.widhtY * factorX,
			Height: borderPoint.Y - points.Y + factorBorderY,
		}

		// TODO: надо посмотреть про высоты отрисовки.
		// if borderPoint.Y < points.Y && result.Y + result.Height != ((borderPoint.Y + borderPoint.Height) - 10) {
		// 	result.Height = result.Height + 10
		// }
		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}

func (a *audToAudController) middleFinalPreparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	switch axis {
	case a.constData.axisX:
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

		if points.Height == a.constData.heightX || points.Height == -a.constData.heightX { // возможно тут над будет изменить
			factorX = 1
			initFactorY = 0
		}

		if points.Widht == -a.constData.widhtY {
			factorX = 0
		}

		result := models.Coordinates{
			X: points.X + (points.Widht * factorX),
			Y: points.Y + (points.Height * initFactorY),
			Widht: borderPoint.X - (points.X + (points.Widht * factorX)),
			Height: a.constData.heightX * factorY,
		}

		if borderPoint.X < points.X && result.X + result.Widht != ((borderPoint.X + borderPoint.Widht) - 10) {
			result.Widht = result.Widht - 5
		}
		return result, appError.AppError{}
	case a.constData.axisY:
		var factorLenPath int
		var faactorX = 1
		var factorXWidht int

		if len(a.points) == 1 {
			if points.Height != a.constData.heightX && points.Height != -a.constData.heightX {
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

		if len(a.points) != 1 && a.points[0].Height == a.constData.heightX || a.points[0].Height == -a.constData.heightX { // возможно тут над будет изменить
			if a.points[1].Height == a.constData.heightX || a.points[1].Height == -a.constData.heightX {
				factorLenPath = 0
			}
		}

		if points.Widht > 0 {factorXWidht=1} else {factorXWidht=-1}
		result := models.Coordinates{
			X: points.X + (points.Widht * faactorX),
			Y: points.Y + (points.Height * factorLenPath),
			Widht: a.constData.widhtY * factorXWidht,
			Height: borderPoint.Y - (points.Y + (points.Height * factorLenPath)),
		}
		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}