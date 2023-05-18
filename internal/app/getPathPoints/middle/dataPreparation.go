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
		return models.Coordinates{
			X:      points.X + points.Widht, // Лестница(4) - 1-174
			Y:      points.Y + points.Height,
			Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X - points.Widht, // Лестница(4) - 1-174
			Height: -5,
		}, appError.AppError{}


	case m.constData.axisY:
		if points.Widht > 0 {
			factor = -1
		} else {
			factor = 1
		}

		if m.typeTransition >= 2 {
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
		// if borderPoint.Y < points.Y {
		// 	factorForY = -1
		// } else {
		// 	factorForY = 1
		// }
		fmt.Println("Work")
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