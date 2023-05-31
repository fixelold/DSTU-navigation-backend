package sectorToSector

import (
	"navigation/internal/models"
)

func (m *sectorToSectorController) leftDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht, // added for 1-353 to 1-116
            Y: points.Y,
            Widht: 5,
            Height: borderPoint.Y - points.Y,
		}
	} else if borderPoint.Y > points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y, // убрал
            Widht: -5, // добавил -
            Height: borderPoint.Y - points.Y + points.Height,
		}
	}

	return result
}

func (m *sectorToSectorController) leftUpY(borderPoint, points models.Coordinates) models.Coordinates {
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
            Y: points.Y + points.Height, // добавил points.Height
            Widht: -5, // добавил -
            Height: borderPoint.Y - points.Y + 5, // добавил 5
		}
	}


	return result
}

func (m *sectorToSectorController) rightDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht, // добавил points.Widht
            Y: points.Y + points.Height,
            Widht: 5,
            Height: borderPoint.Y - (points.Y + points.Height),
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

func (m *sectorToSectorController) rightUpY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X + points.Widht,
            Y: points.Y, // убрал points.Height
            Widht: 5, // убрал -
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