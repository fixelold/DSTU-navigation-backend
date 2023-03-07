package drawPath

import (
	"fmt"
	"navigation/internal/logging"
	"navigation/internal/models"
)

func (h *handler) drawPath(start, end string, sectors []int) ([]models.Coordinates, error) {

	entry, exit := min(sectors[0], sectors[1])
	auditory, audBorderPoints, sectorBorderPoints, err := h.getData(start, entry, exit)
	if err != nil {
		return nil, err
	}

	h.logger.Infoln("Draw Init Path")
	d := NewPath(*auditory, *audBorderPoints, *sectorBorderPoints, sectors[1], start, h.repository, logging.GetLogger())

	err = d.DrawInitPath()
	if err != nil {
		return nil, err
	}

	h.logger.Infoln("Draw Sector to sector")
	for i := 1; i < len(sectors)-1; i++ {
		entry, exit := min(sectors[i], sectors[i+1])
		borderSector, err := h.repository.getSectorBorderPoint2(entry, exit)
		fmt.Println("ONE - ", sectorBorderPoints)
		fmt.Println("TWO - ", borderSector)
		if err != nil {
			return nil, err
		}

		err = d.DrawPathSector2Sector(*borderSector)
		if err != nil {
			return nil, err
		}
	}

	entry, exit = min(sectors[len(sectors)-1], sectors[len(sectors)-2])
	auditory, audBorderPoints, sectorBorderPoints, err = h.getData(end, entry, exit)
	if err != nil {
		return nil, err
	}

	do := NewPath(*auditory, *audBorderPoints, *sectorBorderPoints, sectors[1], end, h.repository, logging.GetLogger())

	h.logger.Infoln("Draw Final Path")
	err = do.DrawInitPath()
	if err != nil {
		return nil, err
	}

	d.Path = append(d.Path, do.Path...)
	return d.Path, nil
}

// TODO сделать проверку на равно. На всякий
func min(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// получение информацию о аудиториях и секторах
func (h *handler) getData(aud string, entry, exit int) (*models.Coordinates, *models.Coordinates, *models.Coordinates, error) {
	audBorderPoints, err := h.repository.getAudBorderPoint(aud)
	if err != nil {
		return &models.Coordinates{}, &models.Coordinates{}, &models.Coordinates{}, err
	}

	auditory, err := h.repository.getAuditoryPosition(aud)
	if err != nil {
		return &models.Coordinates{}, &models.Coordinates{}, &models.Coordinates{}, err
	}

	sectorBorderPoints, err := h.repository.getSectorBorderPoint2(entry, exit)
	if err != nil {
		return &models.Coordinates{}, &models.Coordinates{}, &models.Coordinates{}, err
	}

	return auditory, audBorderPoints, sectorBorderPoints, nil
}
