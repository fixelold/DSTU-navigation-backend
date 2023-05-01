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
	Points []models.Coordinates
	LastSector bool
	constData constData
	sectorNumber int
	OldAxis int

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

func (s *sectorToSectorController) Sector2SectorPoints(borderSector models.Coordinates, lenPoint int) ([]models.Coordinates,appError.AppError) {
	iterator := lenPoint
	err := s.building(iterator, borderSector)
	if err.Err != nil {
		err.Wrap("sector2Sector")
		return nil, err
	}

	return s.Points, appError.AppError{}
}