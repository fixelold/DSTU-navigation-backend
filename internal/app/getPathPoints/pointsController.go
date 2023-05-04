package getPathPoints

import (
	"navigation/internal/app/getPathPoints/middle"
	sectorToSector "navigation/internal/app/getPathPoints/sector2sector"
	"navigation/internal/app/getPathPoints/start"
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
)

var ()

const (
	AxisX = 1 // указывает на ось x.
	AxisY = 2 // указывает на ось y.

	WidhtX  = 10 // ширина на оси x.
	HeightX = 5  // высота на оси x.

	WidhtY  = 5  // ширина на оси y.
	HeightY = 10 // высота на оси y.

	plus  = 0 // значение будет положительным.
	minus = 1 // значение будет отрицательным.
)

// эти константы будут использовать для рассчета данных. Они буду передаваться в switch.

// const (
// 	audStartPoints  	= 1 // для начального пути от границ аудитории.
// 	auditory2Sector 	= 2
// 	path2Sector     	= 3
// 	sector2Sector   	= 4
// )

const (
	secondSector = 1
)

type controller struct {
	logger     *logging.Logger // логирования.
	client postgresql.Client      // для обращения к базе данных.

	StartAuditory string // номер начальной аудитории.
	EndAuditory   string // номер конечной аудитории.
	sectors       []int  // массив номеров секторов.

	transition       int
	transitionNumber int

	data data
	points []models.Coordinates
}

func NewPointsController(
	audStart, audEnd string,
	sectors []int,
	logger *logging.Logger,
	client postgresql.Client,
	transition, transitionNumber int) pointsController {
	return &controller{
		logger:           logger,
		client:       client,
		StartAuditory:    audStart,
		EndAuditory:      audEnd,
		sectors:          sectors,
		transition:       transition,
		transitionNumber: transitionNumber,
	}
}

type pointsController interface {
	controller() ([]models.Coordinates, appError.AppError)
	start(audNumber string) appError.AppError
	middle(entry, exit int) appError.AppError
	sector2sector() appError.AppError
	// getPointsAuditory2Sector(entry, exit int) appError.AppError
	// getPointsAuditory2Transition(entry, exit int) appError.AppError
	// getPointsSector2Sector() appError.AppError
}

func (p *controller) controller() ([]models.Coordinates, appError.AppError) {
	var response []models.Coordinates
	var err appError.AppError
	/*
		находим минимальное значение между номерами двух секторов.
		необходимо для внутренней логики.
	*/
	entry, exit := min(p.sectors[0], p.sectors[1])

	data, err := newData(p.StartAuditory, entry, exit, p.sectors[secondSector], p.logger, p.client, p.transition, p.transitionNumber)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	p.data = *data

	if p.transition == transitionYes {
		// err := p.getPointsAuditory2Transition(entry, exit)
		// if err.Err != nil {
		// 	return nil, err
		// }

		response = append(response, p.data.points...)
	} else if p.transition == transitionNo {

		err := p.start(p.StartAuditory)
		if err.Err != nil {
			return nil, err
		}

		err = p.middle(entry, exit)
		if err.Err != nil {
			return nil, err
		}

		err = p.sector2sector()
		if err.Err != nil {
			return nil, err
		}

		// err := p.getPointsAuditory2Sector(entry, exit)
		// if err.Err != nil {
		// 	return nil, err
		// }

		// response = append(response, p.data.points...)
		// err = p.getPointsSector2Sector()
		// if err.Err != nil {
		// 	return nil, err
		// }
		// response = append(response, p.data.points...)
		entry, exit = min(p.sectors[len(p.sectors)-1], p.sectors[len(p.sectors)-2])

		// получаем новый объекта типа 'data'. С данными этого типа будет происходить вся работа.
		newData, err := newData(p.EndAuditory, entry, exit, p.sectors[len(p.sectors)-1], p.logger, p.client, p.transition, p.transitionNumber)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return nil, err
		}
 
		p.data = *newData
		response = append(response, p.points...)
		p.points = []models.Coordinates{}

		err = p.start(p.EndAuditory)
		if err.Err != nil {
			return nil, err
		}

		err = p.middle(entry, exit)
		if err.Err != nil {
			return nil, err
		}

		// err = p.getPointsAuditory2Sector(entry, exit)
		// if err.Err != nil {
		// 	return nil, err
		// }
		// response = append(response, p.data.points...)
		// err = p.getPointsSector2Sector()
		// if err.Err != nil {
		// 	return nil, err
		// }
		// response = append(response, p.data.points...)

		response = append(response, p.points...)
	}

	return response, err
}

/*
расчет точек от аудитории до сектора
entry - входной сектор
exit - выходной сектор
entry всегда должен быть меньше exit
*/
func (p *controller) start(audNumber string) appError.AppError {
	start := start.NewStartController(p.data.audBorderPoints, p.client, audNumber, plus, minus, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY)
	data, err := start.StartPath()
	if err.Err != nil {
		err.Wrap("start")
		return err
	}
	p.points = append(p.points, data...)

	return appError.AppError{}
}

func (p *controller) middle(entry, exit int) appError.AppError {
	repository := NewRepository(p.client, p.logger)
	middle := middle.NewMiddleController(p.data.sectorNumber, p.client, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY, p.logger)
	borderSector, err := repository.getSectorBorderPoint(entry, exit)
	if err.Err != nil {
		err.Wrap("middle")
		return err
	}
	middle.Points = append(middle.Points, p.points...)
	
	data, err := middle.MiddlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("middle")
	}

	p.points = append(p.points, data...)

	return appError.AppError{}
}

func (p *controller) sector2sector() appError.AppError {
	repository := NewRepository(p.client, p.logger)
	sector2sector := sectorToSector.NewSectorToSectorController(p.data.sectorNumber, p.client, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY, p.logger)
	sector2sector.Points = append(sector2sector.Points, p.points...)
	sector2sector.OldAxis = 3 // delete
	for i := 1; i < len(p.sectors)-1; i++ {
		entry, exit := min(p.sectors[i], p.sectors[i+1])

		borderSector, err := repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return err
		}

		data, err := sector2sector.Sector2SectorPoints(borderSector, len(sector2sector.Points) - 1)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return err
		}

		// p.points = append(p.points, data...)
		sector2sector.Points = append(sector2sector.Points, data[len(data) - 1])

	}

	p.points = append(p.points, sector2sector.Points...)

	return appError.AppError{}
}


// func (p *controller) getPointsAuditory2Sector(entry, exit int) appError.AppError {

// 	// построение начального пути. От границы аудитории.
// 	// p.logger.Logger.Infoln("===================Start 'start auditory points'===================")
// 	err := p.data.setAudStartPoints()
// 	if err.Err != nil {
// 		err.Wrap("getPathPoints")
// 		return err
// 	}
// 	// p.logger.Logger.Infoln("===================End 'start auditory points'===================")

// 	borderSector, err := p.repository.getSectorBorderPoint(entry, exit)
// 	if err.Err != nil {
// 		err.Wrap("getPathPoints")
// 		return err
// 	}

// 	// p.logger.Logger.Infoln("===================Start 'auditory to end sector'===================")
// 	// построение пути вплоть до вхождение в область точек сектора.
// 	err = p.data.middlePoints(borderSector)
// 	if err.Err != nil {
// 		err.Wrap("getPathPoints")
// 		return err
// 	}

// 	// p.logger.Logger.Infoln("===================End 'auditory to end sector'===================")

// 	return appError.AppError{}
// }

// func (p *controller) getPointsAuditory2Transition(entry, exit int) appError.AppError {
// 	// построение начального пути. От границы аудитории.
// 	err := p.data.setAudStartPoints()
// 	if err.Err != nil {
// 		err.Wrap("getPathPoints")
// 		return err
// 	}

// 	borderSector, err := p.repository.getTransitionSectorBorderPoint(entry, exit)
// 	if err.Err != nil {
// 		err.Wrap("getPathPoints")
// 		return err
// 	}

// 	err = p.data.middlePoints(borderSector)
// 	if err.Err != nil {
// 		err.Wrap("getPathPoints")
// 		return err
// 	}

// 	return appError.AppError{}
// }

// func (p *controller) getPointsSector2Sector() appError.AppError {
	
// 	for i := 1; i < len(p.sectors)-1; i++ {

// 		entry, exit := min(p.sectors[i], p.sectors[i+1])

// 		borderSector, err := repository.getSectorBorderPoint(entry, exit)
// 		if err.Err != nil {
// 			err.Wrap("getPathPoints")
// 			return err
// 		}
// 		fmt.Println("iterator - ", i)
// 		if i == 2 {
// 			p.data.sectorType = 1
// 		}
// 		p.logger.Logger.Infoln("===================Start 'sector to sector'===================")
// 		err = p.data.sector2Sector(borderSector)
// 		if err.Err != nil {
// 			err.Wrap("getPathPoints")
// 			return err
// 		}
// 		p.logger.Logger.Infoln("===================End 'sector to sector'===================")
// 		// if i >= 2 {
// 		// 	p.data.points[len(p.data.points) - i - 1].Y -= 5
// 		// 	fmt.Println("data - ", p.data.points)
// 		// }

// 	}
// 	return appError.AppError{}
// }

func min(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
