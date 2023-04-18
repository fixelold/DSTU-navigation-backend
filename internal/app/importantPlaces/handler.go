package importantPlaces

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"

	"github.com/gin-gonic/gin"
)

const (
	places = "/places"
)

var (
	bindError = appError.NewAppError("can't decode data")
)

type handler struct {
	logger *logging.Logger
	repository Repository
	userMiddleware middleware.UserMiddleware
}

func NewHandler(logger *logging.Logger, repository Repository, um middleware.UserMiddleware) handlers.Handler {
	return &handler {
		logger: logger,
		repository: repository,
		userMiddleware: um,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	// userMiddleware middleware.UserMiddleware
	jwtMiddleware := h.userMiddleware.JwtMiddleware()
	places := router.Group(places)
	places.Use(middleware.CORSMiddleware())
	places.GET("", h.Read)
	places.GET("/list", h.List)
	places.Use(jwtMiddleware.MiddlewareFunc())
	{
		places.POST("/create", h.Create)
		places.PUT("/update", h.Update)
		places.DELETE("/delete", h.Delete)
	}
}

func (h *handler) Create(c *gin.Context) {}

func (h *handler) Read(c *gin.Context) {}

func (h *handler) Update(c *gin.Context) {}

func (h *handler) Delete(c *gin.Context) {}

func (h *handler) List(c *gin.Context) {}