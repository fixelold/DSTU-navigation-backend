package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (d *data) sector2Sector(borderSector models.Coordinates) appError.AppError {
	iterator := (len(d.points) - 1)
	err := d.otherPathPoints(iterator, borderSector, sector2Sector)
	if err.Err != nil {
		err.Wrap("sector2Sector")
		return err
	}

	return appError.AppError{}
}