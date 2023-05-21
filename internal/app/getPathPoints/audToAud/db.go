package audToAud

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"

	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/models"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(client postgresql.Client) Repository {
	return &repository{
		client: client,
	}
}

var (
	txError    = appError.NewAppError("can't start transaction")
	queryError = appError.NewAppError("failed to complete the request")
	execError  = appError.NewAppError("exec request error")
)

// проверка, чтобы точки пути не находились в границах аудитории.
func (r *repository) checkBorderAud(coordinates models.Coordinates, audNumber string) (bool, appError.AppError) {
	request :=
		`SELECT x, y, widht, height
	FROM auditorium_position 
	WHERE x <= $1 AND $1 <= (x+widht)
	AND y <= $2 AND $2 <= (y+height) 
	AND id_auditorium = (SELECT id FROM auditorium WHERE number = $3)`

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
		coordinates.Y+coordinates.Height,
		audNumber)

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

func (r *repository) checkBorderAud2(coordinates models.Coordinates, sectorNumber int) (bool, appError.AppError) {
	//TODO: тут бы немного подправить. 
	//TODO: Т.к неособо уверен, что $1 и $2 правильно расчитываются. 
	//TODO: И работают для всех случаев.
	request := ``
	if coordinates.Widht < 0 && coordinates.Height > 0 {
		temp := coordinates
		coordinates.X = temp.X + temp.Widht
		coordinates.Y = temp.Y
		coordinates.Widht = temp.X - (temp.X + temp.Widht)
		coordinates.Height = temp.Height

		request =
			`SELECT x, y, widht, height
			FROM auditorium_position
			JOIN auditorium
			ON auditorium.id = auditorium_position.id_auditorium
			JOIN sector
			ON sector.id = auditorium.id_sector
			WHERE ($1 >= x AND (x+widht) <= $2
			AND $3 >= y AND $4 <= (y+height)
			AND sector.number = $5)
			OR (
				$1 <= x AND (x+widht) <= $2 
				AND $3 >= y AND $4 <= (y+height) 
				AND sector.number = $5);`
	} else {
		request =
			`SELECT x, y, widht, height
			FROM auditorium_position
			JOIN auditorium
			ON auditorium.id = auditorium_position.id_auditorium
			JOIN sector
			ON sector.id = auditorium.id_sector
			WHERE ($1 <= x AND x <= $2
			AND $3 >= y AND $4 <= (y+height)
			AND sector.number = $5)
			OR (
				$1 <= x AND (x+widht) <= $2 
				AND $3 >= y AND $4 <= (y+height) 
				AND sector.number = $5);`
	}

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
		coordinates.X,
		coordinates.X + coordinates.Widht,
		coordinates.Y,
		coordinates.Y + coordinates.Height,
		sectorNumber)

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