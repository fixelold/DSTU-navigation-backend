package getPathPoints

import "navigation/internal/models"

func (d *data) setPoints(pointsType, axis, sign int, points models.Coordinates) (models.Coordinates, error) {
	var path models.Coordinates
	var err error

	switch pointsType {
	case audStartPoints:
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

	default:
		return models.Coordinates{}, nil
	}
}
