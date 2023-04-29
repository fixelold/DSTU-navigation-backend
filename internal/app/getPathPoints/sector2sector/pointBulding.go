package sectorToSector

import (
	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (s *sectorToSectorController) building(iterator int, borderSector models.Coordinates) appError.AppError {
	boolean := true
	repository := NewRepository(s.client, s.logger)
	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height, s.constData.axisX, s.constData.axisY)
	for boolean {
		if s.checkOccurrence(s.points[iterator], axis, borderSector) {

			s.pathAlignment(borderSector, axis)

			axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)

			points := s.preparation(axis, borderSector, s.points[iterator])

			points = s.setPoints(borderSector, points, s.points[iterator], axis)

			s.points = append(s.points, points)
			boolean = false
		} else {

			points := s.preparation(axis, borderSector, s.points[iterator])

			points = s.setPoints(borderSector, points, s.points[iterator], axis)

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

			s.points = append(s.points, points)
		}

		iterator += 1
	}

	return appError.AppError{}
}

// точки от начала пути до вхождение в пределы сектора
func (s *sectorToSectorController) setPoints(borderPoints, points, lastPathPoint models.Coordinates, axis int) (models.Coordinates) {
	p := models.Coordinates{
		X: (points.X),
		Y: (points.Y)}
	p.Widht = points.Widht
	p.Height = points.Height
	return p
}


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
	lenght := len(s.points)
	path := s.points[lenght-1]
	switch axis {
	case s.constData.axisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			points.Widht = sectorPoints - path.X
			points.Height = s.constData.heightX
			s.points[lenght-1].Widht = points.Widht
		} else if sectorPoints < path.X {
			points.Widht = sectorPoints - path.X
			points.Height = s.constData.heightX
			s.points[lenght-1].Widht = points.Widht
		}
	case s.constData.axisY:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		if sectorPoints > path.Y {
			points.Widht = s.constData.widhtY
			points.Height = sectorPoints - path.Y
			s.points[lenght-1].Height = points.Height
		} else if sectorPoints < path.Y {
			points.Widht = s.constData.widhtY
			points.Height = sectorPoints - path.Y
			s.points[lenght-1].Height = points.Height
		}
	default:
		s.logger.Errorln("Path Alignment default")
	}
}
