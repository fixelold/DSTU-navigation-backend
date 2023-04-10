package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
	"strconv"
)

var (
	stairs = 4
)

type data struct {
	audPoints          models.Coordinates // координаты аудитории.
	audBorderPoints    models.Coordinates // координаты места отрисовки пути (одна из границ аудитории).
	sectorBorderPoints models.Coordinates // координаты одной из границ сектора.

	sectorNumber     int    // номер сектора.
	nextSectorNumber int    // номер следующего сектора.
	audNumber        string // номер аудитории.

	logger     *logging.Logger // логирования.
	repository Repository      // для обращения к базе данных.

	points []models.Coordinates // массив координат. Для построения пути.
}

func newData(audNumber string, sectorEntry, sectorExit, nextSectorNumber int, logger *logging.Logger, repository Repository) (*data, appError.AppError) {
	var err appError.AppError
	data := &data{
		audNumber:        audNumber,
		sectorNumber: sectorExit, //TODO: тут может быть ошибка. Может пердаваться не верный сектор. 
		nextSectorNumber: nextSectorNumber,
		logger:           logger,
		repository:       repository,
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
	if len(strconv.Itoa(exit)) == stairs {
		d.sectorBorderPoints, err = d.repository.getTransitionSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPoints")
			return err
		}
	} else {
		d.sectorBorderPoints, err = d.repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPoints")
			return err
		}
	}

	return err
}
