package sectorToSector

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"

	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
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
	execError  = appError.NewAppError("exec request error")
)

// проверка, чтобы точки пути не находились в границах аудитории.
func (r *repository) checkBorderAud(coordinates models.Coordinates) (bool, appError.AppError) {
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
			execError.Wrap("checkBorderAud")
			execError.Err = pgErr
			return false, *execError
		}
		execError.Wrap("checkBorderAud")
		execError.Err = err
		return false, *execError
	}
	_ = tx.Commit(context.Background())

	if res.RowsAffected() != 0 {
		return false, appError.AppError{}
	}

	// Возможно тут надо это добавить в else.
	return true, appError.AppError{}
}

// проверка, чтобы точки пути не находились в границах сектора.
func (r *repository) checkBorderSector(coordinates models.Coordinates) (bool, appError.AppError) {

	request :=
		`SELECT x, y, widht, height
	FROM sector_border_points 
	WHERE x <= $1 AND $1 <= (x+widht)
	AND y <= $2 AND $2 <= (y+height)`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap("checkBorderSector")
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
			execError.Wrap("checkBorderSector")
			execError.Err = pgErr
			return false, *execError
		}
		execError.Wrap("checkBorderSector")
		execError.Err = err
		return false, *execError
	}
	_ = tx.Commit(context.Background())

	if res.RowsAffected() != 0 {
		return false, appError.AppError{}
	}

	return true, appError.AppError{}
}