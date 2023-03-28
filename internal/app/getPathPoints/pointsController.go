package getPathPoints

import (
	"navigation/internal/logging"
)

type pointsController struct {
	logger     *logging.Logger
	repository Repository
}

func NewPointsController(logger *logging.Logger, repository Repository) *pointsController {
	return &pointsController{
		logger: logger,
		repository: repository,
	}
}

func (p *pointsController) getPathPoints() error {

	return nil
}