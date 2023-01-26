package locationDetermination

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"strconv"
	"strings"
)

var (
	User000001 = appError.NewError("locationDetermination", "GetSelector", "Input does not match desired length", "-", "US-000001")
	User000002 = appError.NewError("locationDetermination", "GetSelector", "Not convert string to int", "-", "US-000002")
	User000003 = appError.NewError("locationDetermination", "GetSelector", "Errir in getting sector", "-", "US-000003")
)

type location struct {
	repository Repository
	logger     *logging.Logger
}

func NewLocation(logger *logging.Logger, repository Repository) *location {
	return &location{
		logger:     logger,
		repository: repository,
	}
}

func (l *location) GetSector(number string) (uint, error) {
	var err error

	splitText := strings.Split(number, "-")
	if len(splitText) != 2 {
		l.logger.Errorf("function GetSelector. Input does not match desired length expected: %d, received: %d", 2, len(splitText))
		return 0, User000001
	}

	audienceNumber := splitText[1]

	building, err := strconv.Atoi(splitText[0])
	if err != nil {
		l.logger.Errorf("function GetSelector not convert string to int, err: %s", err)
		User000002.ChangeDescription(err.Error())
		return 0, User000002
	}

	sector, err := l.repository.GetSector(audienceNumber, uint(building))
	if err != nil {
		l.logger.Errorf("the getSector function call returned %s", err.Error())
		User000003.ChangeDescription(err.Error())
		return 0, User000003
	}

	return sector, nil
}
