package middle

import (
	"fmt"

	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates, finalHeight int) (models.Coordinates, appError.AppError) {
	var factor int // необходимо для рпавильного расчета высоты
	// var factorForY int
	switch axis {
	case m.constData.axisX:
		fmt.Println("middle work 1")
		// 1-420 - лестница
		if borderPoint.X > points.X {
			return models.Coordinates{
				X:      points.X, // Лестница(4) - 1-174
				Y:      points.Y + points.Height,
				Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X, // Лестница(4) - 1-174
				Height: -5,
			}, appError.AppError{}
	
		}
		return models.Coordinates{
			X:      points.X + points.Widht, // Лестница(4) - 1-174
			Y:      points.Y + points.Height,
			Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X - points.Widht, // Лестница(4) - 1-174
			Height: -5,
		}, appError.AppError{}


	case m.constData.axisY:
		fmt.Println("middle work 2")
		if points.Widht > 0 {
			factor = -1
		} else {
			factor = 1
		}

		if m.typeTransition >= 2 {
			fmt.Println("middle work 3")
			if (m.thisSectorNumber % 10) == (m.sectorNumber % 10) && (m.sectorNumber % 10) == 2 && borderPoint.X > points.X {
				// 1-340 - лестница
				return models.Coordinates{
					X:      points.X + points.Widht,
					Y:      points.Y + points.Height,
					Widht:  -m.constData.widhtY,
					Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height - finalHeight),
				}, appError.AppError{}
			}
		}
		if borderPoint.Y > points.Y {
			fmt.Println("middle work 4")
			return models.Coordinates{
				X:      points.X + points.Widht,
				Y:      points.Y,
				Widht:  -m.constData.widhtY * factor,
				Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height) + m.constData.widhtY,
			}, appError.AppError{}
		}
		// } else if borderPoint.Y < points.Y {
		// 	return models.Coordinates{
		// 		X:      points.X + points.Widht,
		// 		Y:      points.Y,
		// 		Widht:  -m.constData.widhtY * factor,
		// 		Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height) + m.constData.widhtY,
		// 	}, appError.AppError{}
		// }
		fmt.Println("Work Omega")
		return models.Coordinates{
			X:      points.X + points.Widht,
			Y:      points.Y + points.Height,
			Widht:  m.constData.widhtY * factor,
			Height: (borderPoint.Y + (borderPoint.Height / 2)) - (points.Y + points.Height), // 1-399 - лестница
		}, appError.AppError{}
	default:
		return models.Coordinates{}, *appError.NewAppError("switch error")
	}
}