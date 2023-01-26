package locationDetermination

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

func (r *repository) GetSector(id string, building uint) (uint, error) {
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

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return 0, err
	}

	err = tx.QueryRow(
		context.Background(),
		req,
		id,
		building).Scan(&sector.Number)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return 0, newErr
		}
		r.logger.Error(err)
		return 0, err
	}
	_ = tx.Commit(context.Background())
	return sector.Number, nil
}
