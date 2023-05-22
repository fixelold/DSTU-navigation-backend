package sectorToSector

import (
	"navigation/internal/models"
)

func (s *sectorToSectorController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	var path models.Coordinates

	if axis == s.constData.axisX {
		if borderPoint.X > (points.X + points.Widht) {path = s.prePathRightX(borderPoint, points)
		}else if borderPoint.X < (points.X + points.Widht) {path = s.prePathLeftX(borderPoint, points)}
	
	}else if axis == s.constData.axisY {
		if borderPoint.Y > (points.Y + points.Height) {path = s.prePathDownY(borderPoint, points)
			}else if borderPoint.Y < (points.Y + points.Height) {path = s.prePathUpY(borderPoint, points)}
	}
 
	return path
}

func (s *sectorToSectorController) finalPreparation(axis int, borderPoint, points models.Coordinates, exeption bool) models.Coordinates {
	var path models.Coordinates
	if exeption {
		if axis == s.constData.axisX {
			if points.Height < 0 {
				if points.Widht == 5 {path = s.downLeftX(borderPoint, points)
				} else if points.Widht == -5 {path = s.downRightX(borderPoint, points)}
			
			} else if points.Height > 0 {
				if points.Widht == 5 {path = s.upLeftX(borderPoint, points)
					} else if points.Widht == -5 {path = s.upRightX(borderPoint, points)}	
			}
		} else if axis == s.constData.axisY {
			if points.Widht < 0 {
				if points.Height == 5 {path = s.leftDownY(borderPoint, points)
				} else if points.Height == -5 {path = s.leftUpY(borderPoint, points)}
			
			} else if points.Widht > 0 {
				if points.Height == 5 {path = s.rightDownY(borderPoint, points)
					} else if points.Height == -5 {path = s.rightUpY(borderPoint, points)}	
			}
		} 
	} else {
		if axis == s.constData.axisX {path = s.finalX(borderPoint, points)
		} else if axis == s.constData.axisY {path = s.finalY(borderPoint, points)}

	}
	return path
}