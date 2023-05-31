package sectorToSector

import (
	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)
 
func (s *sectorToSectorController) building(iterator int, borderSector models.Coordinates) appError.AppError {
	// ось для перехода в другой сектор
	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height, s.constData.axisX, s.constData.axisY)
	var b = false
	for i := 0; true; i++ {
		// if i == 1 {
		// 	break
		// } 
		// проверка вхождение координат пути в координаты границ сектора
		lenght := len(s.Points)
		if s.checkOccurrence(s.Points[lenght - 1], axis, borderSector) {
			var points models.Coordinates
			axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)
			points = s.finalPreparation(axis, borderSector, s.Points[lenght - 1], b)

			s.Points = append(s.Points, points)
			break
		} 
		// расчет точек пути
		points := s.preparation(axis, borderSector, s.Points[lenght - 1])
		b = true

		s.Points = append(s.Points, points)
	}

	return appError.AppError{}
}

// точки от начала пути до вхождение в пределы сектора


// проверка на вхождение точек пути в пределы сектора.
func (s *sectorToSectorController) checkOccurrence(points models.Coordinates, axis int, borderSector models.Coordinates) bool {
	switch axis {
	case s.constData.axisX:
		ph := points.X + points.Widht
		x1 := borderSector.X
		x2 := borderSector.X + borderSector.Widht
		if x1 <= ph && ph <= x2 {
			return true
		} else {
			return false
		}
	case s.constData.axisY:
		ph := points.Y + points.Height
		y1 := borderSector.Y
		y2 := borderSector.Y + borderSector.Height
		if y1 <= ph && ph <= y2 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

// выравнивание пути
func (s *sectorToSectorController) pathAlignment(sectorBorderPoint models.Coordinates, axis int) {
	lenght := len(s.Points)
	path := s.Points[lenght-1]
	switch axis {
	case s.constData.axisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			// if temo == 0 {
			// 	s.Points[lenght-2] = 
			// }
			points.Widht = sectorPoints - path.X
			points.Height = s.constData.heightX
			s.Points[lenght-1].Widht = points.Widht
		} else if sectorPoints < path.X {
			points.Widht = (sectorBorderPoint.X + (sectorBorderPoint.Widht / 2)) - (path.X + s.constData.widhtY)
			points.Height = s.constData.heightX
			s.Points[lenght-1].Widht = points.Widht
		}
	case s.constData.axisY:
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
	

		s.Points[lenght-1].Widht = s.constData.widhtY
		s.Points[lenght-1].Height = sectorPoints - s.Points[lenght-1].Y
	default:
		s.logger.Errorln("Path Alignment default")
	}
}
