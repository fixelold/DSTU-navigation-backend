package audToAud

import "navigation/internal/models"

func (s *audToAudController) preparation(axis int, borderPoint models.Coordinates) models.Coordinates {
	var coordinates models.Coordinates

	XX := borderPoint.X
	YX := (borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2

	XY := (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2
	YY := borderPoint.Y

	if axis == s.constData.axisX {
		coordinates.X = XX
		coordinates.Y = YX
		coordinates.Widht = s.constData.widhtX
		coordinates.Height = s.constData.heightX

	} else if axis == s.constData.axisY {
		coordinates.X = XY
		coordinates.Y = YY
		coordinates.Widht = s.constData.widhtY
		coordinates.Height = s.constData.heightY
	}

	return coordinates
}