package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

type coloringAuditory struct {
	StartAuditoryNumber string             `json:"-"`
	EndAuditoryNumber   string             `json:"-"`
	StartAuditoryPoints models.Coordinates `json:"start"`
	EndAuditoryPoints   models.Coordinates `json:"end"`

	logging    *logging.Logger `json:"-"`
	repository Repository     `json:"-"`
}

func NewColoringAudience(start, end string, logging *logging.Logger, repository Repository) *coloringAuditory {
	return &coloringAuditory{
		StartAuditoryNumber: start,
		EndAuditoryNumber: end,
		logging: logging,
		repository: repository,
	}
}

func (c *coloringAuditory) getAuditoryPoints() appError.AppError {
	var err appError.AppError
	err.Wrap("getAuditoryPoints")

	c.StartAuditoryPoints, err = c.repository.getAudPoints(c.StartAuditoryNumber)
	if err.Err != nil {
		return err
	}

	c.EndAuditoryPoints, err = c.repository.getAudPoints(c.EndAuditoryNumber)
	if err.Err != nil {
		return err
	}

	return err
}
