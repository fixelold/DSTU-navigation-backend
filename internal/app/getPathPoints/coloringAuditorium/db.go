package getPathPoints

import (
	"context"
	"errors"
	"fmt"

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

// получаем координаты переходного сеткора по его номеру.
func (r *repository) getTransitionPoints(number int) (models.Coordinates, appError.AppError) {
	var position models.Coordinates
	var dbErr appError.AppError
	fmt.Println("number transition: ", number)
	request :=
		`SELECT x, y, widht, height 
	FROM transition_position 
	JOIN transition 
	ON transition_position.id_transition = transition.id 
	WHERE transition.number = $1;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		dbErr.Wrap("getTransitionPoints")
		dbErr = *txError
		dbErr.Err = err
		return models.Coordinates{}, dbErr
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
			dbErr = *queryError
			dbErr.Wrap("getTransitionPoints")
			dbErr.Err = pgErr
			return models.Coordinates{}, dbErr
		}
		dbErr.Wrap("getTransitionPoints")
		dbErr.Err = err
		return models.Coordinates{}, dbErr
	}
	_ = tx.Commit(context.Background())
	return position, appError.AppError{}
}