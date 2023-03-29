package getPathPoints

import "navigation/internal/models"

func (d *data) sector2Sector(borderSector models.Coordinates) error {
	iterator := (len(d.points) - 1)
	if err := d.otherPathPoints(iterator, borderSector); err != nil {
		return err
	}

	return nil
}