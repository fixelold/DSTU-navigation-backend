package pathBuilder

import (
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

	splitText := strings.Split(start, "-")
	if len(splitText) != 2 {
		h.logger.Errorf("function GetSelector. Input does not match desired length expected: %d, received: %d", 2, len(splitText))
		return 0, 0, User000001
	}

	buildingStart, err := strconv.Atoi(splitText[0])
	if err != nil {
		h.logger.Errorf("function GetSelector not convert string to int, err: %s", err)
		User000002.ChangeDescription(err.Error())
		return 0, 0, User000002
	}

	startAuditory := splitText[1]

	splitText = strings.Split(end, "-")
	if len(splitText) != 2 {
		h.logger.Errorf("function GetSelector. Input does not match desired length expected: %d, received: %d", 2, len(splitText))
		return 0, 0, User000001
	}


	buildingEnd, err := strconv.Atoi(splitText[0])
	if err != nil {
		h.logger.Errorf("function GetSelector not convert string to int, err: %s", err)
		User000002.ChangeDescription(err.Error())
		return 0, 0, User000002
	}


	endAuditory := splitText[1]

	sectorStart, err := h.repository.GetSector(startAuditory, uint(buildingStart))
	if err != nil {
		h.logger.Errorf("the getSector function call returned %s", err.Error())
		User000003.ChangeDescription(err.Error())
		return 0,0, User000003
	}

	sectorEnd, err := h.repository.GetSector(endAuditory, uint(buildingEnd))
	if err != nil {
		h.logger.Errorf("the getSector function call returned %s", err.Error())
		User000003.ChangeDescription(err.Error())
		return 0, 0, User000003
	}

	return sectorStart, sectorEnd, nil
}
