package sectorToSector

import (
	"fmt"

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
	switch axis {
	case s.constData.axisX:
		lenght := len(s.Points)
		var factorX int

		if points.Widht != -5 && points.Widht != 5 {
			factorX = 0
		} else {
			factorX = 1
		}
		result := models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y + (points.Height * factorX),
			Widht: borderPoint.X - (points.X + points.Widht),
			Height: -s.constData.heightX,
		}

		if points.Height != 5 && points.Height != -5 && borderPoint.Y < points.Y {
			s.Points[lenght - 1].Height -= 5
		}
		return result
	case s.constData.axisY:
		fmt.Println("Work")
		// var factorHeight int
		lenght := len(s.Points)
		var factorWidht int
		var factorX int // если путь идет справа сверху

		if (points.X > borderPoint.X || points.X < borderPoint.X) && (points.Y < borderPoint.Y || points.Y > borderPoint.Y) {
			if points.Y > borderPoint.Y  {
				s.Points[lenght - 1].Widht -= 5
			}
			factorX = 1
		} else {
			factorX = 0
		}

		if points.Widht > 0 && borderPoint.Y < points.Y {
			factorWidht = 1
		} else {
			factorWidht = -1
		}

		result := models.Coordinates {
			X: points.X + (points.Widht * factorX),
			Y: points.Y + points.Height,
			Widht: s.constData.widhtY * factorWidht,
			Height: borderPoint.Y - (points.Y + points.Height),
		}
		return result
	default:
		return models.Coordinates{}
	}
}