package middle

import (
	"fmt"

	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	// var factor int // необходимо для рпавильного расчета высоты
	// var factorForY int
	// switch axis {
	// case m.constData.axisX:
	// 	fmt.Println("middle work 1")
	// 	// 1-420 - лестница
	// 	if borderPoint.X > points.X {
	// 		return models.Coordinates{
	// 			X:      points.X, // Лестница(4) - 1-174
	// 			Y:      points.Y + points.Height,
	// 			Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X, // Лестница(4) - 1-174
	// 			Height: 5,
	// 		}, appError.AppError{}
	
	// 	}
	// 	return models.Coordinates{
	// 		X:      points.X + points.Widht, // Лестница(4) - 1-174
	// 		Y:      points.Y + points.Height,
	// 		Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X - points.Widht, // Лестница(4) - 1-174
	// 		Height: -5,
	// 	}, appError.AppError{}


	// case m.constData.axisY:
	// 	fmt.Println("middle work 2")
	// 	if points.Widht > 0 {
	// 		factor = -1
	// 	} else {
	// 		factor = 1
	// 	}

	// 	if m.typeTransition >= 2 {
	// 		fmt.Println("middle work 3")
	// 		if (m.thisSectorNumber % 10) == (m.sectorNumber % 10) && (m.sectorNumber % 10) == 2 && borderPoint.X > points.X {
	// 			// 1-340 - лестница
	// 			return models.Coordinates{
	// 				X:      points.X + points.Widht,
	// 				Y:      points.Y + points.Height,
	// 				Widht:  -m.constData.widhtY,
	// 				Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height - finalHeight),
	// 			}, appError.AppError{}
	// 		}
	// 	}
	// 	if borderPoint.Y > points.Y {
	// 		fmt.Println("middle work 4")
	// 		return models.Coordinates{
	// 			X:      points.X + points.Widht,
	// 			Y:      points.Y,
	// 			Widht:  -m.constData.widhtY * factor,
	// 			Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height) + m.constData.widhtY,
	// 		}, appError.AppError{}
	// 	}
	// 	// } else if borderPoint.Y < points.Y {
	// 	// 	return models.Coordinates{
	// 	// 		X:      points.X + points.Widht,
	// 	// 		Y:      points.Y,
	// 	// 		Widht:  -m.constData.widhtY * factor,
	// 	// 		Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height) + m.constData.widhtY,
	// 	// 	}, appError.AppError{}
	// 	// }
	// 	fmt.Println("Work Omega")
	// 	return models.Coordinates{
	// 		X:      points.X + points.Widht,
	// 		Y:      points.Y + points.Height,
	// 		Widht:  -m.constData.widhtY * factor,
	// 		Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height), // 1-399 - лестница
	// 	}, appError.AppError{}
	// default:
	// 	return models.Coordinates{}, *appError.NewAppError("switch error")
	// }

	switch axis {
	case m.constData.axisX:
		var factorX int
		var factorBorderX int
		var factorY int

		if borderPoint.X > points.X {
			factorX = 0 
			factorBorderX = 10
		} else {
			factorX = 1
			factorBorderX = 10 // тут было -10
		}

		if points.Height > 0 {factorY=1} else {factorY=-1}
		result := models.Coordinates{
			X: points.X + (points.Widht * factorX),
			Y: points.Y + points.Height,
			Widht: (borderPoint.X - points.X + (points.Widht * factorX)) + factorBorderX ,
			Height: m.constData.heightX * factorY,
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
		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}

func (m *middleController) finalPreparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	switch axis {
	case m.constData.axisX:
		result := models.Coordinates{}
		return result, appError.AppError{}
	case m.constData.axisY:
		var factorLenPath int
		var faactorX = 1
		var factorXWidht int

		if len(m.Points) == 1 {
			if points.Height != m.constData.heightX && points.Height != -m.constData.heightX {
				faactorX = 0} else {faactorX = 1}
			factorLenPath = 1
		} else {
			factorLenPath = 0
		}

		if points.Widht > 0 {factorXWidht=1} else {factorXWidht=-1}
		result := models.Coordinates{
			X: points.X + (points.Widht * faactorX),
			Y: points.Y + (points.Height * factorLenPath),
			Widht: m.constData.widhtY * factorXWidht,
			Height: borderPoint.Y - (points.Y + (points.Height * factorLenPath)),
		}
		fmt.Println("sector: ", borderPoint)
		fmt.Println("poits: ", points)
		fmt.Println("result: ", result)
		return result, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}