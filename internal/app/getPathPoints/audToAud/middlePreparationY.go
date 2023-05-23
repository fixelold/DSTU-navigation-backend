package audToAud

import (
	"navigation/internal/models"
)

func (m *middleController) leftDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht, // added points.Widht
            Y: points.Y + points.Height, // added points.Height
            Widht: 5,
            Height: borderPoint.Y - (points.Y + points.Height), // added points.Height
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: 5,
            Height: borderPoint.Y - (points.Y + points.Height),
		}
	}

	return result
}

func (m *middleController) leftUpY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y, // убрал points.Height
            Widht: 5,
            Height: borderPoint.Y - points.Y,
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y,
            Widht: 5,
            Height: borderPoint.Y - points.Y,
		}
	}


	return result
}

func (m *middleController) rightDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y,
            Widht: -5,
            Height: borderPoint.Y - points.Y,
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y + points.Height,
            Widht: -5,
            Height: borderPoint.Y - (points.Y + points.Height),
		}
	}
	return result

}

func (m *middleController) rightUpY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y, // убрал points.Height
            Widht: -5,
            Height: borderPoint.Y - points.Y,
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y,
            Widht: -5,
            Height: borderPoint.Y - points.Y,
		}
	}

	return result
}