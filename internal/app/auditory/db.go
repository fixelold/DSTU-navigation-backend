package auditory

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"

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

func (r *repository) Read(number string) (auditory, error) {
	var aud auditory
	req :=
		`SELECT * FROM auditorium 
		WHERE number = $1`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return auditory{}, err
	}

	err = tx.QueryRow(
		context.Background(),
		req,
		number).Scan(&aud.ID, &aud.Number, &aud.AuditoryID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return auditory{}, newErr
		}
		r.logger.Error(err)
		return auditory{}, err
	}
	_ = tx.Commit(context.Background())
	return aud, nil
}

func (r *repository) GetDescription(number string) (models.AuditoryDescription, error) {
	var audDEscription models.AuditoryDescription
	req :=
		`SELECT * FROM auditory_description 
		WHERE id_auditory = (SELECT id FROM auditorium WHERE number = $1)`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return models.AuditoryDescription{}, err
	}

	err = tx.QueryRow(
		context.Background(),
		req,
		number).Scan(&audDEscription.ID, &audDEscription.AuditoryID, &audDEscription.Description)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return models.AuditoryDescription{}, newErr
		}
		r.logger.Error(err)
		return models.AuditoryDescription{}, err
	}
	_ = tx.Commit(context.Background())
	return audDEscription, nil
}

func (r *repository) Update(description, number string) error {
	request := `
	UPDATE auditory_description 
	SET description = $1 
	WHERE id_auditory = (SELECT id FROM auditorium WHERE number = $2)
	`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error()) // Прочитать про Tracef
		return err
	}

	_, err = tx.Exec(context.Background(),
		request,
		description,
		number)
	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(
				"SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
				pgErr.SQLState(),
			)
			r.logger.Error(newErr)
			return newErr
		}
		r.logger.Error(err)
		return err
	}
	_ = tx.Commit(context.Background())
	return nil
}
