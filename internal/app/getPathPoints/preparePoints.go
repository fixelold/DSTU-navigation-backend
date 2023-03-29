package getPathPoints

import "navigation/internal/models"

// подготовка данных
func (d *data) preparePoints(pointsType, axis int, borderPoint, points models.Coordinates) models.Coordinates {
	switch pointsType {
	// начальных путь от границ аудитории.
	case audStartPoints:
		var coordinates models.Coordinates

		XX := borderPoint.X + 1
		YX := (borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2

		XY := (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2
		YY := borderPoint.Y + 1

		if axis == AxisX {
			coordinates.X = XX
			coordinates.Y = YX
			coordinates.Widht = WidhtX
			coordinates.Height = HeightX

		} else if axis == AxisY {
			coordinates.X = XY
			coordinates.Y = YY
			coordinates.Widht = WidhtY
			coordinates.Height = HeightY
		}

		return coordinates
	// case Auditory2Sector:

	// 	if axis == AxisX {
	// 		if d.Path[0].Height == HeightY || d.Path[0].Widht == WidhtX{
	// 			return models.Coordinates{
	// 				X:      path.X + path.Widht,
	// 				Y:      d.Path[0].Y + d.Path[0].Height - HeightX,
	// 				Widht:  WidhtX,
	// 				Height: HeightX}
	// 		} else {
	// 			return models.Coordinates{
	// 				X:      path.X + path.Widht,
	// 				Y:      d.Path[0].Y + d.Path[0].Height,
	// 				Widht:  WidhtX,
	// 				Height: HeightX}
	// 		}
	// 	} else {
	// 		return models.Coordinates{
	// 			X:      path.X + path.Widht,
	// 			Y:      d.Path[0].Y + d.Path[0].Height,
	// 			Widht:  WidhtY,
	// 			Height: HeightY}
	// 	}
	// от начального пути до момента вхоэжение в пределы сектора
	case path2Sector:
		if axis == AxisX {
			return models.Coordinates{
				X:      points.X + points.Widht,
				Y:      d.points[0].Y + d.points[0].Height,
				Widht:  borderPoint.X - points.X,
				Height: HeightX}
		} else {
			if borderPoint.Y > points.Y {
				return models.Coordinates{
					X:      points.X + points.Widht,
					Y:      points.Y + points.Height - HeightX,
					Widht:  WidhtY,
					Height: borderPoint.Y - points.Y}
			} else {
				return models.Coordinates{
					X:      points.X + points.Widht,
					Y:      points.Y + points.Height,
					Widht:  WidhtY,
					Height: borderPoint.Y - (points.Y + points.Height)}
			}
		}

	// case Sector2Sector:
	// 	if axis == AxisX {
	// 		return models.Coordinates{
	// 			X:      path.X + path.Widht - WidhtY,
	// 			Y:      path.Y + path.Height,
	// 			Widht:  WidhtY,
	// 			Height: borderPoint.Y - (path.Y + path.Height)}
	// 	} else {
	// 		return models.Coordinates{
	// 			X:      path.X + path.Widht - WidhtY,
	// 			Y:      path.Y + path.Height,
	// 			Widht:  borderPoint.X - (path.X + path.Widht),
	// 			Height: HeightX}
	// 	}
	default:
		return models.Coordinates{}
	}
}