package drawPath

import "navigation/internal/models"

func checkBorder(path, auditory *models.Reactangle) bool {
	pointX := path.X + path.Widht
	pointY := path.Y + path.Height

	if auditory.X <= pointX && pointX <= auditory.Widht {
		return false
	} else if auditory.Y <= pointY && pointY <= auditory.Height {
		return false
	}

	return true
}
