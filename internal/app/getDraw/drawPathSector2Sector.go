package drawPath

import (
	"navigation/internal/models"
)

func (d *Path) DrawPathSector2Sector(borderSector models.Coordinates) error {
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

			p := d.prepare2(Sector2Sector, axis, borderSector, d.Path[i])

			points := d.getPoints2(p, d.Path[i], borderSector, axis)

			d.Path = append(d.Path, points)
			boolean = false
		} else {
			// TODO: возможно тут ошибка. Надо будет проверить и, если что, подправить.
			p := d.prepare2(Sector2Sector, axis, borderSector, d.Path[i])

			points := d.getPoints2(p, d.Path[i], borderSector, axis)
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
