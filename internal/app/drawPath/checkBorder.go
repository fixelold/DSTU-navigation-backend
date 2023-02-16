package drawPath

import (
	"fmt"
	"navigation/internal/models"
)

func checkBorder(axis int, path, auditory *models.Reactangle) bool {
	pointX := path.X + path.Widht
	pointY := path.Y + path.Height

	switch axis {
	case AxisX:
		if auditory.X <= pointX && pointX <= auditory.Widht+auditory.X {
			fmt.Println(auditory)
			fmt.Println(path)
			fmt.Println("Error")
			return false
		}
	case AxisY:
		if auditory.Y <= pointY && pointY <= auditory.Height+auditory.Y {
			return false
		}
	}
	return true
}
