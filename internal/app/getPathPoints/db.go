package getPathPoints

import (
	"context"
	"errors"
	"fmt"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"

	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) getAudPoints(number string) (models.Coordinates, error) {
	var position models.Coordinates
	request :=
		`SELECT x, y, widht, height 
	FROM auditorium_position 
	JOIN auditorium 
	ON auditorium_position.id_auditorium = auditorium.id 
	WHERE auditorium.number = $1;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return nil, err
	}

	err = tx.QueryRow(
		context.Background(),
		request,
		number).Scan(
		&position.X,
		&position.Y,
		&position.Widht,
		&position.Height)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}
		r.logger.Error(err)
		return nil, err
	}
	_ = tx.Commit(context.Background())
	return &position, nil
}

func (r *repository) getAudBorderPoint(number string) (models.Coordinates, error) {
	var borderPoint models.Coordinates
	request :=
		`SELECT x, y, widht, height 
	FROM aud_border_points 
	JOIN auditorium 
	ON aud_border_points.id_auditorium = auditorium.id 
	WHERE auditorium.number = $1;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return nil, err
	}

	err = tx.QueryRow(
		context.Background(),
		request,
		number).Scan(
		&borderPoint.X,
		&borderPoint.Y,
		&borderPoint.Widht,
		&borderPoint.Height)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}
		r.logger.Error(err)
		return nil, err
	}
	_ = tx.Commit(context.Background())
	return &borderPoint, nil
}

func (r *repository) getSectorBorderPoint(entry, exit int) (models.Coordinates, error) {
	var borderPoint models.Coordinates
	request :=
		`SELECT x, y, widht, height 
	FROM sector_border_points 
	WHERE entry = $1 
	AND exit = $2;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return nil, err
	}

	err = tx.QueryRow(
		context.Background(),
		request,
		entry,
		exit).Scan(
		&borderPoint.X,
		&borderPoint.Y,
		&borderPoint.Widht,
		&borderPoint.Height)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil, newErr
		}
		r.logger.Error(err)
		return nil, err
	}
	_ = tx.Commit(context.Background())
	return &borderPoint, nil
}
