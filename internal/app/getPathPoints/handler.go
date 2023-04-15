package getPathPoints

import (
	"fmt"
	"navigation/internal/appError"
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	urlPath = "/points" // url путь

	transitionYes = 1 // переход между этажами есть
	transitionNo  = 2 // перехода между этажами нет

	file = "handler.go"
	getPointsFuntion = "getPoints"
	getAuddiencePointsFunction = "getAuddiencePoints"
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
	points := router.Group(urlPath)
	points.Use(middleware.CORSMiddleware)
	points.POST("/points", h.getPoints)
	points.GET("/aud-points", h.getAuddiencePoints)
}

type requestData struct {
	Start   string `json:"start" binding:"required"`
	End     string `json:"end" binding:"required"`
	Sectors []int  `json:"sectors"`

	transition       int `json:"transition" binding:"required"`
	transitionNumber int `json:"transition_number" binding:"required"`
}

func (h *handler) getPoints(c *gin.Context) {
	var err appError.AppError
	var data requestData

	err.Err = c.ShouldBindJSON(&data)
	if err.Err != nil {
		err.Wrap(fmt.Sprintf("file: %s, function: %s", file, getPointsFuntion))
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	p := NewPointsController(data.Start, data.End, data.Sectors, h.logger, h.repository, data.transition, data.transitionNumber)

	response, err := p.controller()
	if err.Err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// type request struct {
// 	Start            string `form:"start" binding:"required"`
// 	End              string `form:"end" binding:"required"`
// 	Transition       int    `form:"transition" binding:"required"`
// 	TransitionNumber int    `form:"transition_number"`
// }

func (h *handler) getAuddiencePoints(c *gin.Context) {
	var request requestData
	var err appError.AppError

	err.Err = c.ShouldBindQuery(&request)
	if err.Err != nil {
		err.Wrap(fmt.Sprintf("file: %s, function: %s", file, getAuddiencePointsFunction))
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	audPoints := NewColoring(request.Start, request.End, h.logger, h.repository, request.transition, request.transitionNumber)
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
