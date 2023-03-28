package getPathPoints

import "navigation/internal/models"

func (d *data) preparePoints(pointsType, axis int, borderPoint models.Coordinates) models.Coordinates {
	switch pointsType {
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

	// case Path2Sector:
	// 	if axis == AxisX {
	// 		if borderPoint.X > path.X {
	// 			return models.Coordinates{
	// 				X:      path.X + path.Widht,
	// 				Y:      d.Path[0].Y + d.Path[0].Height,
	// 				Widht:  borderPoint.X - path.X,
	// 				Height: HeightX}
	// 		} else {
	// 			return models.Coordinates{
	// 				X:      path.X + path.Widht,
	// 				Y:      d.Path[0].Y + d.Path[0].Height,
	// 				Widht:  borderPoint.X - path.X,
	// 				Height: HeightX}
	// 		}
	// 	} else {
	// 		if borderPoint.Y > path.Y {
	// 			return models.Coordinates{
	// 				X:      path.X + path.Widht,
	// 				Y:      path.Y + path.Height - HeightX,
	// 				Widht:  WidhtY,
	// 				Height: borderPoint.Y - path.Y}
	// 		} else {
	// 			return models.Coordinates{
	// 				X:      path.X + path.Widht,
	// 				Y:      path.Y + path.Height,
	// 				Widht:  WidhtY,
	// 				Height: borderPoint.Y - (path.Y + path.Height)}
	// 		}
	// 	}

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