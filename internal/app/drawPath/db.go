package drawPath

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

func (r *repository) getAuditoryPosition(number string) (*models.Coordinates, error) {
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

func (r *repository) getBorderPoint(number string) (*models.Coordinates, error) {
	var borderPoint models.Coordinates
	request :=
		`SELECT x, y, widht, height 
	FROM border_points 
	JOIN auditorium 
	ON border_points.id_auditorium = auditorium.id 
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

func (r *repository) checkBorderAud(coordinates models.Coordinates) (bool, error) {
	request :=
	`SELECT x, y, widht, height
	FROM auditorium_position 
	WHERE x <= $1 AND $1 <= (x+widht)
	AND y <= $2 AND $2 <= (y+height)`
	// Возможно тут надо добавить вместо x написать x+widht и т.д с y.

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return false, err
	}

	res, err := tx.Exec(
		context.Background(),
		request,
		coordinates.X + coordinates.Widht,
		coordinates.Y + coordinates.Height)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return false, newErr
		}
		r.logger.Error(err)
		return false, err
	}
	_ = tx.Commit(context.Background())

	if res.RowsAffected() != 0 {
		return false, nil
	}

	// Возможно тут надо это добавить в else.
	return true, nil
}

func (r *repository) checkBorderSector(coordinates models.Coordinates) (bool, error) {
	request :=
	`SELECT x, y, widht, height
	FROM border_points 
	WHERE x <= $1 AND $1 <= (x+widht)
	AND y <= $2 AND $2 <= (y+height)`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return false, err
	}

	res, err := tx.Exec(
		context.Background(),
		request,
		coordinates.X + coordinates.Widht,
		coordinates.Y + coordinates.Height)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return false, newErr
		}
		r.logger.Error(err)
		return false, err
	}
	_ = tx.Commit(context.Background())

	if res.RowsAffected() != 0 {
		return false, nil
	}

	return true, nil
}