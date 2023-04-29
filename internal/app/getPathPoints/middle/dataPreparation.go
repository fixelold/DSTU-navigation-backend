package middle

import (
	"strconv"

	"navigation/internal/models"
)

func (m *middleController) preparation(axis int, borderPoint, points models.Coordinates) models.Coordinates {
	if axis == m.constData.axisX {
		if len(strconv.Itoa(m.sectorNumber)) == 4 { //stairs
			return models.Coordinates{
				X:      m.points[0].X,
				Y:      points.Y + points.Height,
				Widht:  borderPoint.X - points.X,
				Height: m.constData.heightX}
		} else {
			if borderPoint.X < points.X {
				return models.Coordinates{
					X:      points.X + points.Widht,
					Y:      points.Y + points.Height,
					Widht:  borderPoint.X - points.X,
					Height: m.constData.heightX}
			} else {
				return models.Coordinates{
					X:      points.X + points.Widht,
					Y:      points.Y + points.Height,
					Widht:  borderPoint.X - points.X,
					Height: m.constData.heightX}
			}
		}
	} else {
		if borderPoint.Y > points.Y {
			return models.Coordinates{
				X:      points.X + points.Widht,
				Y:      points.Y + points.Height - m.constData.heightX,
				Widht:  m.constData.widhtY,
				Height: borderPoint.Y - points.Y}
		} else {
			return models.Coordinates{
				X:      points.X + points.Widht,
				Y:      points.Y + points.Height,
				Widht:  m.constData.widhtY,
				Height: borderPoint.Y - (points.Y + points.Height)}
		}
	}
}