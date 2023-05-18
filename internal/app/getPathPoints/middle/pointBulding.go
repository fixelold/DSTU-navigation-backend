package middle

import (
	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (m *middleController) building(borderSector models.Coordinates) appError.AppError {
	repository := NewRepository(m.client, m.logger) // для обращение к базе данных
	// ось для перехода в другой сектор
	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height, m.constData.axisX, m.constData.axisY)

	for i := 0; true; i++ {
		// if i != 1 {
		// 	break
		// } 
		// проверка вхождение координат пути в координаты границ сектора
		if m.checkOccurrence(m.Points[i], axis, borderSector) {
			axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)

			m.pathAlignment(borderSector, axis)
			
			// расчет точек пути
			finalHeight := -m.constData.heightX // 1-340 - лестница
			points, err := m.preparation(axis, borderSector, m.Points[i], finalHeight)
			if err.Err != nil {
				err.Wrap("building")
				return err
			}
			m.Points = append(m.Points, points)
			break
		} 
		// расчет точек пути
		points, err := m.preparation(axis, borderSector, m.Points[i], m.constData.heightY)
		if err.Err != nil {
			err.Wrap("building")
			return err
		}


		ok2, err := repository.checkBorderAud2(points, m.thisSectorNumber)
		if err.Err != nil {
			err.Wrap("building")
			return err
		}

		// изменения оси построения, если точки входят в пределы аудитории
		if !ok2 {
			axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)
			points, err = m.preparation(axis, borderSector, m.Points[i], m.constData.heightY)
			if err.Err != nil {
				err.Wrap("building")
				return err
			}
			axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)
		}
		m.Points = append(m.Points, points)
	}

	return appError.AppError{}
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
	lenght := len(m.Points)
	switch axis {
	case m.constData.axisX:
		m.Points[lenght-1].Height = (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2 - (m.Points[lenght-1].Y)
	case m.constData.axisY:
		if m.typeTransition >= 2 && sectorBorderPoint.Y == m.Points[lenght-1].Y {
			m.Points[lenght-2].Height = m.Points[lenght-2].Height + 10 // TODO: тута надо сделать обратку и для - 10
			m.Points[lenght-1].Y = m.Points[lenght-1].Y + 10 // TODO: тоже самое
		} else {
			// надо для того, чтобы еще аудитория находится прямо возле лестници, чтобы не было не красиво
			// len(m.Points) > 1 
			if len(m.Points) >= 1 {
				var result int
				if sectorBorderPoint.X > m.Points[lenght-1].X {
					result = (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2 - (m.Points[lenght-1].X) // от лестницы (143) до сектора 142
				} else {
					result = (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2 - (m.Points[lenght-1].X) - m.constData.widhtY // - m.constData.widhtY Лестница(122) к 121
				}
				if result == -25 { // это надо если аудитория находится под лестницей

				} else {
					m.Points[lenght-1].Widht = result
				}
			}
		}
	default:
		m.logger.Errorln("Path Alignment default")
	}
}
