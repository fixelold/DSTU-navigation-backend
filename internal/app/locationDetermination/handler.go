package locationDetermination

import (
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

const pathBuildingURL = "/get-sector"

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

type auditoryNumber struct {
	Number string `form:"number" binding:"required"`
}

func (h *handler) pathBuilding(c *gin.Context) {
	var audNumber auditoryNumber

	if err := c.ShouldBindQuery(&audNumber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	// TODO сделать обработку ошибки
	res, err := h.GetSector(audNumber.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, res)
}
