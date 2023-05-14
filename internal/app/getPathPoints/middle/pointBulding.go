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
		if i == 1 {
			break
		}
		// проверка вхождение координат пути в координаты границ сектора
		if m.checkOccurrence(m.Points[i], axis, borderSector) {
			// TODO: отрисовка пути прямо до границ секторного перехода
			break
		}

		// расчет точек пути
		points, err := m.preparation(axis, borderSector, m.Points[i])
		if err.Err != nil {
			err.Wrap("building")
			return err
		}

		// проверка точек на вхождение в пределы аудитории
		ok, err := repository.checkBorderAud(points, m.thisSectorNumber)
		if err.Err != nil {
			err.Wrap("building")
			return err
		}

		ok2, err := repository.checkBorderAud(points, m.thisSectorNumber)
		if err.Err != nil {
			err.Wrap("building")
			return err
		}

		// изменения оси построения, если точки входят в пределы аудитории
		if !ok || !ok2 {
			axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)
			points, err = m.preparation(axis, borderSector, m.Points[i])
			if err.Err != nil {
				err.Wrap("building")
				return err
			}
		}

		m.Points = append(m.Points, points)
	}

	return appError.AppError{}
}

// func (m *middleController) building(borderSector models.Coordinates) appError.AppError {
// 	boolean := true
// 	iterator := 0
// 	repository := NewRepository(m.client, m.logger)
// 	axis := axes.DefenitionAxis(borderSector.Widht, borderSector.Height, m.constData.axisX, m.constData.axisY)
// 	ok := true
// 	var err appError.AppError
// 	for boolean {
// 		if m.checkOccurrence(m.Points[iterator], axis, borderSector) {

// 			m.pathAlignment(borderSector, axis)

// 			axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)

// 			points := m.preparation(axis, borderSector, m.Points[iterator], 0)

// 			points = m.setPoints(borderSector, points, m.Points[iterator], axis)

// 			m.Points = append(m.Points, points)
// 			boolean = false
// 		} else {

// 			points := m.preparation(axis, borderSector, m.Points[iterator], 0)

// 			points = m.setPoints(borderSector, points, m.Points[iterator], axis)

// 			if m.typeTransition == 2 {
// 				if points.Widht < 0 || points.Height < 0 {
// 					chechPoints := models.Coordinates{
// 						X: (points.X + points.Widht),
// 						Y: points.Y,
// 						Widht: points.X,
// 						Height: points.Y + points.Height,
// 					}
// 					ok, err = repository.checkBorderAud(chechPoints, m.thisSectorNumber)
// 					if err.Err != nil {
// 						err.Wrap("otherPathPoints")
// 						return err
// 					}
// 				} else {
// 					if axis == m.constData.axisY {
// 						ok, err = repository.checkBorderAudY(points, m.thisSectorNumber)
// 						if err.Err != nil {
// 							err.Wrap("otherPathPoints")
// 							return err
// 						}
// 					} else {
// 						chechPoints := models.Coordinates{
// 							X: points.X ,
// 							Y: points.Y,
// 							Widht: points.X + points.Widht,
// 							Height: points.Y + points.Height,
// 						}
// 						ok, err = repository.checkBorderAud(chechPoints, m.thisSectorNumber)
// 						if err.Err != nil {
// 							err.Wrap("otherPathPoints")
// 							return err
// 						}
	
// 						if ok {
// 							ok, err = repository.checkBorderAud2(chechPoints, m.thisSectorNumber)
// 							if err.Err != nil {
// 								err.Wrap("otherPathPoints")
// 								return err
// 							}
// 						}
// 					}
// 				}
// 			}

// 			ok2, err := repository.checkBorderSector(points)
// 			if err.Err != nil {
// 				err.Wrap("otherPathPoints")
// 				return err
// 			}

// 			if !ok && !ok2 {
// 				//TODO написать изменения направления или типо что-то такого
// 			}

// 			if !ok {
// 				// tmp := m.Points[0]
// 				// m.Points = append(m.Points[0:], tmp)
// 				axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)

// 				points = m.preparation(axis, borderSector, m.Points[iterator], m.constData.heightY)
// 				points = m.setPoints(borderSector, points, m.Points[iterator], axis)

// 				axis = axes.ChangeAxis(axis, m.constData.axisX, m.constData.axisY)
// 			}

// 			m.Points = append(m.Points, points)
// 		}

// 		iterator += 1
// 	}

// 	return appError.AppError{}
// }

// // точки от начала пути до вхождение в пределы сектора
// func (m *middleController) setPoints(borderPoints, points, lastPathPoint models.Coordinates, axis int) (models.Coordinates) {
// 	p := models.Coordinates{
// 		X: (points.X),
// 		Y: (points.Y)}
// 	p.Widht = points.Widht
// 	p.Height = points.Height
// 	return p
// }


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
	path := m.Points[lenght-1]
	switch axis {
	case m.constData.axisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			points.Widht = sectorPoints - path.X
			points.Height = m.constData.heightX
			m.Points[lenght-1].Widht = points.Widht
		} else if sectorPoints < path.X {
			points.Widht = sectorPoints - path.X
			points.Height = m.constData.heightX
			m.Points[lenght-1].Widht = points.Widht
		}
	case m.constData.axisY:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		if sectorPoints > path.Y {
			points.Widht = m.constData.widhtY
			points.Height = sectorPoints - path.Y
			m.Points[lenght-1].Height = points.Height
		} else if sectorPoints < path.Y {
			points.Widht = m.constData.widhtY
			points.Height = sectorPoints - path.Y
			m.Points[lenght-1].Height = points.Height
		}
	default:
		m.logger.Errorln("Path Alignment default")
	}
}
