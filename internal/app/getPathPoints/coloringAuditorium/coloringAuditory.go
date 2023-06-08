package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
)
var (
	TransitionNumber error = appError.NewAppError("transition number is empty")
)

type coloring struct {
	StartAuditoryNumber string             `json:"-"`
	EndAuditoryNumber   string             `json:"-"`
	StartAuditoryPoints models.Coordinates `json:"start"`
	EndAuditoryPoints   models.Coordinates `json:"end"`
	numberTransition    int

	logging    *logging.Logger `json:"-"`
	client postgresql.Client     `json:"-"`

	transition int
}

func NewColoring(start, end string, logging *logging.Logger, client postgresql.Client, transition, numberTransition int) *coloring {
	return &coloring{
		StartAuditoryNumber: start,
		EndAuditoryNumber:   end,
		logging:             logging,
		client:          client,
		transition:          transition,
		numberTransition: numberTransition,
	}
}

const (
	transitionNo = 1
	transitionYes = 2
	transitionToAud = 4
)

func (c *coloring) GetColoringPoints() appError.AppError {
	var err appError.AppError

	if c.transition == 3 {
		c.transition = transitionYes
	}

	switch c.transition {
	case transitionYes: // через переходный сектор
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

	case transitionNo: // переходного сектора нет
		c.StartAuditoryPoints, err = c.getColoringAudPoints(c.StartAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}
		c.EndAuditoryPoints, err = c.getColoringAudPoints(c.EndAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}

	case transitionToAud: // аудиторя находятся в одном секторе
		c.StartAuditoryPoints, err = c.getColoringTransitionPoints(c.EndAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}
		
		c.EndAuditoryPoints, err = c.getColoringAudPoints(c.EndAuditoryNumber)
		if err.Err != nil {
			err.Wrap("getAuditoryPoints")
			return err
		}
	}

	return err
}

// получение точек адуитория для ее окраски
func (c *coloring) getColoringAudPoints(number string) (models.Coordinates, appError.AppError) {
	var err appError.AppError
	repository := NewRepository(c.client, c.logging)
	coordinates, err := repository.getAudPoints(number)
	if err.Err != nil {
		err.Wrap("getColoringAudPoints")
		return models.Coordinates{}, err
	}

	return coordinates, err
}

// получение точек переходного сеткора для окраски
func (c *coloring) getColoringTransitionPoints(number string) (models.Coordinates, appError.AppError) {
	var err appError.AppError
	repository := NewRepository(c.client, c.logging)
	if c.numberTransition == 0 {
		err.Err = TransitionNumber
		err.Wrap("getColoringTransitionPoints")
		return models.Coordinates{}, err
	}

	coordinates, err := repository.getTransitionPoints(c.numberTransition)
	if err.Err != nil {
		err.Wrap("getColoringAudPoints")
		return models.Coordinates{}, err
	}

	return coordinates, err
}