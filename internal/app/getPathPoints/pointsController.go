package getPathPoints

import (
	"strconv"

	"navigation/internal/app/getPathPoints/audToAud"
	"navigation/internal/app/getPathPoints/audToTransition"
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
	aud2Aud() ([]models.Coordinates, appError.AppError)
	// start(audNumber string, typeTransition int) appError.AppError
	// middle(entry, exit int) appError.AppError
	// sector2sector() appError.AppError
	transitionController() ([]models.Coordinates, appError.AppError)
}

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


func (p *controller) controller() ([]models.Coordinates, appError.AppError) {
	var response []models.Coordinates
	var err appError.AppError
	/*
		находим минимальное значение между номерами двух секторов.
		необходимо для внутренней логики.
	*/

	entry, exit := min(p.sectors[0], p.sectors[1])
	data, err := newData(p.StartAuditory, "", entry, exit, p.sectors[secondSector], p.logger, p.client, p.transition, p.transitionNumber)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}
	p.data = *data

	err = p.start(p.StartAuditory)
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

	entry, exit = min(p.sectors[len(p.sectors)-1], p.sectors[len(p.sectors)-2])

	// получаем новый объекта типа 'data'. С данными этого типа будет происходить вся работа.
	newData, err := newData(p.EndAuditory, "", entry, exit, p.sectors[len(p.sectors)-1], p.logger, p.client, p.transition, p.transitionNumber)
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

	response = append(response, p.points...)

	return response, err
}

/*
расчет точек от аудитории до сектора
entry - входной сектор
exit - выходной сектор
entry всегда должен быть меньше exit
*/

// func (p *controller) transitionStart(audNumber string) appError.AppError {
	
// }

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
	var borderSector models.Coordinates
	var err appError.AppError
	repository := NewRepository(p.client, p.logger)
	middle := middle.NewMiddleController(p.transition, p.sectors[0], p.data.sectorNumber, p.client, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY, p.logger)
	if p.transition == aud2Aud {
		borderSector = p.data.sectorBorderPoints
	} else {
		borderSector, err = repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("middle")
			return err
		}
	}
	middle.Points = append(middle.Points, p.points...)
	data, err := middle.MiddlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("middle")
	}
	p.points = []models.Coordinates{}
	p.points = append(p.points, data...)

	return appError.AppError{}
}

func (p *controller) middleToTransition(entry int) appError.AppError {
	var borderSector models.Coordinates
	repository := NewRepository(p.client, p.logger)
	var err appError.AppError
	middle := audToTransition.NewAudToTransition(p.transition, p.sectors[0], p.data.sectorNumber, p.client, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY, p.logger)
	borderSector, err = repository.getTransitionSectorBorderPoint(entry)
	if err.Err != nil {
		err.Wrap("middle to transition")
		return err
	}

	middle.Points = append(middle.Points, p.points...)
	
	data, err := middle.MiddlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("middle to transition")
	}

	p.points = append(p.points, data...)

	return appError.AppError{}
}

func (p *controller) sector2sector() appError.AppError {
	var data []models.Coordinates
	repository := NewRepository(p.client, p.logger)
	sector2sector := sectorToSector.NewSectorToSectorController(p.data.sectorNumber, p.client, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY, p.logger)
	data = append(data, p.points...)
	sector2sector.OldAxis = 3 // delete
	sector2sector.Points = append(sector2sector.Points, data[1:]...)
	for i := 1; i < len(p.sectors)-1; i++ {
		entry, exit := min(p.sectors[i], p.sectors[i+1])

		borderSector, err := repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return err
		}

		_, err = sector2sector.Sector2SectorPoints(borderSector, len(sector2sector.Points) - 1)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return err
		}
		// p.points = append(p.points, data...)
		// fmt.Println("data - ",  data[len(data) - 1])
		// sector2sector.Points = append(sector2sector.Points, data[len(data) - 1])

	}

	p.points = append(p.points, sector2sector.Points...)

	return appError.AppError{}
}

func (p *controller) transitionController() ([]models.Coordinates, appError.AppError) {
	var response []models.Coordinates
	// var err appError.AppError
	var exit int
	if len(p.sectors) == 1 {
		exit = p.sectors[0]
	} else {
		_, exit = min(p.sectors[0], p.sectors[1])
	}

	if p.transition == stair || len(strconv.Itoa(exit)) == 4 {
		entry, exit := p.sectors[0], p.sectors[1]
		data, err := newData(p.StartAuditory, "", entry, exit, p.sectors[secondSector], p.logger, p.client, stair, p.transitionNumber)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return nil, err
		}
		p.data = *data
		// entry, exit := p.sectors[0], p.sectors[1]

		err = p.start(p.StartAuditory)
		if err.Err != nil {
			return nil, err
		}

		err = p.middleToTransition(exit)
		if err.Err != nil {
			return nil, err
		}

		response = append(response, p.points...)

	} else if p.transition == elevator {
		entry, exit := min(p.sectors[0], p.sectors[1])
		data, err := newData(p.StartAuditory, "", entry, exit, p.sectors[secondSector], p.logger, p.client, elevator, p.transitionNumber)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return nil, err
		}
		p.data = *data

		err = p.start(p.StartAuditory)
		if err.Err != nil {
			return nil, err
		}

		if len(strconv.Itoa(exit)) == 4 {
			err = p.middleToTransition(entry)
			if err.Err != nil {
				return nil, err
			}
		} else {
			err = p.middle(entry, exit)
			if err.Err != nil {
				return nil, err
			}

			// это сделано т.к с фронта возвращается не (143, 142, 141), а (1043, 143, 142, 141)
			entry, exit = min(p.sectors[0], p.sectors[1])
			//p.transition = stair возможно, надо будет это раскоментить или что-то сделать!!!
			p.transitionNumber = p.sectors[len(p.sectors) - 1]
			data, err = newData(p.StartAuditory, "", entry, exit, p.sectors[secondSector], p.logger, p.client, transitionToAud, p.transitionNumber)
			if err.Err != nil {
				err.Wrap("getPathPoints")
				return nil, err
			}
			p.data = *data
			response = append(response, p.points...)
			p.points = []models.Coordinates{}

			err = p.start(p.StartAuditory)
			if err.Err != nil {
				return nil, err
			}

			err = p.middle(entry, exit)
			if err.Err != nil {
				return nil, err
			}
		}

		response = append(response, p.points...)

	} else if p.transition == transitionToAud {
		if len(strconv.Itoa(p.sectors[0])) == 4{
			temp := p.sectors[1:]
			p.sectors = temp
		}

		if len(p.sectors) == 1 {
			
			p.transition = stair
			entry := p.sectors[0]
			data, err := newData(p.EndAuditory, "", entry, entry, p.sectors[0], p.logger, p.client, p.transition, p.transitionNumber)
			if err.Err != nil {
				err.Wrap("getPathPoints")
				return nil, err
			}
			p.data = *data
	
	
			err = p.start(p.EndAuditory)
			if err.Err != nil {
				return nil, err
			}

			err = p.middleToTransition(p.transitionNumber)
			if err.Err != nil {
				return nil, err
			}
	
			response = append(response, p.points...)

		} else {
			// это сделано т.к с фронта возвращается не (143, 142, 141), а (1043, 143, 142, 141)
			entry, exit := min(p.sectors[0], p.sectors[1])
			//p.transition = stair возможно, надо будет это раскоментить или что-то сделать!!!
			data, err := newData(p.StartAuditory, "", entry, exit, p.sectors[secondSector], p.logger, p.client, p.transition, p.transitionNumber)
			if err.Err != nil {
				err.Wrap("getPathPoints")
				return nil, err
			}
			p.data = *data

			err = p.start(p.StartAuditory)
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

			entry, exit = min(p.sectors[len(p.sectors)-1], p.sectors[len(p.sectors)-2])

			// получаем новый объекта типа 'data'. С данными этого типа будет происходить вся работа.
			p.transition = noTransition
			newData, err := newData(p.EndAuditory, "", entry, exit, p.sectors[len(p.sectors)-1], p.logger, p.client, p.transition, p.transitionNumber)
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
	
			response = append(response, p.points...)
		}
	}

	return response, appError.AppError{}
}

func (p *controller) aud2Aud() ([]models.Coordinates, appError.AppError) {
	var response []models.Coordinates
	p.transition = aud2Aud
	entry := p.sectors[0]

	data, err := newData( p.EndAuditory, p.StartAuditory, entry, entry, p.sectors[len(p.sectors)-1], p.logger, p.client, p.transition, p.transitionNumber)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	p.data = *data

	err = p.start(p.EndAuditory)
	if err.Err != nil {
		return nil, err
	}

	newdata, err := newData(p.StartAuditory, p.EndAuditory, entry, entry, p.sectors[0], p.logger, p.client, p.transition, p.transitionNumber)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}
	p.data = *newdata
	endPoints := p.points[0]
	p.points = []models.Coordinates{}

	err = p.start(p.StartAuditory)
	if err.Err != nil {
		return nil, err
	}

	err = p.middleAudToAud(endPoints)
	if err.Err != nil {
		return nil, err
	}

	response = append(response, p.points...)


	return response, appError.AppError{}
}

func (p *controller) middleAudToAud(endPints models.Coordinates) appError.AppError {
	var borderSector models.Coordinates
	// repository := NewRepository(p.client, p.logger)
	var err appError.AppError
	middle := audToAud.NewAudToAud(p.transition, p.sectors[0], p.data.sectorNumber, endPints, p.client, AxisX, AxisY, WidhtX, HeightX, WidhtY, HeightY, p.logger)
	borderSector = p.data.sectorBorderPoints
	middle.Points = append(middle.Points, p.points...)
	
	data, err := middle.MiddlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("middle to transition")
	}

	p.points = append(p.points, data...)

	return appError.AppError{}
}
func min(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
