package main

import (
	"context"
	"navigation/internal/app/auditory"
	"navigation/internal/app/drawPath"
	"navigation/internal/app/pathBuilder"
	"navigation/internal/app/user"
	"navigation/internal/config"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	appContext := context.Background()
	logger := logging.GetLogger()
	router := gin.Default()
	v1Group := router.Group("/api/v1")

	appConfig := config.GetConfig()
	pgConn := postgresql.NewClient(appContext, *appConfig)

	userPrepareRepo := user.NewRepository(pgConn, logger)
	userPrepare := user.NewUser(logger, userPrepareRepo)
	userPrepare.Create()

	pathBuildingRepo := pathBuilder.NewRepository(pgConn, logger)
	pathBuldingController := pathBuilder.NewHandler(logger, pathBuildingRepo)

	drawPathRepo := drawPath.NewRepository(pgConn, logger)
	drawPathController := drawPath.NewHandler(logger, drawPathRepo)

	// users := user.NewRepository(pgConn, logger)
	userController := user.NewHandler(logger, middleware.UserMiddleware{Client: pgConn, Logger: logger})

	auditoryRepo := auditory.NewRepository(pgConn, logger)
	auditoryController := auditory.NewHandler(logger, auditoryRepo, middleware.UserMiddleware{Client: pgConn, Logger: logger})

	pathBuldingController.Register(v1Group)
	drawPathController.Register(v1Group)
	userController.Register(v1Group)
	auditoryController.Register(v1Group)

	logger.Fatalln(router.Run(":8080"))
}
