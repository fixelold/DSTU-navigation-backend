package drawPath

import "navigation/internal/models"

func (h *handler) drawPath(start, end string, sectors []int) ([]models.Coordinates, error) {
	audBorderPoints, err := h.repository.getAudBorderPoint(start)
	if err != nil {
		return nil, err
	}

	auditory, err := h.repository.getAuditoryPosition(start)
	if err != nil {
		return nil, err
	}

	sectorBorderPoints, err := h.repository.getSectorBorderPoint(sectors[1])
	if err != nil {
		return nil, err
	}

	d := NewPath(*auditory, *audBorderPoints, *sectorBorderPoints ,133, "1-333", h.repository)

	err = d.DrawInitPath()
	if err != nil {
		return nil, err
	}

	//TODO тут будет цикл. Вообще все в цикл обернуть надо бы.

	entry, exit := min(sectors[1], sectors[2])
	borderSector, err := h.repository.getSectorBorderPoint2(entry, exit)
	if err != nil {
		return nil, err
	}

	err = d.DrawPathSector2Sector(*borderSector)
	if err != nil {
		return nil, err
	}

	return d.Path, nil
}

// TODO сделать проверку на равно. На всякий
func min(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}