package sectorToSector

import (
	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)
 
func (s *sectorToSectorController) building(iterator int, borderSector models.Coordinates) appError.AppError {
	boolean := true
	temp := 0
	repository := NewRepository(s.client, s.logger)
	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height, s.constData.axisX, s.constData.axisY)
	lastPathSector := false
	for boolean {
		// fmt.Println("old data - ", s.Points[iterator])
		// if iterator == 2 {
		// 	boolean = false
		// 	break
		// }
		if s.checkOccurrence(s.Points[iterator], axis, borderSector) {
			var points models.Coordinates
			if temp != 0 {
				s.pathAlignment(borderSector, axis)
			} 

			axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)

			if lastPathSector == true {
				points = s.preparation(axis, borderSector, s.Points[iterator], true)
			} else {
				points = s.preparation(axis, borderSector, s.Points[iterator], false)
			}

			s.Points = append(s.Points, points)
	
			s.OldAxis = axis
			boolean = false
		} else {
			temp += 1
			lastPathSector = true
			points := s.preparation(axis, borderSector, s.Points[iterator], false)

			ok, err := repository.checkBorderAud(points)
			if err.Err != nil {
				err.Wrap("otherPathPoints")
				return err
			}
			ok2, err := repository.checkBorderSector(points)
			if err.Err != nil {
				err.Wrap("otherPathPoints")
				return err
			}

			if !ok && !ok2 {
				//TODO написать изменения направления или типо что-то такого
			}

			s.Points = append(s.Points, points)
		}

		iterator += 1
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
		// если расскоментировать, то надо смотреть на путь от лифта сектора 5 до ауд 1-333а
		// points := models.Coordinates{
		// 	X: (path.X),
		// 	Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		// if sectorPoints > path.Y {
		// 	points.Widht = s.constData.widhtY
		// 	points.Height = sectorPoints - path.Y
		// 	s.Points[lenght-1].Height = points.Height
		// } else if sectorPoints < path.Y {
		// 	points.Widht = s.constData.widhtY
		// 	points.Height = sectorPoints - path.Y
		// 	s.Points[lenght-1].Height = points.Height
		// }

		s.Points[lenght-1].Widht = s.constData.widhtY
		s.Points[lenght-1].Height = sectorPoints - s.Points[lenght-1].Y

		// if sectorBorderPoint.Y > points.Y {
		// 	points.Widht = s.constData.widhtY
		// 	points.Height = 
		// }
	default:
		s.logger.Errorln("Path Alignment default")
	}
}
