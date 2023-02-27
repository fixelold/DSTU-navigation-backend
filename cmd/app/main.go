package main

import (
	"context"
	"navigation/internal/app/drawPath"
	"navigation/internal/app/pathBuilder"
	"navigation/internal/config"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"

	"github.com/gin-gonic/gin"
)

func main() {
	appContext := context.Background()
	logger := logging.GetLogger()
	router := gin.Default()
	v1Group := router.Group("/api/v1")

	appConfig := config.GetConfig()
	pgConn := postgresql.NewClient(appContext, *appConfig)

	pathBuildingRepo := pathBuilder.NewRepository(pgConn, logger)
	pathBuldingController := pathBuilder.NewHandler(logger, pathBuildingRepo)

	drawPathRepo := drawPath.NewRepository(pgConn, logger)
	drawPathController := drawPath.NewHandler(logger, drawPathRepo)

	pathBuldingController.Register(v1Group)
	drawPathController.Register(v1Group)

	logger.Fatalln(router.Run(":8080"))
}
