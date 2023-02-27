package pathBuilder

import (
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
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
	pathBuilding.Use(middleware.CORSMiddleware)
	pathBuilding.GET("", h.pathBuilding)
}

type auditorys struct {
	Start string `form:"start" binding:"required"`
	End   string `form:"end" binding:"required"`
}

type response struct {
	Sectors []int `json:"sectors"`
}

func (h *handler) pathBuilding(c *gin.Context) {
	var auditorys auditorys
	var response response

	if err := c.ShouldBindQuery(&auditorys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	start, end, err := h.GetSector(auditorys.Start, auditorys.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// TODO сделать обработку ошибки
	response.Sectors, err = h.Builder(start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}
