package getPathPoints

import (
	"navigation/internal/models"
)

func (d *data) middlePoints() error {
	iterator := 0
	boolean := true
	axis := d.defenitionAxis(d.sectorBorderPoints.Widht, d.sectorBorderPoints.Height)

	for boolean {
		if d.checkOccurrence(d.points[iterator], axis) {

			d.pathAlignment(d.sectorBorderPoints, axis)

			axis = d.changeAxis(axis)

			points := d.preparePoints(path2Sector, axis, d.sectorBorderPoints, d.points[iterator])

			points, err := d.setPointsPath2Sector(d.sectorBorderPoints, points, d.Path[iterator], axis)
			if err != nil {
				return err
			}

			d.points = append(d.points, points)
			boolean = false
		} else {

			p := d.prepare2(Auditory2Sector, axis, d.SectorBorderPoint, d.Path[iterator])

			points := d.getPoints2(p, d.Path[iterator], d.SectorBorderPoint, axis)
			if points == (models.Coordinates{}) {
				return User000004
			}

			ok, err := d.Repository.checkBorderAud(points)
			if err != nil {
				return User000004
			}

			ok2, err := d.Repository.checkBorderSector(points)
			if err != nil {
				return User000004
			}

			if !ok && !ok2 {
				//TODO написать изменения направления или типо что-то такого
			}

			d.Path = append(d.Path, points)
		}

		iterator += 1
	}

	return nil
}

// проверка на вхождение точек пути в пределы сектора.
func (d *data) checkOccurrence(points models.Coordinates, axis int) bool {
	switch axis {
	case AxisX:
		ph := points.X + points.Widht
		x1 := d.sectorBorderPoints.X
		x2 := d.sectorBorderPoints.X + d.sectorBorderPoints.Widht
		if x1 <= ph && ph <= x2 {
			return true
		} else {
			return false
		}
	case AxisY:
		ph := points.Y + points.Height
		y1 := d.sectorBorderPoints.Y
		y2 := d.sectorBorderPoints.Y + d.sectorBorderPoints.Height
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