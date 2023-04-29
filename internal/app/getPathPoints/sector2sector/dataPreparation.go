package sectorToSector

import (
	"navigation/internal/models"
)

func (s *sectorToSectorController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	if axis == s.constData.axisX {
		return models.Coordinates{
			X: points.X + points.Widht,
			Y: points.Y + points.Height,
			Widht: borderPoint.X - (points.X + points.Widht),
			Height: s.constData.heightX,
		}

	} else {
		return models.Coordinates{
			X: points.X,
			Y: points.Y + points.Height,
			Widht: s.constData.widhtY,
			Height: borderPoint.Y - (points.Y + points.Height),
		}
	}
}
