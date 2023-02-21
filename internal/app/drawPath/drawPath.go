package drawPath

func (h *handler) drawPath(start, end string, sectors []int) ([][]int, error) {
	var points [][]int
	borderPoints, err := h.repository.getBorderPoint(start)
	if err != nil {
		return nil, err
	}

	auditory, err := h.repository.getAuditoryPosition(start)
	if err != nil {
		return nil, err
	}

	//TODO: получение координат секторов

	d := NewDrawPathAud2Sector(*auditory, *borderPoints ,133, "1-333", h.repository, )

	err = d.DrawInitPath()
	if err != nil {
		return nil, err
	}

	points = append(points, d.Path)

	return points, nil
}
