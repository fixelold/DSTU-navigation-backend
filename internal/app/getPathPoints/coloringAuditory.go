package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

type coloring struct {
	StartAuditoryNumber string             `json:"-"`
	EndAuditoryNumber   string             `json:"-"`
	StartAuditoryPoints models.Coordinates `json:"start"`
	EndAuditoryPoints   models.Coordinates `json:"end"`

	logging    *logging.Logger `json:"-"`
	repository Repository     `json:"-"`

	transition int
}

func NewColoring(start, end string, logging *logging.Logger, repository Repository, transition int) *coloring {
	return &coloring{
		StartAuditoryNumber: start,
		EndAuditoryNumber: end,
		logging: logging,
		repository: repository,
		transition: transition,
	}
}

func (c *coloring) GetColoringPoints() appError.AppError {
	var err appError.AppError

	switch c.transition {
	case transitionYes:
		c.StartAuditoryPoints, err = c.getColoringAudPoints(c.StartAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}

		c.EndAuditoryPoints, err = c.getColoringTransitionPoints(c.EndAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}
	
	case transitionNo:
		c.StartAuditoryPoints, err = c.getColoringAudPoints(c.StartAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}
		c.EndAuditoryPoints, err = c.getColoringAudPoints(c.StartAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}
	}

	return err
}

func (c *coloring) getColoringAudPoints(number string) (models.Coordinates, appError.AppError) {
	var err appError.AppError

	coordinates, err := c.repository.getAudPoints(number)
	if err.Err != nil {
		err.Wrap("getColoringAudPoints")
		return models.Coordinates{}, err
	}

	return coordinates, err
}

func (c *coloring) getColoringTransitionPoints(number string) (models.Coordinates, appError.AppError) {
	var err appError.AppError

	coordinates, err := c.repository.getAudPoints(number)
	if err.Err != nil {
		err.Wrap("getColoringAudPoints")
		return models.Coordinates{}, err
	}

	return coordinates, err
}