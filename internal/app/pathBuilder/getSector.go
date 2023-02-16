package pathBuilder

import (
	"fmt"
	"navigation/internal/appError"
	"strconv"
	"strings"
)

var (
	User000001 = appError.NewError("locationDetermination", "GetSelector", "Input does not match desired length", "-", "US-000001")
	User000002 = appError.NewError("locationDetermination", "GetSelector", "Not convert string to int", "-", "US-000002")
	User000003 = appError.NewError("locationDetermination", "GetSelector", "Errir in getting sector", "-", "US-000003")
)

func (h *handler) GetSector(start, end string) (int, int, error) {
	var err error

	startAud, startBuild, err := separationAudidotyNumber(start)
	if err != nil {
		return 0, 0, err
	}

	endAud, endBuild, err := separationAudidotyNumber(end)
	if err != nil {
		return 0, 0, err
	}

	fmt.Println(startAud)
	fmt.Println(endAud)

	sectorStart, err := h.repository.GetSector(startAud, uint(startBuild))
	if err != nil {
		h.logger.Errorf("the getSector function call returned %s", err.Error())
		User000003.ChangeDescription(err.Error())
		return 0, 0, User000003
	}

	sectorEnd, err := h.repository.GetSector(endAud, uint(endBuild))
	if err != nil {
		h.logger.Errorf("the getSector function call returned %s", err.Error())
		User000003.ChangeDescription(err.Error())
		return 0, 0, User000003
	}

	return sectorStart, sectorEnd, nil
}

func separationAudidotyNumber(number string) (string, int, error) {
	splitText := strings.Split(number, "-")
	if len(splitText) != 2 {
		return "", 0, User000001
	}

	building, err := strconv.Atoi(splitText[0])
	if err != nil {
		return "", 0, User000002
	}

	return number, building, nil
}
