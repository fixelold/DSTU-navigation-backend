package sectorToSector

import (
	"fmt"

	"navigation/internal/models"
)

func (s *sectorToSectorController) preparation(axis int, borderPoint, points models.Coordinates, lastPathSector bool) models.Coordinates {
	if lastPathSector {
		fmt.Println("Work 2")
		if axis == s.constData.axisX {
			return models.Coordinates{
				X: points.X + points.Widht,
				Y: points.Y + points.Height,
				Widht: borderPoint.X - (points.X + points.Widht),
				Height: s.constData.heightX,
			}
	
		} else {
			return models.Coordinates{
				X: points.X + points.Widht,
				Y: points.Y + points.Height,
				Widht: s.constData.widhtY,
				Height: borderPoint.Y - (points.Y + points.Height),
			}
		}
	} else {
		fmt.Println("Work 1: ", points.X + points.Widht)
		if axis == s.constData.axisX {
			return models.Coordinates{
				X: points.X + points.Widht,
				Y: points.Y,
				Widht: borderPoint.X - (points.X + points.Widht - 10),
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

}