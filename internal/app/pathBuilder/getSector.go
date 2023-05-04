package pathBuilder

import (
	"errors"
	"strconv"
	"strings"

	"navigation/internal/appError"
)

var (
	splitError = appError.NewAppError("can't split text")
)

func (h *handler) GetSector(start, end string) (int, int, appError.AppError) {
	var err appError.AppError
	
	startAud, startBuild, err := separationAudidotyNumber(start)
	if err.Err != nil {
		err.Wrap("file GetSector")
		return 0, 0, err
	}

	endAud, endBuild, err := separationAudidotyNumber(end)
	if err.Err != nil {
		err.Wrap("file GetSector")
		return 0, 0, err
	}

	sectorStart, err := h.repository.GetSector(startAud, uint(startBuild))
	if err.Err != nil {
		err.Wrap("file GetSector")
		return 0, 0, err
	}

	sectorEnd, err := h.repository.GetSector(endAud, uint(endBuild))
	if err.Err != nil {
		err.Wrap("file GetSector")
		return 0, 0, err
	}

	return sectorStart, sectorEnd, err
}

func separationAudidotyNumber(number string) (string, int, appError.AppError) {
	var err appError.AppError

	splitText := strings.Split(number, "-")
	if len(splitText) != 2 {
		splitError.Err = errors.New("wrong line lenght, exected: %s, received: %s")
		splitError.Wrap("separationAudidotyNumber")
		return "", 0, *splitError
	}

	building, error := strconv.Atoi(splitText[0])
	if err.Err != nil {
		err.Err = error
		err.Wrap("separationAudidotyNumber")
		return "", 0, err
	}

	return number, building, err
}
