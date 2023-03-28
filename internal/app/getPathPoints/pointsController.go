package getPathPoints

import (
	"navigation/internal/logging"
	"navigation/internal/models"
)

type pointsController struct {
	logger     *logging.Logger // логирования.
	repository Repository      // для обращения к базе данных.

	audStart string               // номер начальной аудитории.
	audEnd   string               // номер конечной аудитории.
	sectors  []int                // массив номеров секторов
	points   []models.Coordinates // массив координат. Для построения пути.
}

func NewPointsController(logger *logging.Logger, repository Repository) *pointsController {
	return &pointsController{
		logger:     logger,
		repository: repository,
	}
}

func (p *pointsController) getPathPoints() error {

	return nil
}

// TODO: добавить ошибки, если переменные одинаковые.
func min(a, b int) (int, int, error) {
	if a < b {
		return a, b, nil
	} else if a > b {
		return b, a, nil
	}
	return a, b, nil
}