package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

type data struct {
	audPoints          models.Coordinates // координаты аудитории.
	audBorderPoints    models.Coordinates // координаты места отрисовки пути (одна из границ аудитории).
	sectorBorderPoints models.Coordinates // координаты одной из границ сектора.

	nextSectorNumber int    // номер сектора.
	audNumber        string // номер аудитории.

	logger     *logging.Logger // логирования.
	repository Repository      // для обращения к базе данных.

	points   []models.Coordinates // массив координат. Для построения пути.
}

func newData(audNumber string, sectorEntry, sectorExit, nextSectorNumber int, logger *logging.Logger, repository Repository) (*data, appError.AppError) {
	var err appError.AppError
	data := &data{
		audNumber:    audNumber,
		nextSectorNumber: nextSectorNumber,
		logger:       logger,
		repository:   repository,
	}

	err = data.getPoints(sectorEntry, sectorExit) 
	if err.Err != nil {
		err.Wrap("newData")
		return nil, err
	}

	return data, err
}
 
// получение audPoints, audBorderPoints, sectorBorderPoints
func (d *data) getPoints(entry, exit int) appError.AppError {
	var err appError.AppError
	// получаем координаты аудитории по ее номеру.
	d.audPoints, err = d.repository.getAudPoints(d.audNumber)
	if err.Err != nil {
		err.Wrap("getPoints")	
		return err
	}

	// получаем координаты границ аудитории по ее номеру.
	d.audBorderPoints, err = d.repository.getAudBorderPoint(d.audNumber)
	if err.Err != nil {
		err.Wrap("getPoints")	
		return err
	}

	// получаем координаты одной из границ сектора. По значению входа и выхода из него.
	d.sectorBorderPoints, err = d.repository.getSectorBorderPoint(entry, exit)
	if err.Err != nil {
		err.Wrap("getPoints")	
		return err
	}

	return err
}
