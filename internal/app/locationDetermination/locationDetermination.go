package locationDetermination

import (
	"context"
	"errors"
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
	"strconv"
	"strings"

	"github.com/jackc/pgconn"
)

var (
	User000001 = appError.NewError("locationDetermination", "GetSelector", "Input does not match desired length", "-", "US-000001")
	User000002 = appError.NewError("locationDetermination", "GetSelector", "Not convert string to int", "-", "US-000002")
	User000003 = appError.NewError("locationDetermination", "GetSelector", "Errir in getting sector", "-", "US-000003")
)

type location struct {
	audienceNumber string
	building       int
	client         postgresql.Client
	logger         *logging.Logger
}

func NewLocation(audienceNumber string, logger *logging.Logger, client postgresql.Client) *location {
	return &location{
		audienceNumber: audienceNumber,
		logger:         logger,
		client: client,
	}
}

func (l *location) GetSector() (uint, error) {
	var err error

	splitText := strings.Split(l.audienceNumber, "-")
	if len(splitText) != 2 {
		l.logger.Errorf("function GetSelector. Input does not match desired length expected: %d, received: %d", 2, len(splitText))
		return 0, User000001
	}

	l.audienceNumber = splitText[1]

	l.building, err = strconv.Atoi(splitText[0])
	if err != nil {
		l.logger.Errorf("function GetSelector not convert string to int, err: %s", err)
		User000002.ChangeDescription(err.Error())
		return 0, User000002
	}

	sector, err := l.getSector()
	if err != nil {
		l.logger.Errorf("the getSector function call returned %s", err.Error())
		User000003.ChangeDescription(err.Error())
		return 0, User000003
	}

	return sector, nil
}

func (l *location) getSector() (uint, error) {
	var sector models.Sector
	req :=
		`SELECT 
	sector.number
	FROM auditorium
	JOIN sector ON auditorium.id_sector=sector.id
	JOIN floor ON sector.id_floor=floor.id
	JOIN building ON floor.id_building=building.id
	WHERE auditorium.number = $1 
	AND building.number = $2;`

	tx, err := l.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		l.logger.Tracef("can't start transaction: %s", err.Error())
		return 0, err
	}

	err = tx.QueryRow(
		context.Background(),
		req,
		l.audienceNumber,
		l.building).Scan(&sector.Number)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			l.logger.Error(newErr)
			return 0, newErr
		}
		l.logger.Error(err)
		return 0, err
	}
	_ = tx.Commit(context.Background())
	return sector.Number, nil
}
