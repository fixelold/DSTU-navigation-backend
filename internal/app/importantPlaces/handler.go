package importantPlaces

import (
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
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

func (h *handler) Create(c *gin.Context) {
	var err appError.AppError
	var places models.ImportantPlaces

	err.Err = c.ShouldBindJSON(&places) 
	if err.Err != nil {
		err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "Create"))
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	response, err := h.repository.Create(places)
	if err.Err != nil && !strings.EqualFold(err.ToString(), pgx.ErrNoRows.Error()) {
		err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "Create"))
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if response.ID == 0 {
		// err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "Create"))
		// h.logger.Error(err.Error())
		c.JSON(http.StatusConflict, gin.H{"error": "important place already exists"})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func (h *handler) Read(c *gin.Context) {}

func (h *handler) Update(c *gin.Context) {
	//проверка на существования старой записи
	//проверка на то, что новые данные не конфликтуют
	//обновление данных
}

func (h *handler) Delete(c *gin.Context) {
	//проверка записи на существования
}

func (h *handler) List(c *gin.Context) {}