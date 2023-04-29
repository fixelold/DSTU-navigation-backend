package sectorToSector

import (
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
)

type constData struct {
	// positiveCoordinate int
	// negativeCoordinate int
	axisX int
	axisY int
	widhtX int
	heightX int
	widhtY int
	heightY int
}

type sectorToSectorController struct {
	points []models.Coordinates
	constData constData
	sectorNumber int

	client postgresql.Client
	logger *logging.Logger
}

func NewSectorToSectorController(
	sectorNumber int, 
	client postgresql.Client, 
	axisX, axisY, widhtX, heightX, widhtY, heightY int, 
	logger *logging.Logger) *sectorToSectorController {
	return &sectorToSectorController{
		sectorNumber: sectorNumber,
		client: client,
		logger: logger,
		constData: constData{
			axisX: axisX,
			axisY: axisY,
			widhtX: widhtX,
			heightX: heightX,
			widhtY: widhtY,
			heightY: heightY,
		},
	}
}

func (s *sectorToSectorController) sector2Sector(borderSector models.Coordinates) appError.AppError {
	iterator := (len(s.points) - 1)
	err := s.building(iterator, borderSector)
	if err.Err != nil {
		err.Wrap("sector2Sector")
		return err
	}

	return appError.AppError{}
}