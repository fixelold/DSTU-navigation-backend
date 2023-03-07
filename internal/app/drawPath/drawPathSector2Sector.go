package drawPath

import (
	"navigation/internal/models"
)

func (d *Path) DrawPathSector2Sector(borderSector models.Coordinates) error {
	var widht, height int
	iterator := 0
	axis := d.defenitionAxis(d.SectorBorderPoint.Widht, d.SectorBorderPoint.Height)
	boolean := true

	for boolean {
		i := (len(d.Path) - 1) + iterator
		if d.checkPath2Sector(d.Path[i], axis) {
			switch axis {
			case AxisX:
				widht = d.SectorBorderPoint.X - (d.Path[iterator].X + d.Path[iterator].Widht)
				height = HeightX
			case AxisY:
				widht = WidhtY
				height =  d.SectorBorderPoint.Y - (d.Path[iterator].Y + d.Path[iterator].Y)
			default:
				d.logger.Errorln("Function drawPathSector. Error switch default")
			}
			points := d.getPoints2Sector(widht, height, axis, d.Path[iterator], d.SectorBorderPoint)

			d.Path = append(d.Path, points)
			boolean = false
		} else {
				// TODO: возможно стоит вынести в отдельную функцию
				switch axis {
				case AxisX:
					widht = WidhtY
					height = HeightY
				case AxisY:
					widht = WidhtX
					height = HeightX
				default:
					d.logger.Errorln("Function DrawPathSector2Sector. Error switch default")
				}
				points := d.getPoints2Sector(widht, height, axis, d.Path[iterator], d.SectorBorderPoint)
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
