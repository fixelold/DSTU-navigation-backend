package getPathPoints

import "navigation/internal/models"

func (d *data) middlePoints(borderSector models.Coordinates) error {
	if err := d.otherPathPoints(0, borderSector, path2Sector); err != nil {
		return err
	}

	return nil
}
