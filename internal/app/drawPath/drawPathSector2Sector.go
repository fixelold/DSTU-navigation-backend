package drawPath

import "navigation/internal/models"

func (d *Path) DrawPathSector2Sector(borderSector models.Coordinates) error {
	iterator := 0
	axis := d.defenitionAxis(d.SectorBorderPoint.Widht, d.SectorBorderPoint.Height)
	boolean := true

	for boolean {
		if d.checkPath2Sector(d.Path[iterator], axis) {
			points := d.getDrawPoints2Sector(d.Path[iterator], axis)

			d.Path = append(d.Path, points)
			boolean = false
		} else {
			// определяем в каком направлении рисовать
			points := d.getDrawPoints(d.Path[iterator], axis)
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