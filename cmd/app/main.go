package main

import (
	"context"
	"fmt"
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
	res, _ := p.Builder(31, 37)
	fmt.Println(res)

}
