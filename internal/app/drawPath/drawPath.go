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

	res, err := DrawPathAuditory(borderPoints, auditory)
	if err != nil {
		return nil, err
	}

	points = append(points, res)

	return points, nil
}