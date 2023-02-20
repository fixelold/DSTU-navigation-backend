package drawPath

import (
	"navigation/internal/models"
)

func checkBorderAud(axis int, path, auditory models.Coordinates) bool {
	pointX := path.X + path.Widht
	pointY := path.Y + path.Height

	switch axis {
	case AxisX:
		if auditory.X <= pointX && pointX <= auditory.Widht+auditory.X {
			return false
		}
	case AxisY:
		if auditory.Y <= pointY && pointY <= auditory.Height+auditory.Y {
			return false
		}
	}
	return true
}
