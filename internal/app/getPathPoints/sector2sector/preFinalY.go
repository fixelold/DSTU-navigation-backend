package sectorToSector

import "navigation/internal/models"

func (m *sectorToSectorController) leftDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X,
            Y: points.Y,
            Widht: 5,
            Height: borderPoint.Y - points.Y,
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

func (m *sectorToSectorController) leftUpY(borderPoint, points models.Coordinates) models.Coordinates {
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

func (m *sectorToSectorController) rightDownY(borderPoint, points models.Coordinates) models.Coordinates {
	var result models.Coordinates
	if borderPoint.Y < points.Y {
		result = models.Coordinates{
			X: points.X,
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

func (m *sectorToSectorController) rightUpY(borderPoint, points models.Coordinates) models.Coordinates {
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