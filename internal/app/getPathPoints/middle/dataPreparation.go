package middle

import (
	"fmt"

	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) (models.Coordinates, appError.AppError) {
	var factor int // необходимо для рпавильного расчета высоты

	// var factor2 int // необходимо для рпавильного расчета высоты
	// if points.X > borderPoint.X {
	// 	factor2 = -1
	// } else {
	// 	factor2 = 1
	// }
	
	switch axis {
	case m.constData.axisX:
		if points.Height > 0 {
			factor = -1
		} else {
			factor = 1
		}
		fmt.Println("points: ", points, borderPoint)
		return models.Coordinates{
			X:      points.X,
			Y:      points.Y + points.Height,
			Widht:  (borderPoint.X + (borderPoint.Widht / 2)) - points.X,
			Height: m.constData.heightX * factor,
		}, appError.AppError{}
		// return models.Coordinates {
		// 	X: points.X + points.Widht,
		// 	Y: (points.Y + points.Height) + (m.constData.heightX * factor),
		// 	Widht: m.constData.widhtX * factor2,
		// 	Height: m.constData.heightX,
		// }, appError.AppError{}

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

// func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
// 	if axis == m.constData.axisX {
// 		if len(strconv.Itoa(m.sectorNumber)) == 4 { //stairs
// 			return models.Coordinates{
// 				X:      points.X + points.Widht,
// 				Y:      points.Y + points.Height,
// 				Widht:  borderPoint.X - points.X,
// 				Height: m.constData.heightX}
// 		} else {
// 			if borderPoint.X < points.X {
// 				return models.Coordinates{
// 					X:      points.X + points.Widht,
// 					Y:      points.Y + points.Height,
// 					Widht:  borderPoint.X - points.X,
// 					Height: m.constData.heightX}
// 			} else {
// 				return models.Coordinates{
// 					X:      points.X + points.Widht,
// 					Y:      points.Y + points.Height,
// 					Widht:  borderPoint.X - points.X,
// 					Height: m.constData.heightX}
// 			}
// 		}
// 	} else {
// 		if borderPoint.Y > points.Y {
// 			return models.Coordinates{
// 				X:      points.X + points.Widht,
// 				Y:      points.Y + points.Height - m.constData.heightX,
// 				Widht:  m.constData.widhtY,
// 				Height: borderPoint.Y - points.Y + temp}
// 		} else {
// 			return models.Coordinates{
// 				X:      points.X + points.Widht,
// 				Y:      points.Y + points.Height,
// 				Widht:  m.constData.widhtY,
// 				Height: borderPoint.Y - (points.Y + points.Height)}
// 		}
// 	}
// }