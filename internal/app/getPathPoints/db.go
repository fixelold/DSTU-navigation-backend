package getPathPoints

import (
	"context"
	"errors"
	"fmt"
	"navigation/internal/appError"
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

var (
	txError    = appError.NewAppError("can't start transaction")
	queryError = appError.NewAppError("failed to complete the request")
)

// получаем координаты аудитории по ее номеру.
func (r *repository) getAudPoints(number string) (models.Coordinates, appError.AppError) {
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
		txError.Wrap("getAudPoints")
		txError.Err = err
		return models.Coordinates{}, *txError
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
			queryError.Wrap("getAudPoints")
			queryError.Err = pgErr
			return models.Coordinates{}, *queryError
		}
		queryError.Wrap("getAudPoints")
		queryError.Err = err
		return models.Coordinates{}, *queryError
	}
	_ = tx.Commit(context.Background())
	return position, appError.AppError{}
}

// получаем координаты границ аудитории по ее номеру.
func (r *repository) getAudBorderPoint(number string) (models.Coordinates, appError.AppError) {
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
		txError.Wrap("getAudBorderPoint")
		txError.Err = err
		return models.Coordinates{}, *txError
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
			queryError.Wrap("getAudBorderPoint")
			queryError.Err = pgErr
			return models.Coordinates{}, *queryError
		}
		queryError.Wrap("getAudBorderPoint")
		queryError.Err = err
		return models.Coordinates{}, *queryError
	}
	_ = tx.Commit(context.Background())
	return borderPoint, appError.AppError{}
}

// получаем координаты одной из границ сектора. По значению входа и выхода из него.
func (r *repository) getSectorBorderPoint(entry, exit int) (models.Coordinates, appError.AppError) {
	var borderPoint models.Coordinates
	request :=
		`SELECT x, y, widht, height 
	FROM sector_border_points 
	WHERE entry = $1 
	AND exit = $2;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap("getSectorBorderPoint")
		txError.Err = err
		return models.Coordinates{}, *txError
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
			queryError.Wrap("getSectorBorderPoint")
			queryError.Err = pgErr
			return models.Coordinates{}, *queryError
		}
		queryError.Wrap("getSectorBorderPoint")
		queryError.Err = err
		return models.Coordinates{}, *queryError
	}
	_ = tx.Commit(context.Background())
	return borderPoint, appError.AppError{}
}

// проверка, чтобы точки пути не находились в границах аудитории.
func (r *repository) checkBorderAud(coordinates models.Coordinates) (bool, appError.AppError) {
	r.logger.Infoln("db - check border auditory")
	request :=
		`SELECT x, y, widht, height
	FROM auditorium_position 
	WHERE x <= $1 AND $1 <= (x+widht)
	AND y <= $2 AND $2 <= (y+height)`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap("checkBorderAud")
		txError.Err = err
		return false, *txError
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
			queryError.Wrap("checkBorderAud")
			queryError.Err = pgErr
			return false, *queryError
		}
		queryError.Wrap("checkBorderAud")
			queryError.Err = err
		return false, *queryError
	}
	_ = tx.Commit(context.Background())

	if res.RowsAffected() != 0 {
		return false, appError.AppError{}
	}

	// Возможно тут надо это добавить в else.
	return true, appError.AppError{}
}

func (r *repository) checkBorderSector(coordinates models.Coordinates) (bool, error) {
	r.logger.Infoln("db - check border sector")
	request :=
		`SELECT x, y, widht, height
	FROM sector_border_points 
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

	return true, nil
}
