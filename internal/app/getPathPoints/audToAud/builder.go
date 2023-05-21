package audToAud

import (
	"errors"
	"fmt"

	axes "navigation/internal/app/getPathPoints/axis"
	"navigation/internal/appError"
	"navigation/internal/models"
)

var (
	switchAxisError = appError.NewAppError("switch error")
	switchSignError = appError.NewAppError("switch error")
)
 
// для начального пути от границ сектора
func (s *audToAudController) startBuilding(points models.Coordinates, axis, sign int) (models.Coordinates, appError.AppError) {
	var path models.Coordinates
	switchAxisError.Wrap("setPointsAudStart")
	switchSignError.Wrap("setPointsAudStart")
	switchAxisError.Err = errors.New(fmt.Sprintf("no suitable value, expected: %d or %d received: %d", s.constData.axisX, s.constData.axisY, axis))
	switchSignError.Err = errors.New(fmt.Sprintf("no suitable value, expected: %d or %d received: %d", s.constData.positiveCoordinate, s.constData.negativeCoordinate, sign))
	switch axis {
	case s.constData.axisX:
		switch sign {
		case s.constData.positiveCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case s.constData.negativeCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = -points.Widht
			path.Height = points.Height

		default:
			return path, *switchSignError
		}
	case s.constData.axisY:
		switch sign {
		case s.constData.positiveCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case s.constData.negativeCoordinate:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = -points.Height

		default:
			return path, *switchSignError
		}
	default:
		return path, *switchAxisError
	}

	return path, appError.AppError{}
}

func (a *audToAudController) middleBuilding() (appError.AppError) {
	fmt.Println("Work")
	repository := NewRepository(a.client) // для обращение к базе данных
	// ось для перехода в другой сектор
	axis := axes.DefenitionAxis(a.endAudBorderPoint.Widht, a.endAudBorderPoint.Height, a.constData.axisX, a.constData.axisY)

	for i := 0; true; i++ {
		// if i == 3 {
		// 	break
		// } 
		// проверка вхождение координат пути в координаты границ сектора
		if a.checkOccurrence(a.points[i], axis, a.endAudBorderPoint) {
			axis = axes.ChangeAxis(axis, a.constData.axisX, a.constData.axisY)
			
			// расчет точек пути
			points, err := a.middleFinalPreparation(axis, a.endAudBorderPoint, a.points[i])
			if err.Err != nil {
				err.Wrap("building")
				return err
			}
			a.points = append(a.points, points)
			break
		} 
		// расчет точек пути
		points, err := a.middlePreparation(axis, a.endAudBorderPoint, a.points[i])
		if err.Err != nil {
			err.Wrap("building")
			return err
		}


		ok, err := repository.checkBorderAud2(points, a.sectorNumber)
		if err.Err != nil {
			err.Wrap("building")
			return err
		}

		// изменения оси построения, если точки входят в пределы аудитории
		if !ok {
			axis = axes.ChangeAxis(axis, a.constData.axisX, a.constData.axisY)
			points, err = a.middlePreparation(axis, a.endAudBorderPoint, a.points[i])
			if err.Err != nil {
				err.Wrap("building")
				return err
			}
			axis = axes.ChangeAxis(axis, a.constData.axisX, a.constData.axisY)
		}
		a.points = append(a.points, points)
	}

	return appError.AppError{}
}

// проверка на вхождение точек пути в пределы сектора.
func (a *audToAudController) checkOccurrence(points models.Coordinates, axis int, borderSector models.Coordinates) bool {
	switch axis {
	case a.constData.axisX:
		ph := points.X + points.Widht
		x1 := borderSector.X
		x2 := borderSector.X + borderSector.Widht
		if x1 <= ph && ph <= x2 {
			return true
		} else {
			return false
		}
	case a.constData.axisY:
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