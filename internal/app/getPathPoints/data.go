package getPathPoints

import (
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

func newData(audNumber string, sectorEntry, sectorExit, nextSectorNumber int, logger *logging.Logger, repository Repository) (*data, error) {
	data := &data{
		audNumber:    audNumber,
		nextSectorNumber: nextSectorNumber,
		logger:       logger,
		repository:   repository,
	}

	if err := data.getPoints(sectorEntry, sectorExit); err != nil {
		return nil, err
	}

	return data, nil
}

// получение audPoints, audBorderPoints, sectorBorderPoints
func (d *data) getPoints(entry, exit int) error {
	var err error

	// получаем координаты аудитории по ее номеру.
	d.audPoints, err = d.repository.getAudPoints(d.audNumber)
	if err != nil {
		return err
	}

	// получаем координаты границ аудитории по ее номеру.
	d.audBorderPoints, err = d.repository.getAudBorderPoint(d.audNumber)
	if err != nil {
		return err
	}

	// получаем координаты одной из границ сектора. По значению входа и выхода из него.
	d.sectorBorderPoints, err = d.repository.getSectorBorderPoint(entry, exit)
	if err != nil {
		return err
	}

	return nil
}
