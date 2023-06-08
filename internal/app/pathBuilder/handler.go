package pathBuilder

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
)

const getSectors = "/get-sectors"

var (
	bindQueryError = appError.NewAppError("can't decode query")

	noTransition = 1
	stairs   = 2
	elevator = 3
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(logger *logging.Logger, repository Repository) handlers.Handler {
	return &handler {
		logger:     logger,
		repository: repository,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	pathBuilding := router.Group(getSectors)
	pathBuilding.Use(middleware.CORSMiddleware())
	pathBuilding.GET("", h.getSectors)
}

type request struct {
	Start               string `form:"start" binding:"required"`
	End                 string `form:"end" binding:"required"`
	TypeTranstionSector int    `form:"type_transtion_sector" binding:"required"`
}

type response struct {
	Sectors []int `json:"sectors"`
}

// получение секторов, через которые надо пройти пользователю
func (h *handler) getSectors(c *gin.Context) {
	var err appError.AppError
	var request request
	var response response
	var transitionNumber int

	if err := c.ShouldBindQuery(&request); err != nil {
		bindQueryError.Err = err
		h.logger.Error(bindQueryError.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}
 
	start, end, err := h.GetSector(request.Start, request.End, request.TypeTranstionSector)
	if err.Err != nil {
		err.Wrap("handler getSectors")
		h.logger.Error(err.Error())
		fmt.Println("Work sectors: ", start, end)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.ToString()})
		return
	}

	if start == end {
		response.Sectors = append(response.Sectors, start)
		c.JSON(http.StatusOK, response)
		return
	}

	if request.TypeTranstionSector != noTransition {

		if request.TypeTranstionSector == stairs {
			transitionNumber, err = h.stairs(start)
			if err.Err != nil {
				err.Wrap("getSector")
				h.logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
				return
			}
		}

		if request.TypeTranstionSector == elevator {
			transitionNumber, err = h.elevator(start)
			if err.Err != nil {
				h.logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
				return
			}
		}
	}

	// TODO сделать обработку ошибки
	response.Sectors, err = h.Builder(start, end, transitionNumber, request.TypeTranstionSector)
	if err.Err != nil {
		err.Wrap("handler getSectors")
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}
