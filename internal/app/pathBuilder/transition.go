package pathBuilder

import (
	"navigation/internal/appError"
)

func (h *handler) stairs(sector int) (int, appError.AppError) {
	transitionSector, err := h.repository.GetTransitionSector(sector, stairs)
	if err.Err != nil {
		err.Wrap("stairs")
		return 0, err
	}

	return transitionSector, err
}

func (h *handler) elevator(start, sector int) (int, appError.AppError) {
	sectorNumber :=  start % 10
	var transitionNumber int

	if sectorNumber == 1 || sectorNumber == 2 || sectorNumber == 3 {
		transitionNumber = 1000 + (start % 100 / 10 * 10) + 2
	
	} else if sectorNumber == 4 || sectorNumber == 5 || sectorNumber == 6 || sectorNumber == 7 {
		transitionNumber = 1000 + (start % 100 / 10 * 10) + 5
	}

	return transitionNumber, appError.AppError{}
}