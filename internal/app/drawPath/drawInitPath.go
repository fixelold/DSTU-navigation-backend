package drawPath

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
)

const (
	AxisX = 1
	AxisY = 2

	WidhtX  = 130
	HeightX = 30

	WidhtY  = 30
	HeightY = 130

	plus  = 0
	minus = 1
)

type drawPathAud2Sector struct {
	AudienceCoordinates models.Coordinates
	AudienceBorderPoint models.Coordinates
	SectorBorderPoint models.Coordinates
	SectorNumber        int
	AudienceNumber      string
	Path                []models.Coordinates
	Repository          Repository
	// добавить repository, чтобы можно было обращаться в БД.
}

func NewDrawPathAud2Sector(
	audienceCoordinates, 
	audienceBorderPoint,
	sectorBorderPoint models.Coordinates, 
	sectorNumber int, 
	audienceNumber string, 
	repository Repository) *drawPathAud2Sector {
	return &drawPathAud2Sector{
		AudienceCoordinates: audienceCoordinates,
		AudienceBorderPoint: audienceBorderPoint,
		SectorBorderPoint: sectorBorderPoint,
		SectorNumber:        sectorNumber,
		AudienceNumber:      audienceNumber,
		Repository:          repository,
	}
}

var (
	User000004 = appError.NewError("drawPath", "GetSelector", "Input does not match desired length", "-", "US-000004")
)

func (d *drawPathAud2Sector) DrawInitPath() error {

	err := d.drawPathAuditory()
	if err != nil {
		return err
	}

	return nil
}

func (d *drawPathAud2Sector) drawPathAuditory() error {
	var err error
	axis := d.defenitionAxis(d.AudienceBorderPoint.Widht, d.AudienceBorderPoint.Height)

	switch axis {

	case AxisX:
		err := d.drawAudX()
		if err != nil {
			logging.GetLogger().Errorln("DrawPathAuditory case AxisX. Error - ", err)
			return err
		}

	case AxisY:
		err := d.drawAudY()
		if err != nil {
			logging.GetLogger().Errorln("DrawPathAuditory case AxisY. Error - ", err.Error())
			return err
		}

	default:
		logging.GetLogger().Errorln("DrawPathAuditory case default. Error - ", err)
		err = User000004
	}

	return err
}

func (d *drawPathAud2Sector) drawPathSector() error {
	iterator := 0
	axis := d.defenitionAxis(d.SectorBorderPoint.Widht, d.SectorBorderPoint.Height)
	boolean := true

	for boolean {
		if d.checkPath2Sector(d.Path[iterator], axis) {
			points := d.getDrawPoints2Sector(d.Path[iterator], axis)

			d.Path = append(d.Path, points)
			boolean = false
		} else {
			// определяем в каком направлении рисовать
			points := d.getDrawPoints(d.Path[iterator], axis)
			if points == (models.Coordinates{}) {
				return User000004
			}

			ok, err := d.Repository.checkBorderAud(points)
			if err != nil {
				return User000004
			}

			ok2, err := d.Repository.checkBorderSector(points)
			if err != nil {
				return User000004
			}

			if !ok && !ok2 {
				//TODO написать изменения направления или типо что-то такого
			} 

			d.Path = append(d.Path, points)
		}

		iterator += 1
	}

	return nil
}

func (d *drawPathAud2Sector) getDrawPoints(path models.Coordinates, axis int) (models.Coordinates) {
	points := models.Coordinates{
		X: (path.X + path.Widht),
		Y: (path.X + path.Widht),}

	switch axis {
	case AxisX:
		//y := (d.Path[0].X + d.Path[0].Widht) - HeightX
		sectorPoints := (d.SectorBorderPoint.Y + (d.SectorBorderPoint.Height + d.SectorBorderPoint.Y))/ 2
		if sectorPoints > path.X {
			points.Widht = WidhtY
			points.Height = HeightY
			return points
		} else {
			points.Widht = -WidhtY
			points.Height = -HeightY
			return points
		}
	case AxisY:
		sectorPoints := (d.SectorBorderPoint.X + (d.SectorBorderPoint.Widht + d.SectorBorderPoint.X))/ 2
		if sectorPoints > path.X {
			points.Widht = WidhtX
			points.Height = HeightX
			return points
		} else {
			points.Widht = -WidhtX
			points.Height = -HeightX
			return points
		}
	default:
		return models.Coordinates{}
	}
}

func (d *drawPathAud2Sector) getDrawPoints2Sector(path models.Coordinates, axis int) (models.Coordinates) {
	points := models.Coordinates{
		X: (path.X + path.Widht),
		Y: (path.Y + path.Height),}

	switch axis {
	case AxisX:
		sectorPoints := (d.SectorBorderPoint.Y + (d.SectorBorderPoint.Height + d.SectorBorderPoint.Y))/ 2
		if sectorPoints > path.X {
			points.Widht = WidhtY
			points.Height = d.SectorBorderPoint.Height - (path.Y + path.Height)
			return points
		} else {
			points.Widht = -WidhtY
			points.Height = -d.SectorBorderPoint.Height - (path.Y + path.Height)
			return points
		}
	case AxisY:
		sectorPoints := (d.SectorBorderPoint.X + (d.SectorBorderPoint.Widht + d.SectorBorderPoint.X))/ 2
		if sectorPoints > path.X {
			points.Widht = d.SectorBorderPoint.Widht - (path.X + path.Widht)
			points.Height = HeightX
			return points
		} else {
			points.Widht = -d.SectorBorderPoint.Widht - (path.X + path.Widht)
			points.Height = -HeightX
			return points
		}
	default:
		return models.Coordinates{}
	}
}

func (d *drawPathAud2Sector) checkPath2Sector(path models.Coordinates, axis int) bool {
	switch axis {
	case AxisX:
		ph := path.Y + path.Height
		y1 := d.SectorBorderPoint.Y
		y2 := d.SectorBorderPoint.Y + d.SectorBorderPoint.Height
		if y1 <= ph && ph <= y2 {
			return true
		} else {
			return false
		}
	case AxisY:
		ph := path.X + path.Widht
		x1 := d.SectorBorderPoint.X
		x2 := d.SectorBorderPoint.X + d.SectorBorderPoint.Widht
		if x1 <= ph && ph <= x2 {
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

func (d *drawPathAud2Sector) defenitionAxis(width, height int) int {
	if width == 1 {
		return AxisX
	} else if height == 1 {
		return AxisY
	} else {
		return 0
	}
}
