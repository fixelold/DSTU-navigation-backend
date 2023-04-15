package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

var ()

const (
	AxisX = 1 // указывает на ось x.
	AxisY = 2 // указывает на ось y.

	WidhtX  = 15 // ширина на оси x.
	HeightX = 10 // высота на оси x.

	WidhtY  = 10 // ширина на оси y.
	HeightY = 15 // высота на оси y.

	plus  = 0 // значение будет положительным.
	minus = 1 // значение будет отрицательным.
)

// эти константы будут использовать для рассчета данных. Они буду передаваться в switch.
const (
	audStartPoints  = 1 // для начального пути от границ аудитории.
	auditory2Sector = 2
	path2Sector     = 3
	sector2Sector   = 4
)

const (
	secondSector = 1
)

type controller struct {
	logger     *logging.Logger // логирования.
	repository Repository      // для обращения к базе данных.

	StartAuditory string // номер начальной аудитории.
	EndAuditory   string // номер конечной аудитории.
	sectors       []int  // массив номеров секторов.

	transition       int
	transitionNumber int

	data data
	// points []models.Coordinates
}

func NewPointsController(
	audStart, audEnd string,
	sectors []int,
	logger *logging.Logger,
	repository Repository,
	transition, transitionNumber int) pointsController {
	return &controller{
		logger:           logger,
		repository:       repository,
		StartAuditory:    audStart,
		EndAuditory:      audEnd,
		sectors:          sectors,
		transition:       transition,
		transitionNumber: transitionNumber,
	}
}

type pointsController interface {
	controller() ([]models.Coordinates, appError.AppError)
	getPointsAuditory2Sector(entry, exit int) appError.AppError
	getPointsAuditory2Transition(entry, exit int) appError.AppError
	getPointsSector2Sector() appError.AppError
}

func (p *controller) controller() ([]models.Coordinates, appError.AppError) {
	var response []models.Coordinates
	var err appError.AppError
	/*
		находим минимальное значение между номерами двух секторов.
		необходимо для внутренней логики.
	*/
	entry, exit := min(p.sectors[0], p.sectors[1])

	data, err := newData(p.StartAuditory, entry, exit, p.sectors[secondSector], p.logger, p.repository, p.transition, p.transitionNumber)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	p.data = *data

	if p.transition == transitionYes {
		err := p.getPointsAuditory2Transition(entry, exit)
		if err.Err != nil {
			return nil, err
		}

	} else if p.transition == transitionNo {
		err := p.getPointsAuditory2Sector(entry, exit)
		if err.Err != nil {
			return nil, err
		}

		response = append(response, p.data.points...)
		err = p.getPointsSector2Sector()
		if err.Err != nil {
			return nil, err
		}
		response = append(response, p.data.points...)
		entry, exit = min(p.sectors[len(p.sectors)-1], p.sectors[len(p.sectors)-2])

		// получаем новый объекта типа 'data'. С данными этого типа будет происходить вся работа.
		data, err := newData(p.EndAuditory, entry, exit, p.sectors[len(p.sectors)-1], p.logger, p.repository, p.transition, p.transitionNumber)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return nil, err
		}

		p.data = *data

		err = p.getPointsAuditory2Sector(entry, exit)
		if err.Err != nil {
			return nil, err
		}
		response = append(response, p.data.points...)
		err = p.getPointsSector2Sector()
		if err.Err != nil {
			return nil, err
		}
		response = append(response, p.data.points...)
	}

	return response, err
}

/*
расчет точек от аудитории до сектора
entry - входной сектор
exit - выходной сектор
entry всегда должен быть меньше exit
*/
func (p *controller) getPointsAuditory2Sector(entry, exit int) appError.AppError {

	// построение начального пути. От границы аудитории.
	err := p.data.setAudStartPoints()
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return err
	}

	borderSector, err := p.repository.getSectorBorderPoint(entry, exit)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return err
	}

	// построение пути вплоть до вхождение в область точек сектора.
	err = p.data.middlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return err
	}

	return appError.AppError{}
}

func (p *controller) getPointsAuditory2Transition(entry, exit int) appError.AppError {
	// построение начального пути. От границы аудитории.
	err := p.data.setAudStartPoints()
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return err
	}

	borderSector, err := p.repository.getTransitionSectorBorderPoint(entry, exit)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return err
	}

	err = p.data.middlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return err
	}

	return appError.AppError{}
}

func (p *controller) getPointsSector2Sector() appError.AppError {

	for i := 1; i < len(p.sectors)-1; i++ {

		entry, exit := min(p.sectors[i], p.sectors[i+1])

		borderSector, err := p.repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return err
		}

		err = p.data.sector2Sector(borderSector)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return err
		}
	}
	return appError.AppError{}
}

func min(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
