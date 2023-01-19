package main

import (
	"context"
	"navigation/internal/app/pathBuilder"
	"navigation/internal/config"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
)

func main() {
	appContext := context.Background()
	logger := logging.GetLogger()

	appConfig := config.GetConfig()
	pgConn := postgresql.NewClient(appContext, *appConfig)

	p := pathBuilder.NewPathBuilder(logger, pgConn)
	_, _ = p.Builder(1, 1)

}
