package getPathPoints

import (
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const drawPathURL = "/draw-path"

var (
	shouldBindQueryError = appError.NewAppError("can't decode query data")
)

const (
	transitionYes = 1
	transitionNo  = 2
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
	drawPath := router.Group(drawPathURL)
	drawPath.Use(middleware.CORSMiddleware)
	drawPath.POST("/points", h.getPoints)
	drawPath.GET("/aud-points", h.getAuddiencePoints)
}

type navigationObject struct {
	Start   string `json:"start" binding:"required"`
	End     string `json:"end" binding:"required"`
	Sectors []int  `json:"sectors" binding:"required"`

	transition       int `json:"transition" binding:"required"`
	transitionNumber int `json:"transition_number" binding:"required"`
}

func (h *handler) getPoints(c *gin.Context) {
	var err appError.AppError
	var navObj navigationObject

	err.Wrap("getPoints")

	if err := c.ShouldBindJSON(&navObj); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	p := NewPointsController(navObj.Start, navObj.End, navObj.Sectors, h.logger, h.repository, navObj.transition, navObj.transitionNumber)

	response, err := p.getPathPoints()
	if err.Err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

type request struct {
	Start            string `form:"start" binding:"required"`
	End              string `form:"end" binding:"required"`
	Transition       int    `form:"transition" binding:"required"`
	TransitionNumber int    `form:"transition_number"`
}

func (h *handler) getAuddiencePoints(c *gin.Context) {
	var request request
	var err appError.AppError

	err.Err = c.ShouldBindQuery(&request)
	if err.Err != nil {
		shouldBindQueryError.Err = err.Err
		shouldBindQueryError.Wrap("getAuddiencePoints")
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	audPoints := NewColoring(request.Start, request.End, h.logger, h.repository, request.Transition, request.TransitionNumber)
	err = audPoints.GetColoringPoints()
	if err.Err != nil {
		err.Wrap("getAuddiencePoints")
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, audPoints)
	return
}
