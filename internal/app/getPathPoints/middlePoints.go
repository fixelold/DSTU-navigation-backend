package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/models"
)

func (d *data) middlePoints(borderSector models.Coordinates) appError.AppError {
	err := d.otherPathPoints(0, borderSector, path2Sector)
	if err.Err != nil {
		err.Wrap("middlePoints")
		return err
	}

	return appError.AppError{}
}
