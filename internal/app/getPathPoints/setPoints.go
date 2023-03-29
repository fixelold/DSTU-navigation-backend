package getPathPoints

import "navigation/internal/models"

// для начального пути от границ сектора
func (d *data) setPointsAudStart(points models.Coordinates, axis, sign int) (models.Coordinates, error) {
	var path models.Coordinates
	var err error
	switch axis {
	case AxisX:
		switch sign {
		case plus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case minus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = -points.Widht
			path.Height = points.Height

		default:
			err = nil
		}
	case AxisY:
		switch sign {
		case plus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = points.Height

		case minus:
			path.X = points.X
			path.Y = points.Y
			path.Widht = points.Widht
			path.Height = -points.Height

		default:
			err = nil
		}
	}
	return path, err
}

// точки от начала пути до вхождение в пределы сектора
func (d *data) setPointsPath2Sector(borderPoints, points, lastPathPoint models.Coordinates, axis int) (models.Coordinates, error) {
	var err error

	p := models.Coordinates{
		X: (points.X),
		Y: (points.Y)}
	p.Widht = points.Widht
	p.Height = points.Height
	return p, err
}