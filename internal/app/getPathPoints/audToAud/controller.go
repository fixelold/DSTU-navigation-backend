package audToAud

import (
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/models"
)

type audToAudController struct {
	points []models.Coordinates
	startAudBorderPoint models.Coordinates
	endAudBorderPoint models.Coordinates
	startAud string
	endAud string
	sectorNumber int
	client postgresql.Client
	constData constData
}

type constData struct {
	positiveCoordinate int
	negativeCoordinate int
	axisX int
	axisY int
	widhtX int
	heightX int
	widhtY int
	heightY int
}


func NewAudToAudController(
	startAudBorderPoint, endAudBorderPoint models.Coordinates,
	startAud, endAud string,
	sectorNumber int,
	client postgresql.Client,
	positiveCoordinate, negativeCoordinate int,
	axisX, axisY, widhtX, heightX, widhtY, heightY int) *audToAudController {
		return &audToAudController{
			startAudBorderPoint: startAudBorderPoint,
			endAudBorderPoint: endAudBorderPoint,
			startAud: startAud,
			endAud: endAud,
			sectorNumber: sectorNumber,
			client: client,
			constData: constData{
				positiveCoordinate: positiveCoordinate,
				negativeCoordinate: negativeCoordinate,
				axisX: axisX,
				axisY: axisY,
				widhtX: widhtX,
				heightX: heightX,
				widhtY: widhtY,
				heightY: heightY,
			},
		}
	}

func (a *audToAudController) controller() ([]models.Coordinates, appError.AppError) {
	//TODO: отрисовка начальных путей у стартовой и конечной аудиторий.
	err := a.getStartPoints()
	if err.Err != nil {
		err.Wrap("startPath")
		return nil, err
	}

	//TODO: отрисовка среднего пути от стартовой аудитории до конечной аудитории в притык.

	return a.points, appError.AppError{}
}