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
