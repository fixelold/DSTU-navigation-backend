package sectorToSector

import (
	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)
 
func (s *sectorToSectorController) building(iterator int, borderSector models.Coordinates) appError.AppError {
	// repository := NewRepository(s.client, s.logger) // для обращение к базе данных
	// ось для перехода в другой сектор
	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height, s.constData.axisX, s.constData.axisY)
	// var b = true
	// i := 0
	for true {
		// if i == 3 {
		// 	break
		// } 
		// i += 1
		// проверка вхождение координат пути в координаты границ сектора
		if s.checkOccurrence(s.Points[iterator], axis, borderSector) {
			var points models.Coordinates
			axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)
			points = s.finalPreparation(axis, borderSector, s.Points[iterator])

			// if (axis == s.constData.axisX && s.Points[iterator].Widht == 5) || (axis == s.constData.axisY && s.Points[iterator].Height == 5)  {
			// 	fmt.Println("Work")
			// 	axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)
			// 	b = true
			// 	points = s.finalPreparation(axis, borderSector, s.Points[iterator])
			// } else {
			// 	axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)
			// 	b = false
			// 	points = s.finalPreparation(axis, borderSector, s.Points[iterator])
			// }

			s.Points = append(s.Points, points)
			break
		} 
		// // расчет точек пути
		// points := s.preparation(axis, borderSector, s.Points[iterator])


		// // TODO: просмотреть проверку ругался на 1-408.
		// ok, err := repository.checkBorderAud2(points, s.thisSectorNumber)
		// if err.Err != nil {
		// 	err.Wrap("building")
		// 	return err
		// }

		// // изменения оси построения, если точки входят в пределы аудитории
		// if !ok {
		// 	axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)
		// 	points = s.preparation(axis, borderSector, s.Points[iterator])
		// 	axis = axes.ChangeAxis(axis, s.constData.axisX, s.constData.axisY)
		// }
		// s.Points = append(s.Points, points)
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
