package drawPath

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
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

type Path struct {
	AudienceCoordinates models.Coordinates
	AudienceBorderPoint models.Coordinates
	SectorBorderPoint   models.Coordinates
	SectorNumber        int
	AudienceNumber      string
	Path                []models.Coordinates
	Repository          Repository
	logger              *logging.Logger
}

func NewPath(
	audienceCoordinates,
	audienceBorderPoint,
	sectorBorderPoint models.Coordinates,
	sectorNumber int,
	audienceNumber string,
	repository Repository,
	logger *logging.Logger) *Path {
	return &Path{
		AudienceCoordinates: audienceCoordinates,
		AudienceBorderPoint: audienceBorderPoint,
		SectorBorderPoint:   sectorBorderPoint,
		SectorNumber:        sectorNumber,
		AudienceNumber:      audienceNumber,
		Repository:          repository,
		logger:              logger,
	}
}

var (
	User000004 = appError.NewError("drawPath", "GetSelector", "Input does not match desired length", "-", "US-000004")
)

func (d *Path) DrawInitPath() error {

	err := d.drawPathAuditory()
	if err != nil {
		return err
	}

	err = d.drawPathSector()
	if err != nil {
		return err
	}

	return nil
}

func (d *Path) drawPathAuditory() error {
	var err error

	axis := d.defenitionAxis(d.AudienceBorderPoint.Widht, d.AudienceBorderPoint.Height)

	err = d.getPoints(axis)
	if err != nil {
		return err
	}

	return nil
}

func (d *Path) drawPathSector() error {
	var widht, height int
	iterator := 0
	axis := d.defenitionAxis(d.SectorBorderPoint.Widht, d.SectorBorderPoint.Height)
	boolean := true

	switch axis {
	case AxisX:
		widht = WidhtY
		height = HeightY
	case AxisY:
		widht = WidhtX
		height = HeightX
	default:
		d.logger.Errorln("Function drawPathSector. Error switch default")
	}

	for boolean {
		if d.checkPath2Sector(d.Path[iterator], axis) {
			switch axis {
			case AxisX:
				widht = d.SectorBorderPoint.X - d.Path[iterator].X
				height = HeightX
			case AxisY:
				widht = WidhtY
				height =  d.SectorBorderPoint.Y - d.Path[iterator].Y
			default:
				d.logger.Errorln("Function drawPathSector. Error switch default")
			}
			d.pathAlignment(d.SectorBorderPoint, axis)
			points := d.getPoints2Sector(widht, height, axis, d.Path[iterator], d.SectorBorderPoint)
			if points == (models.Coordinates{}) {
				return User000004
			}

			d.Path = append(d.Path, points)
			boolean = false
		} else {

			points := d.getPoints2Sector(widht, height, axis, d.Path[iterator], d.SectorBorderPoint)
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

func (d *Path) getPoints2Sector(widht, heihgt, axis int, path, borderPoint models.Coordinates) models.Coordinates {

	switch axis {
	case AxisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y + path.Height - HeightX)}
		sectorPoints := (borderPoint.Y + (borderPoint.Height + borderPoint.Y)) / 2
		if sectorPoints > path.X {
			points.Widht = widht
			points.Height = heihgt
			return points
		} else {
			points.Widht = -widht
			points.Height = -heihgt
			return points
		}
	case AxisY:
		points := models.Coordinates{
			X: (path.X + path.Widht - WidhtY),
			Y: (path.Y)}
		sectorPoints := (borderPoint.X + (borderPoint.Widht + borderPoint.X)) / 2
		if sectorPoints > path.X {
			points.Widht = widht
			points.Height = heihgt
			return points
		} else {
			points.Widht = -widht
			points.Height = -heihgt
			return points
		}
	default:
		d.logger.Errorln("Function - getPoint2Sector. Error - switch default")
		return models.Coordinates{}
	}
}

func (d *Path) getDrawSector2Sector(path, sectorBorderPoint models.Coordinates, axis int) models.Coordinates {
	d.logger.Infoln("draw init path - get draw points 2 sector")

	switch axis {
	case AxisX:
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			points := models.Coordinates{
				X: (path.X + path.Widht),
				Y: (path.Y + path.Height)}
			points.Widht = sectorBorderPoint.X - (path.X + path.Widht)
			points.Height = HeightX
			return points
		} else {
			points := models.Coordinates{
				X: (path.X + path.Widht),
				Y: (path.Y)}
			points.Widht = -sectorBorderPoint.X - (path.X + path.Widht)
			points.Height = -HeightX
			return points
		}
	case AxisY:
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		if sectorPoints > path.Y {
			points := models.Coordinates{
				X: (path.X + path.Widht),
				Y: (path.Y + path.Height)}
			points.Widht = WidhtY
			points.Height = sectorBorderPoint.Y - (path.Y + path.Height)
			return points
		} else {
			points := models.Coordinates{
				X: (path.X),
				Y: (path.Y + path.Height)}
			points.Widht = -WidhtY
			points.Height = sectorBorderPoint.Y - (path.Y + path.Height)
			return points
		}
	default:
		return models.Coordinates{}
	}
}

func (d *Path) checkPath2Sector(path models.Coordinates, axis int) bool {
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

func (d *Path) defenitionAxis(width, height int) int {
	if width == 1 {
		return AxisX
	} else if height == 1 {
		return AxisY
	} else {
		return 0
	}
}

// Выравнивание пути
func (d *Path) pathAlignment(sectorBorderPoint models.Coordinates, axis int) {
	lenght := len(d.Path)
	path := d.Path[lenght-1]
	switch axis {
	case AxisX:
		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.Y + (sectorBorderPoint.Height + sectorBorderPoint.Y)) / 2
		if sectorPoints > path.Y {
			points.Widht = WidhtY
			points.Height = sectorPoints - path.Y
			d.Path[lenght-1].Height = points.Height
		} else if sectorPoints < path.Y {
			points.Widht = WidhtY
			points.Height = sectorPoints - path.Y
			d.Path[lenght-1].Height = points.Height
		}
	case AxisY:

		points := models.Coordinates{
			X: (path.X),
			Y: (path.Y)}
		sectorPoints := (sectorBorderPoint.X + (sectorBorderPoint.Widht + sectorBorderPoint.X)) / 2
		if sectorPoints > path.X {
			points.Widht = sectorPoints - path.X
			points.Height = HeightX
			d.Path[lenght-1].Widht = points.Widht
		} else if sectorPoints < path.X {
			points.Widht = sectorPoints - path.X
			points.Height = HeightX
			d.Path[lenght-1].Widht = points.Widht
		}
	default:
		d.logger.Errorln("Path Alignment default")
	}
}
