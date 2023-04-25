package importantPlaces

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"

	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/models"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
)

const (
	places = "/places"
)

type handler struct {
	logger         *logging.Logger
	repository     Repository
	userMiddleware middleware.UserMiddleware
}

func NewHandler(logger *logging.Logger, repository Repository, um middleware.UserMiddleware) handlers.Handler {
	return &handler{
		logger:         logger,
		repository:     repository,
		userMiddleware: um,
	}
}

// структура будет использоваться в методах read, update и delete.
type request struct {
	ID int `form:"id" binding:"required"`
}

func (h *handler) Register(router *gin.RouterGroup) {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode data"})
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
}

func (h *handler) Read(c *gin.Context) {
	var r request
	var err appError.AppError
	var places models.ImportantPlaces
	err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "Read"))

	if err.Err = c.ShouldBindQuery(&r); err.Err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode data"})
	}

	places, err.Err = h.repository.Read(r.ID)
	if err.Err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	c.JSON(http.StatusOK, places)
}

func (h *handler) Update(c *gin.Context) {
	var r request // тут будет хранится id записи, которую надо обновить
	var err appError.AppError
	var oldPlace models.ImportantPlaces // тут будет хранится запись из бд
	var newPlace models.ImportantPlaces
	err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "Update"))

	if err.Err = c.ShouldBindQuery(&r); err.Err != nil {
		// h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode data"})
		return
	}

	if err.Err = c.ShouldBindJSON(&newPlace); err.Err != nil {
		// h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode data"})
		return
	}

	// проверка на существование данных, которые надо обновить
	oldPlace, err.Err = h.repository.Read(r.ID)
	if err.Err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	// проверка на существование новой аудитории
	// _, err.Err = h.repository.Read(newPlace.AuditoryID)
	// if err.Err != nil {
	// 	fmt.Println("data - ", newPlace.AuditoryID)
	// 	fmt.Println("error - ", err.Err)
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
	// 	return
	// }

	newPlace, err.Err = h.repository.Update(oldPlace, newPlace)
	if err.Err != nil && !strings.EqualFold(err.ToString(), pgx.ErrNoRows.Error()) {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// if newPlace.ID == 0 {
	// 	c.JSON(http.StatusConflict, gin.H{"error": "important place already exists"})
	// 	return
	// }

	c.JSON(http.StatusOK, newPlace)
}

func (h *handler) Delete(c *gin.Context) {
	var r request // тут будет хранится id записи, которую надо обновить
	var err appError.AppError
	err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "Delete"))

	if err.Err = c.ShouldBindQuery(&r); err.Err != nil {
		// h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode data"})
		return
	}

	_, err.Err = h.repository.Read(r.ID)
	if err.Err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	err.Err = h.repository.Delete(r.ID)
	if err.Err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

type listRequest struct {
	NumberBiuld int `form:"number_build" binding:"required"`
}

func (h *handler) List(c *gin.Context) {
	var r listRequest
	var places []models.ImportantPlaces
	var err appError.AppError
	err.Wrap(fmt.Sprintf("package: %s, file: %s, function: %s", "importantPlaces", "handler.go", "List"))

	if err.Err = c.ShouldBindQuery(&r); err.Err != nil {
		// h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode data"})
		return
	}

	places, err.Err = h.repository.List(r.NumberBiuld)
	if err.Err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
	}

	c.JSON(http.StatusOK, places)
}
