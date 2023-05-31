package audToAud

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

type middleController struct {
	Points []models.Coordinates
	constData constData
	sectorNumber int
	thisSectorNumber int
	typeTransition int
	endPoints models.Coordinates // начальный путь конечной аудитории

	client postgresql.Client
	logger *logging.Logger
}

func NewAudToAud(
	typeTransition int,
	thisSectorNumber int,
	sectorNumber int, 
	endPoints models.Coordinates,
	client postgresql.Client, 
	axisX, axisY, widhtX, heightX, widhtY, heightY int, 
	logger *logging.Logger) *middleController {
	return &middleController{
		typeTransition: typeTransition,
		thisSectorNumber: thisSectorNumber,
		sectorNumber: sectorNumber,
		endPoints: endPoints,
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

func (m *middleController) MiddlePoints(borderSector models.Coordinates) ([]models.Coordinates, appError.AppError) {
	err := m.building(borderSector)
	if err.Err != nil {
		err.Wrap("middlePoints")
		return nil, err
	}

	return m.Points, appError.AppError{}
}
