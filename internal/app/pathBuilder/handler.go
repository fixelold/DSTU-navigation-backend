package pathBuilder

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const getSectors = "/get-sectors"

var (
	bindQueryError = appError.NewAppError("can't decode query")

	stairs   = 1
	elevator = 2
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(logger *logging.Logger, repository Repository) handlers.Handler {
	return &handler{
		logger:     logger,
		repository: repository,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	pathBuilding := router.Group(getSectors)
	pathBuilding.Use(middleware.CORSMiddleware)
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

func (h *handler) getSectors(c *gin.Context) {
	var err appError.AppError
	var request request
	var response response
	var transitionSector int

	if err := c.ShouldBindQuery(&request); err != nil {
		bindQueryError.Err = err
		h.logger.Error(bindQueryError.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	start, end, err := h.GetSector(request.Start, request.End)
	if err.Err != nil {
		err.Wrap("handler getSectors")
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if request.TypeTranstionSector == stairs {
		transitionSector, err = h.stairs(start)
		if err.Err != nil {
			err.Wrap("getSector")
			h.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}
	}

	// TODO сделать обработку ошибки
	response.Sectors, err = h.Builder(start, end, transitionSector)
	if err.Err != nil {
		err.Wrap("handler getSectors")
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}
