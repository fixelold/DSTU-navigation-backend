package sectorToSector

import (
	"navigation/internal/models"
)

func (s *sectorToSectorController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	// var factor int
	// // var factorY int
	
	// if borderPoint.X < points.X {
	// 	factor = s.constData.widhtY
	// 	// factorY = -1
	// }
	// if lastPathSector {
	// 	if axis == s.constData.axisX {
	// 		fmt.Println("Work 1")
	// 		return models.Coordinates{
	// 			X: points.X + factor,
	// 			Y: points.Y + points.Height,
	// 			Widht: borderPoint.X - (points.X - points.Widht),
	// 			Height: -s.constData.heightX,
	// 		}
	
	// 	} else {
	// 		fmt.Println("Work 2")
	// 		return models.Coordinates{
	// 			X: points.X + points.Widht - s.constData.widhtY,
	// 			Y: points.Y + points.Height,
	// 			Widht: s.constData.widhtY,
	// 			Height: borderPoint.Y - (points.Y + points.Height),
	// 		}
	// 	}
	// } else {
	// 	//((borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2) > points.Y + points.Height
	// 	if points.Height < 0 {
	// 		factor = 0
	// 	} else {
	// 		factor = 1
	// 	}
	// 	if axis == s.constData.axisX {
	// 		fmt.Println("Work 3")
	// 		return models.Coordinates{
	// 			X: points.X + points.Widht,
	// 			Y: points.Y + (points.Height * factor),
	// 			Widht: borderPoint.X - (points.X + points.Widht - 10),
	// 			Height: -s.constData.heightX,
	// 		}
	
	// 	} else {
	// 		fmt.Println("Work 4: ")
	// 		return models.Coordinates{
	// 			X: points.X,
	// 			Y: points.Y + points.Height,
	// 			Widht: -s.constData.widhtY, // от 143 сектора к сектору 142
	// 			Height: borderPoint.Y - (points.Y + points.Height),
	// 		}
	// 	}
	// }

	switch axis {
	case s.constData.axisX:
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

		if points.Height == s.constData.heightX || points.Height == -s.constData.heightX { // возможно тут над будет изменить
			factorX = 1
			initFactorY = 0
			factorBorderX = -15
		}

		if points.Widht == -s.constData.widhtY {
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
			Widht: (borderPoint.X - points.X + (points.Widht * factorX)) + factorBorderX ,
			Height: s.constData.heightX * factorY,
		}
		return result
	case s.constData.axisY:
		
		var factorWidht int

		if points.Widht > 0 {
			factorWidht = 1
		} else {
			factorWidht = -1
		}
	
		result := models.Coordinates {
			X: points.X,
			Y: points.Y + points.Height,
			Widht: s.constData.widhtY * factorWidht,
			Height: borderPoint.Y - (points.Y + points.Height),
		}
		return result
	default:
		return models.Coordinates{}
	}
}