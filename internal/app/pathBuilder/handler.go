package pathBuilder

import (
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const pathBuildingURL = "/path-building"

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
	pathBuilding := router.Group(pathBuildingURL)
	pathBuilding.GET("", h.pathBuilding)
}

type auditorys struct {
	Start int `form:"start" binding:"required"`
	End   int `form:"end" binding:"required"`
}

func (h *handler) pathBuilding(c *gin.Context) {
	var auditorys auditorys

	if err := c.ShouldBindQuery(&auditorys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	// TODO сделать обработку ошибки
	res, err := h.Builder(auditorys.Start, auditorys.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
