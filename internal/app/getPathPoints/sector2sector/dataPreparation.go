package sectorToSector

import (
	"navigation/internal/models"
)

func (s *sectorToSectorController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	switch axis {
	case s.constData.axisX:
		
		var factorX int
		var lenPath int

		if points.X > borderPoint.X {
			lenPath = 15
		} else {
			lenPath = 15
		}

		if points.Widht != -5 && points.Widht != 5 {
			factorX = 0
		} else {
			factorX = 1
		}

		if points.Height == -s.constData.axisX {
			factorX = 1
		}

		result := models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y + (points.Height * factorX),
			Widht: borderPoint.X - (points.X + points.Widht) + lenPath,
			Height: -s.constData.heightX,
		}


		return result
	case s.constData.axisY:
		var factorHeight = 15
		var factorWidht int

		if points.Widht > 0 {
			factorWidht = 1
		} else {
			factorWidht = -1
		}

		// if borderPoint.Y < points.Y {
		// 	factorHeight = 10
		// } else {
		// 	factorHeight = 15
		// }

		result := models.Coordinates {
			X: points.X,
			Y: points.Y + points.Height,
			Widht: s.constData.widhtY * factorWidht,
			Height: borderPoint.Y - (points.Y + points.Height) + factorHeight,
		}
		return result
	default:
		return models.Coordinates{}
	}
}

func (s *sectorToSectorController) finalPreparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	var path models.Coordinates

	if axis == s.constData.axisX {
		path = s.finalX(borderPoint, points)
	} else if axis == s.constData.axisY {
		path = s.finalY(borderPoint, points)
	}

	return path
}