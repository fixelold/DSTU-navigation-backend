package importantPlaces

import (
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
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

func newRepository(
	client postgresql.Client, 
	logger *logging.Logger,
) Repository {
	return &repository {
		client: client,
		logger: logger,
	}
}

func (r *repository) Create(places models.ImportantPlaces) (models.ImportantPlaces, appError.AppError) {}

func (r *repository) Read(id int) (models.ImportantPlaces, appError.AppError) {}

func (r *repository) Update(oldpPlaces models.ImportantPlaces, newPlaces models.ImportantPlaces) (models.ImportantPlaces, appError.AppError) {}

func (r *repository) Delete(id int) (appError.AppError) {}

func (r *repository) List(numberBuild models.ImportantPlaces) ([]models.ImportantPlaces, appError.AppError) {}