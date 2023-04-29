package middle

import (
	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) building(borderSector models.Coordinates) appError.AppError {
	boolean := true
	iterator := 0
	repository := NewRepository(m.client, m.logger)
	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height)
	for boolean {
		if m.checkOccurrence(m.points[iterator], axis, borderSector) {

			m.pathAlignment(borderSector, axis)

			axis = axes.ChangeAxis(axis)

			points := m.preparation(axis, borderSector, m.points[iterator])

			points = m.setPoints(borderSector, points, m.points[iterator], axis)

			m.points = append(m.points, points)
			boolean = false
		} else {

			points := m.preparation(axis, borderSector, m.points[iterator])

			points = m.setPoints(borderSector, points, m.points[iterator], axis)

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

			m.points = append(m.points, points)
		}

		iterator += 1
	}

	return appError.AppError{}
}

// точки от начала пути до вхождение в пределы сектора
func (m *middleController) setPoints(borderPoints, points, lastPathPoint models.Coordinates, axis int) (models.Coordinates) {
	p := models.Coordinates{
		X: (points.X),
		Y: (points.Y)}
	p.Widht = points.Widht
	p.Height = points.Height
	return p
}


// проверка на вхождение точек пути в пределы сектора.
func (m *middleController) checkOccurrence(points models.Coordinates, axis int, borderSector models.Coordinates) bool {
	switch axis {
	case m.constData.axisX:
		ph := points.X + points.Widht
		x1 := borderSector.X
		x2 := borderSector.X + borderSector.Widht
		if x1 <= ph && ph <= x2 {
			return true
		} else {
			return false
		}
	case m.constData.axisY:
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
func (m *middleController) pathAlignment(sectorBorderPoint models.Coordinates, axis int) {
	lenght := len(m.points)
	path := m.points[lenght-1]
	switch axis {
	case m.constData.axisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			points.Widht = sectorPoints - path.X
			points.Height = m.constData.heightX
			m.points[lenght-1].Widht = points.Widht
		} else if sectorPoints < path.X {
			points.Widht = sectorPoints - path.X
			points.Height = m.constData.heightX
			m.points[lenght-1].Widht = points.Widht
		}
	case m.constData.axisY:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		if sectorPoints > path.Y {
			points.Widht = m.constData.widhtY
			points.Height = sectorPoints - path.Y
			m.points[lenght-1].Height = points.Height
		} else if sectorPoints < path.Y {
			points.Widht = m.constData.widhtY
			points.Height = sectorPoints - path.Y
			m.points[lenght-1].Height = points.Height
		}
	default:
		m.logger.Errorln("Path Alignment default")
	}
}
