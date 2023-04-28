package getPathPoints

import (
	"fmt"

	"navigation/internal/appError"
	"navigation/internal/models"
)

func (d *data) otherPathPoints(iterator int, borderSector models.Coordinates, pointsType int) appError.AppError {
	boolean := true
	axis := d.defenitionAxis(borderSector.Widht, borderSector.Height)
	// potomYbrat := 0 // sector2sector
	// ybrat := 0 // sector2sector
	// potomYbrat2 := 0
	for boolean {
		// if ybrat == 1 {
		// 	axis = d.changeAxis(axis)
		// }
		// d.logger.Logger.Infoln("all points: ", d.points)
		if d.checkOccurrence(d.points[iterator], axis, borderSector) {
			if pointsType == sector2Sector {
				fmt.Printf("sector to sector (check = true), axis: %d\n iteration: %d\n", axis, iterator)
			}
			d.pathAlignment(borderSector, axis)

			// if ybrat == 1 {
			// 	axis = d.changeAxis(axis)
			// }

			// if pointsType == sector2Sector && potomYbrat > 0 {
			// 	axis = d.changeAxis(axis)
			// }

			// if pointsType != sector2Sector {
			// 	axis = d.changeAxis(axis)
			// }

			axis = d.changeAxis(axis)

			points := d.preparePoints(pointsType, axis, borderSector, d.points[iterator])

			points = d.setPointsPath2Sector(borderSector, points, d.points[iterator], axis)

			d.points = append(d.points, points)
			fmt.Println("all points: ", d.points)
			boolean = false
		} else {
			if pointsType == path2Sector {
				fmt.Printf("path to sector, points: %d\n", d.points[iterator])
			}

			// if ybrat == 1 {
			// 	axis = d.changeAxis(axis)
			// }
			// if pointsType == path2Sector {
			// 	potomYbrat2 += 1
			// }
			// if pointsType == sector2Sector {
			// 	if ybrat == 0 {
			// 		axis = d.changeAxis(axis)
			// 		ybrat += 1
			// 	}
			// 	potomYbrat += 1
			// }
			if pointsType == sector2Sector {
				fmt.Printf("sector to sector, axis: %d\n iteration: %d\n", axis, iterator)
				boolean = false
			}
			points := d.preparePoints(pointsType, axis, borderSector, d.points[iterator])

			points = d.setPointsPath2Sector(borderSector, points, d.points[iterator], axis)

			ok, err := d.repository.checkBorderAud(points)
			if err.Err != nil {
				err.Wrap("otherPathPoints")
				return err
			}

			ok2, err := d.repository.checkBorderSector(points)
			if err.Err != nil {
				err.Wrap("otherPathPoints")
				return err
			}

			if !ok && !ok2 {
				//TODO написать изменения направления или типо что-то такого
			}

			d.points = append(d.points, points)
		}

		iterator += 1
	}

	return appError.AppError{}
}

// проверка на вхождение точек пути в пределы сектора.
func (d *data) checkOccurrence(points models.Coordinates, axis int, borderSector models.Coordinates) bool {
	switch axis {
	case AxisX:
		ph := points.X + points.Widht
		x1 := borderSector.X
		x2 := borderSector.X + borderSector.Widht
		if x1 <= ph && ph <= x2 {
			return true
		} else {
			return false
		}
	case AxisY:
		ph := points.Y + points.Height
		y1 := borderSector.Y
		y2 := borderSector.Y + borderSector.Height
		if y1 <= ph && ph <= y2 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

// выравнивание пути
func (d *data) pathAlignment(sectorBorderPoint models.Coordinates, axis int) {
	lenght := len(d.points)
	path := d.points[lenght-1]
	switch axis {
	case AxisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			points.Widht = sectorPoints - path.X
			points.Height = HeightX
			d.points[lenght-1].Widht = points.Widht
		} else if sectorPoints < path.X {
			points.Widht = sectorPoints - path.X
			points.Height = HeightX
			d.points[lenght-1].Widht = points.Widht
		}
	case AxisY:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		if sectorPoints > path.Y {
			points.Widht = WidhtY
			points.Height = sectorPoints - path.Y
			d.points[lenght-1].Height = points.Height
		} else if sectorPoints < path.Y {
			points.Widht = WidhtY
			points.Height = sectorPoints - path.Y
			d.points[lenght-1].Height = points.Height
		}
	default:
		d.logger.Errorln("Path Alignment default")
	}
}
