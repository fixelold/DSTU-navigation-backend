package drawPath

import (
	"fmt"
	"navigation/internal/models"
)

func (d *Path) DrawPathSector2Sector(borderSector models.Coordinates) error {
	var widht, height int
	iterator := 0
	axis := d.defenitionAxis(borderSector.Widht, borderSector.Height)
	boolean := true

	if axis == AxisX {
		axis = AxisY
	} else {
		axis = AxisX
	}

	for boolean {
		i := (len(d.Path) - 1) + iterator
		if d.checkPath2Sector(d.Path[i], axis) {
			fmt.Println("KEK - ", axis)
			switch axis {
			case AxisX:
				widht = WidhtY
				height = borderSector.Y - (d.Path[i].Y + d.Path[i].Height)
			case AxisY:
				widht = borderSector.X - (d.Path[i].X + d.Path[i].Widht)
				height = HeightX
			default:
				d.logger.Errorln("Function drawPathSector. Error switch default")
			}

			fmt.Println("border - ", borderSector)
			fmt.Println("path - ", d.Path[i])
			fmt.Println(widht, height)
			points := d.getPoints2Sector(d.Path[i].Y, d.Path[i].Height, widht, height, axis, d.Path[i], d.SectorBorderPoint)
			fmt.Println("new points - ", points)
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
			points := d.getPoints2Sector(d.Path[i].Y, d.Path[i].Height, widht, height, axis, d.Path[iterator], d.SectorBorderPoint)
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
