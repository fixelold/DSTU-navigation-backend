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

	//TODO: получение координат секторов
	sectorBorderPoints, err := h.repository.getSectorBorderPoint(sectors[1])
	if err != nil {
		return nil, err
	}

	d := NewDrawPathAud2Sector(*auditory, *audBorderPoints, *sectorBorderPoints ,133, "1-333", h.repository)

	err = d.DrawInitPath()
	if err != nil {
		return nil, err
	}

	return d.Path, nil
}
