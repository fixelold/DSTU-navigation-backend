package getPathPoints

import (
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

var ()

const (
	AxisX = 1 // указывает на ось x.
	AxisY = 2 // указывает на ось y.

	WidhtX  = 40 // ширина на оси x.
	HeightX = 20 // высота на оси x.

	WidhtY  = 20 // ширина на оси y.
	HeightY = 40 // высота на оси y.

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

type pointsController struct {
	logger     *logging.Logger // логирования.
	repository Repository      // для обращения к базе данных.

	audStart string // номер начальной аудитории.
	audEnd   string // номер конечной аудитории.
	sectors  []int  // массив номеров секторов
}

func NewPointsController(audStart, audEnd string, sectors []int, logger *logging.Logger, repository Repository) *pointsController {
	return &pointsController{
		logger:     logger,
		repository: repository,
		audStart:   audStart,
		audEnd:     audEnd,
		sectors:    sectors,
	}
}

func (p *pointsController) getPathPoints() ([]models.Coordinates, appError.AppError) {
	var err appError.AppError
	/* находим минимальное значение между номерами двух секторов.
	   необходимо для внутренней логики.
	*/
	entry, exit := min(p.sectors[0], p.sectors[1])

	// получаем новый объекта типа 'data'. С данными этого типа будет происходить вся работа.
	data, err := newData(p.audStart, entry, exit, p.sectors[1], p.logger, p.repository)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	// построение начального пути. От границы аудитории.
	err = data.setAudStartPoints()
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	borderSector, err := p.repository.getSectorBorderPoint(entry, exit)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	// построение пути вплоть до вхождение в область точек сектора.
	err = data.middlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	for i := 1; i < len(p.sectors)-1; i++ {

		entry, exit = min(p.sectors[i], p.sectors[i+1])

		borderSector, err := p.repository.getSectorBorderPoint(entry, exit)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return nil, err
		}

		err = data.sector2Sector(borderSector)
		if err.Err != nil {
			err.Wrap("getPathPoints")
			return nil, err
		}
	}

	entry, exit = min(p.sectors[len(p.sectors)-1], p.sectors[len(p.sectors)-2])

	// получаем новый объекта типа 'data'. С данными этого типа будет происходить вся работа.
	dataEnd, err := newData(p.audEnd, entry, exit, p.sectors[len(p.sectors)-1], p.logger, p.repository)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	err = dataEnd.setAudStartPoints()
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	borderSector, err = p.repository.getSectorBorderPoint(entry, exit)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	// построение пути вплоть до вхождение в область точек сектора.
	err = dataEnd.middlePoints(borderSector)
	if err.Err != nil {
		err.Wrap("getPathPoints")
		return nil, err
	}

	fmt.Println("data - ", data)
	fmt.Println("dataEnd - ", dataEnd)
	data.points = append(data.points, dataEnd.points...)
	fmt.Println("mega mega data - ", data.points)

	return data.points, err
}

func min(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
