package getPathPoints

import (
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const drawPathURL = "/draw-path"

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
}

type navigationObject struct {
	Start   string `json:"start" binding:"required"`
	End     string `json:"end" binding:"required"`
	Sectors []int  `json:"sectors" binding:"required"`
}

func (h *handler) getPoints(c *gin.Context) {
	var err error
	var navObj navigationObject

	if err := c.ShouldBindJSON(&navObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	p := NewPointsController(navObj.Start, navObj.End, navObj.Sectors, h.logger, h.repository)

	response, err := p.getPathPoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}
