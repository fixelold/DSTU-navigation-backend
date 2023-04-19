package importantPlaces

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

const (
	file = "db.go"

	createFunction 	= "create"
	readFunction 	= "read"
	updateFunction 	= "update"
	deleteFunction 	= "delete"
	listFunction 	= "list"
)

var (
	txError 	= appError.NewAppError("can't start transaction")
	queryError 	= appError.NewAppError("failed to complite the request")
	scanError  	= appError.NewAppError("can't scan database response")
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(
	client postgresql.Client, 
	logger *logging.Logger,
) Repository {
	return &repository {
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(places models.ImportantPlaces) (models.ImportantPlaces, appError.AppError) {
	var newImportantPlaces models.ImportantPlaces
	req := `INSERT INTO important_places (name, id_auditorium) 
	SELECT $1::varchar(100), $2 
	WHERE NOT EXISTS 
	(SELECT null FROM important_places 
	WHERE (id_auditorium) = ($2)) RETURNING id;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap(fmt.Sprintf("file: %s, function: %s", file, createFunction))
		txError.Err = err
		return models.ImportantPlaces{}, *txError
	}

	err = tx.QueryRow(
		context.Background(),
		req,
		places.Name,
		places.AuditoryID).Scan(&newImportantPlaces.ID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, createFunction))
			queryError.Err = pgErr
			return models.ImportantPlaces{}, *queryError
		}
		queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, createFunction))
		queryError.Err = err
		return models.ImportantPlaces{}, *queryError
	}
	_ = tx.Commit(context.Background())
	return newImportantPlaces, appError.AppError{}	
}

func (r *repository) Read(id int) (models.ImportantPlaces, error) {
	var importantPlaces models.ImportantPlaces
	request :=
	`SELECT *
	FROM important_places 
	WHERE id = $1;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap(fmt.Sprintf("file: %s, function: %s", file, readFunction))
		txError.Err = err
		return models.ImportantPlaces{}, txError.Err
	}

	err = tx.QueryRow(
		context.Background(),
		request,
		id).Scan(
		&importantPlaces.ID,
		&importantPlaces.Name,
		&importantPlaces.AuditoryID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, readFunction))
			queryError.Err = pgErr
			return models.ImportantPlaces{}, queryError.Err
		}
		queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, readFunction))
		queryError.Err = err
		return models.ImportantPlaces{}, queryError.Err
	}
	_ = tx.Commit(context.Background())
	return importantPlaces, nil
}

func (r *repository) Update(oldPlaces models.ImportantPlaces, newPlaces models.ImportantPlaces) (models.ImportantPlaces, appError.AppError) {
	request := `
		UPDATE important_places
		SET name = $1,
		id_auditorium = $2
		WHERE id = $3;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap(fmt.Sprintf("file: %s, function: %s", file, updateFunction))
		txError.Err = err
		return models.ImportantPlaces{}, *txError
	}

	_, err = tx.Exec(context.Background(),
		request,
		newPlaces.Name,
		newPlaces.AuditoryID,
		oldPlaces.ID)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
				pgErr = err.(*pgconn.PgError)
				queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, updateFunction))
				queryError.Err = err
				return models.ImportantPlaces{}, *queryError
		}
		queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, updateFunction))
		queryError.Err = err
		return models.ImportantPlaces{}, *queryError
	}
	_ = tx.Commit(context.Background())
	return newPlaces, appError.AppError{}
}

func (r *repository) Delete(id int) (appError.AppError) {
	request := `
	DELETE FROM important_places
	WHERE id = $1;`

	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap(fmt.Sprintf("file: %s, function: %s", file, deleteFunction))
		txError.Err = err
		return *txError
	}

	_, err = tx.Exec(context.Background(),
		request,
		id)

	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
				pgErr = err.(*pgconn.PgError)
				queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, deleteFunction))
				queryError.Err = err
				return *queryError
		}
		queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, deleteFunction))
		queryError.Err = err
		return *queryError
	}
	_ = tx.Commit(context.Background())
	return appError.AppError{}
}

func (r *repository) List(numberBuild models.ImportantPlaces) ([]models.ImportantPlaces, appError.AppError) {
	var places []models.ImportantPlaces
	req := `SELECT * FROM important_places;`
	tx, err := r.client.Begin(context.Background())
	if err != nil {
		_ = tx.Rollback(context.Background())
		txError.Wrap(fmt.Sprintf("file: %s, function: %s", file, listFunction))
		txError.Err = err
		return nil, *txError
	}

	rows, err := tx.Query(context.Background(), req)
	if err != nil {
		_ = tx.Rollback(context.Background())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, listFunction))
			queryError.Err = pgErr
			return nil, *queryError
		}
		queryError.Wrap(fmt.Sprintf("file: %s, function: %s", file, listFunction))
		queryError.Err = err
		return nil, *queryError
	}

	for rows.Next() {
		var sl models.ImportantPlaces
		err := rows.Scan(&sl.ID, &sl.Name, &sl.AuditoryID)
		if err != nil {
			scanError.Wrap(fmt.Sprintf("file: %s, function: %s", file, listFunction))
			scanError.Err = err
			return nil, *scanError
		}
		places = append(places, sl)
	}

	_ = tx.Commit(context.Background())
	return places, appError.AppError{}
}