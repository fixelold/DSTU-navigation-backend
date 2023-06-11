package user

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

func (r *repository) Create(user models.User) (models.User, error) {
	request := `
		INSERT INTO users(login, password) 
		VALUES ($1, $2) 
		RETURNING id
		`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error()) // Прочитать про Tracef
		return models.User{}, err
	}

	err = tx.QueryRow(
		context.Background(),
		request,
		user.Login,
		user.Password).Scan(&user.ID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return models.User{}, newErr
		}
		r.logger.Error(err)
		return models.User{}, err
	}
	_ = tx.Commit(context.Background())
	return user, nil
}

func (r *repository) Update(newData models.User) error {
	fmt.Println("Work: ", newData.ID, newData.Login, newData.Password)
	request := `
	UPDATE users 
	SET login = $1, password = $2
	WHERE id = $3
	`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error()) // Прочитать про Tracef
		return err
	}

	_, err = tx.Exec(context.Background(),
		request,
		newData.Login,
		newData.Password,
		newData.ID)
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

func (r *repository) FindRoot() (models.User, error) {
	var user models.User
	req :=
		`SELECT id FROM users WHERE login = 'root';`

	tx, err := r.client.Begin(context.Background())
	r.logger.Infoln("tx - ", tx)
	if err != nil {
		_ = tx.Rollback(context.Background())
		r.logger.Tracef("can't start transaction: %s", err.Error())
		return models.User{}, err
	}

	rows, err := tx.Query(context.Background(), req)
	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return models.User{}, newErr
		}
		r.logger.Error(err)
		return models.User{}, err
	}

	for rows.Next() {
		var sl models.User
		err := rows.Scan(&sl.ID)
		if err != nil {
			r.logger.Errorf("getSectorLink function. Scan error: %s", err.Error())
			return models.User{}, err
		}
		user = sl
	}

	_ = tx.Commit(context.Background())
	return user, nil
}