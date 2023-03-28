package getPathPoints

import (
	"navigation/internal/logging"
)

const (
	AxisX = 1
	AxisY = 2

	WidhtX  = 40
	HeightX = 20

	WidhtY  = 20
	HeightY = 40

	plus  = 0
	minus = 1
)

const (
	audStartPoints = 1
	// Auditory2Sector = 1
	// Path2Sector     = 2
	// Sector2Sector   = 3
)

type pointsController struct {
	logger     *logging.Logger // логирования.
	repository Repository      // для обращения к базе данных.

	audStart string // номер начальной аудитории.
	audEnd   string // номер конечной аудитории.
	sectors  []int  // массив номеров секторов
}

func NewPointsController(logger *logging.Logger, repository Repository) *pointsController {
	return &pointsController{
		logger:     logger,
		repository: repository,
	}
}

func (p *pointsController) getPathPoints() error {
	entry, exit, err := min(p.sectors[0], p.sectors[1])
	if err != nil {
		return err
	}

	data, err := newData(p.audStart, entry, exit, p.sectors[1], p.logger, p.repository)
	if err != nil {
		return err
	}

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
