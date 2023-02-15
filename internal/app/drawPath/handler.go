package drawPath

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
	drawPath.GET("", h.getPoints)
}

type navigationObject struct {
	Start   string `form:"start" binding:"required"`
	End     string `form:"end" binding:"required"`
	Sectors []int  `form:"sectors" binding:"requred"`
}

type response struct {
	Points [][]int `json:"points"`
}

func (h *handler) getPoints(c *gin.Context) {
	var err error
	var navObj navigationObject
	var response response

	if err := c.ShouldBindQuery(&navObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	response.Points, err = h.drawPath(navObj.Start, navObj.End, navObj.Sectors)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}
