package main

import (
	"context"
	"fmt"
	"navigation/internal/app/locationDetermination"
	"navigation/internal/config"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
)

func main() {
	appContext := context.Background()
	logger := logging.GetLogger()

	appConfig := config.GetConfig()
	pgConn := postgresql.NewClient(appContext, *appConfig)

	l := locationDetermination.NewLocation("1-367а", logger, pgConn)
	sector, _ := l.GetSector()
	fmt.Println("sector - ", sector)
}
