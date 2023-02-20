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
	SectorNumber        int
	AudienceNumber      string
	Path                []int
	Repository          Repository
	// добавить repository, чтобы можно было обращаться в БД.
}

func NewDrawPathAud2Sector(audienceCoordinates, audienceBorderPoint models.Coordinates, sectorNumber int, audienceNumber string) *drawPathAud2Sector {
	return &drawPathAud2Sector{
		AudienceCoordinates: audienceCoordinates,
		AudienceBorderPoint: audienceBorderPoint,
		SectorNumber:        sectorNumber,
		AudienceNumber:      audienceNumber,
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
	axis := d.defenitionAxis()

	switch axis {

	case AxisX:
		err := d.drawX()
		if err != nil {
			logging.GetLogger().Errorln("DrawPathAuditory case AxisX. Error - ", err)
			return err
		}

	case AxisY:
		err := d.drawY()
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

func (d *drawPathAud2Sector) defenitionAxis() int {
	if d.AudienceBorderPoint.Widht == 1 {
		return AxisX
	} else if d.AudienceBorderPoint.Height == 1 {
		return AxisY
	} else {
		return 0
	}
}
