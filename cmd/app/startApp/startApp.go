package startApp

import (
	"context"
	"navigation/internal/app/auditory"
	"navigation/internal/app/getPathPoints"
	"navigation/internal/app/user"
	"navigation/internal/config"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	connection  postgresql.Client
	logger      *logging.Logger
	router      *gin.Engine
	routerGroup *gin.RouterGroup
}

func (r *Router) StartApp() {
	r.prepareData()
	r.prepareRouter()
	r.logger.Fatalln(r.router.Run(":8080"))
}

func (r *Router) prepareData() {
	appContext := context.Background()
	r.logger = logging.GetLogger()
	r.router = gin.Default()
	r.routerGroup = r.router.Group("/api/v1")
	appConfig := config.GetConfig()
	r.connection = postgresql.NewClient(appContext, *appConfig)
}

func (r *Router) prepareRouter() {
	userPrepareRepo := user.NewRepository(r.connection, r.logger)
	userPrepare := user.NewUser(r.logger, userPrepareRepo)
	userPrepare.Create()

	pathBuildingRepo := pathBuilder.NewRepository(r.connection, r.logger)
	pathBuldingController := pathBuilder.NewHandler(r.logger, pathBuildingRepo)

	drawPathRepo := getPathPoints.NewRepository(r.connection, r.logger)
	drawPathController := getPathPoints.NewHandler(r.logger, drawPathRepo)

	userController := user.NewHandler(r.logger, middleware.UserMiddleware{Client: r.connection, Logger: r.logger})

	auditoryRepo := auditory.NewRepository(r.connection, r.logger)
	auditoryController := auditory.NewHandler(r.logger, auditoryRepo, middleware.UserMiddleware{Client: r.connection, Logger: r.logger})

	pathBuldingController.Register(r.routerGroup)
	drawPathController.Register(r.routerGroup)
	userController.Register(r.routerGroup)
	auditoryController.Register(r.routerGroup)
}
