package drawPath

import drawPath "navigation/internal/app/drawPath/drawPathAud2Sector"

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

	d := drawPath.NewDrawPathAud2Sector(*auditory, *borderPoints ,133, "1-33")

	err = d.DrawInitPath()
	if err != nil {
		return nil, err
	}

	points = append(points, d.Path)

	return points, nil
}
