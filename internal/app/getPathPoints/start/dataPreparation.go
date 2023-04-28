package start

import "navigation/internal/models"

func preparation() {
	var coordinates models.Coordinates

	XX := borderPoint.X + 1
	YX := (borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2
	// YX := borderPoint.Y + 1

	XY := (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2
	// XY := borderPoint.X + 1
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
}