package getPathPoints

import (
	"strconv"

	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
)

var (
	stairs = 4
)

var (
	TransitionError = appError.NewAppError("not info for transition")
)

type data struct {
	audPoints          models.Coordinates // координаты аудитории.
	audBorderPoints    models.Coordinates // координаты места отрисовки пути (одна из границ аудитории).
	sectorBorderPoints models.Coordinates // координаты одной из границ сектора.

	sectorNumber     int    // номер сектора.
	nextSectorNumber int    // номер следующего сектора.
	audNumber        string // номер аудитории.

	logger     *logging.Logger // логирования.
	client postgresql.Client      // для обращения к базе данных.

	points []models.Coordinates // массив координат. Для построения пути.

	transition       int
	transitionNumber int

	sectorType int
}

func newData(audNumber string, 
	sectorEntry, sectorExit, nextSectorNumber int, 
	logger *logging.Logger, 
	client postgresql.Client,
	transition, transitionNumber int) (*data, appError.AppError) {
	var err appError.AppError
	data := &data{
		audNumber:        audNumber,
		sectorNumber: sectorExit, //TODO: тут может быть ошибка. Может пердаваться не верный сектор. 
		nextSectorNumber: nextSectorNumber,
		logger:           logger,
		client:       client,
		transition: transition,
		transitionNumber: transitionNumber,
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
	repository := NewRepository(d.client, d.logger)
	// получаем координаты аудитории по ее номеру.
	if d.transition == transitionYes {
		d.audPoints, err = repository.getTransitionPoints(d.transitionNumber)

		if err.Err != nil {
			err.Wrap("getPoints")
			return err
		}
	} else if d.transition == transitionNo {
		d.audPoints, err = repository.getAudPoints(d.audNumber)
		if err.Err != nil {
			err.Wrap("getPoints")
			return err
		}
	} else {
		err.Err = TransitionError
		err.Wrap("getPoints")
	}

	// получаем координаты границ аудитории по ее номеру.
	d.audBorderPoints, err = repository.getAudBorderPoint(d.audNumber)
	if err.Err != nil {
		err.Wrap("getPoints")
		return err
	}

	// получаем координаты одной из границ сектора. По значению входа и выхода из него.
	if len(strconv.Itoa(exit)) == stairs {
		d.sectorBorderPoints, err = repository.getTransitionSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPoints")
			return err
		}
	} else {
		d.sectorBorderPoints, err = repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPoints")
			return err
		}
	}

	return err
}
