package locationDetermination

import (
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/logging"
	"strconv"
	"strings"
)

type location struct {
	audienceNumber string
	logger         *logging.Logger
}

func NewLocation(audienceNumber string, logger *logging.Logger) *location {
	return &location{
		audienceNumber: audienceNumber,
		logger:         logger,
	}
}

func (l *location) GetSector() (sector uint, error error) {
	splitText := strings.Split(l.audienceNumber, "-")
	if len(splitText) != 2 {
		l.logger.Errorf("function GetSelector. Input does not match desired length expected: %d, received: %d", 2, len(splitText))
		return 0, appError.NewError(
			"locationDetermination",
			"GetSelector",
			"Input does not match desired length",
			"-",
			"US-000001",
		)
	}

	building := splitText[0]
	audienceNumber, err := strconv.Atoi(splitText[1])
	if err != nil {
		l.logger.Errorf("function GetSelector not convert string to int, err: %s", err)
		return 0, appError.NewError(
			"locationDetermination",
			"GetSelector",
			"Not convert string to int",
			err.Error(),
			"US-000002",
		)
	}

	if 225 > audienceNumber && audienceNumber > 213 {
		fmt.Println(building, audienceNumber, "1")
	} else if 238 > audienceNumber && audienceNumber > 230 {
		fmt.Println(building, audienceNumber, "2")
	}

	return 0, nil
}
