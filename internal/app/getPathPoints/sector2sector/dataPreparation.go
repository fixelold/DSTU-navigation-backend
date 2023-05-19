package sectorToSector

import (
	"navigation/internal/models"
)

func (s *sectorToSectorController) preparation(axis int, borderPoint, points models.Coordinates, lastPathSector bool) models.Coordinates {
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

	var factorWidht int

	if borderPoint.Y < points.Y {
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
}