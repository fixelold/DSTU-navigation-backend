package getPathPoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	getPathPoints "navigation/internal/app/getPathPoints/coloringAuditorium"
	"navigation/internal/appError"
	"navigation/internal/database/client/postgresql"
	"navigation/internal/logging"
	"navigation/internal/models"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
)

const (
	urlPath = "/points" // url путь

	// transitionYes = 1 // переход между этажами есть
	// transitionNo  = 2 // перехода между этажами нет
	// transitionToAud = 3

	noTransition = 1
	stair   = 2
	elevator = 3
	transitionToAud = 4
	aud2Aud = 5

	file                       = "handler.go"
	getPointsFuntion           = "getPoints"
	getAuddiencePointsFunction = "getAuddiencePoints"
)

type handler struct {
	logger     *logging.Logger
	client postgresql.Client
}

func NewHandler(logger *logging.Logger, client postgresql.Client) handlers.Handler {
	return &handler{
		logger:     logger,
		client: client,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	points := router.Group(urlPath)
	points.Use(middleware.CORSMiddleware())
	points.POST("/points", h.getPoints)
	points.GET("/aud-points", h.getAuddiencePoints)
}

type requestData struct {
	Start   string `json:"start"`
	End     string `json:"end"`
	Sectors []int  `json:"sectors"`

	Transition       int `json:"transition"`
	TransitionNumber int `json:"transition_number"`
}

func (h *handler) getPoints(c *gin.Context) {
	var err appError.AppError
	var data requestData
	var response []models.Coordinates

	err.Err = c.ShouldBindJSON(&data)
	if err.Err != nil {
		err.Wrap(fmt.Sprintf("file: %s, function: %s", file, getPointsFuntion))
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	p := NewPointsController(data.Start, data.End, data.Sectors, h.logger, h.client, data.Transition, data.TransitionNumber)
	if data.Transition != noTransition {
		response, err = p.transitionController()
		if err.Err != nil {
			h.logger.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}
	} else {
		if len(data.Sectors) == 1 {
			response, err = p.aud2Aud()
			if err.Err != nil {
				h.logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
				return
			}
		} else {
			response, err = p.controller()
			if err.Err != nil {
				h.logger.Error(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
				return
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

type request struct {
	Start            string `form:"start" binding:"required"`
	End              string `form:"end" binding:"required"`
	Transition       int    `form:"transition" binding:"required"`
	TransitionNumber int    `form:"transition_number"`
}

type response struct {
	Start models.Coordinates `json:"start"`
	End models.Coordinates `json:"end"`
}

func (h *handler) getAuddiencePoints(c *gin.Context) {
	var request request
	var err appError.AppError
	var response response

	err.Err = c.ShouldBindQuery(&request)
	if err.Err != nil {
		err.Wrap(fmt.Sprintf("file: %s, function: %s", file, getAuddiencePointsFunction))
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	audPoints := getPathPoints.NewColoring(request.Start, request.End, h.logger, h.client, request.Transition, request.TransitionNumber)
	err = audPoints.GetColoringPoints()
	if err.Err != nil {
		err.Wrap("getAuddiencePoints")
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	response.Start = audPoints.StartAuditoryPoints
	response.End = audPoints.EndAuditoryPoints
	c.JSON(http.StatusOK, response)
}
