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

// получаем координаты аудитории по ее номеру.
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
		return models.Coordinates{}, err
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
			return models.Coordinates{}, newErr
		}
		r.logger.Error(err)
		return models.Coordinates{}, err
	}
	_ = tx.Commit(context.Background())
	return position, nil
}

// получаем координаты границ аудитории по ее номеру.
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
		return models.Coordinates{}, err
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
			return models.Coordinates{}, newErr
		}
		r.logger.Error(err)
		return models.Coordinates{}, err
	}
	_ = tx.Commit(context.Background())
	return borderPoint, nil
}

// получаем координаты одной из границ сектора. По значению входа и выхода из него.
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
		return models.Coordinates{}, err
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
			return models.Coordinates{}, newErr
		}
		r.logger.Error(err)
		return models.Coordinates{}, err
	}
	_ = tx.Commit(context.Background())
	return borderPoint, nil
}

func (r *repository) checkBorderAud(coordinates models.Coordinates) (bool, error) {
	r.logger.Infoln("db - check border auditory")
	request :=
		`SELECT x, y, widht, height
	FROM auditorium_position 
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
		coordinates.X+coordinates.Widht,
		coordinates.Y+coordinates.Height)

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
