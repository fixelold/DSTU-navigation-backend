package auditory

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"navigation/internal/logging"
	"navigation/internal/models"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
)

const auditoryPathURL = "/auditory"

type handler struct {
	logger         *logging.Logger
	repository     Repository
	userMiddleware middleware.UserMiddleware
}

func NewHandler(logger *logging.Logger, repository Repository, userMiddleware middleware.UserMiddleware) handlers.Handler {
	return &handler{
		logger:         logger,
		repository:     repository,
		userMiddleware: userMiddleware,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	jwtMiddleware := h.userMiddleware.JwtMiddleware()
	auditory := router.Group(auditoryPathURL)
	auditory.Use(middleware.CORSMiddleware())
	auditory.GET("/description", h.getDescription)
	auditory.GET("", h.read)
	auditory.Use(jwtMiddleware.MiddlewareFunc())
	{
		auditory.POST("/update", h.update)
	}
}

type audNumber struct {
	Start string `form:"start" binding:"required"`
	End   string `form:"end" binding:"required"`
}

type response struct {
	Start models.AuditoryDescription `json:"start"`
	End   models.AuditoryDescription `json:"end"`
}

type updateDescription struct {
	AuditoryID  string `json:"auditory_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	JWTToken    string `json:"token"`
}

type auditory struct {
	ID int `json:"id"`
	Number int `json:"number"`
	AuditoryID int `json:"auditory_id"`
}

func (h *handler) read(c *gin.Context) {
	var auditorys audNumber
	var auditory auditory
	var err error

	if err := c.ShouldBindQuery(&auditorys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	auditory, err = h.repository.Read(auditorys.Start)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, auditory)
}

func (h *handler) getDescription(c *gin.Context) {
	var auditorys audNumber
	var response response
	var err error

	if err := c.ShouldBindQuery(&auditorys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	response.Start, err = h.repository.GetDescription(auditorys.Start)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	response.End, err = h.repository.GetDescription(auditorys.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *handler) update(c *gin.Context) {
	var err error
	var audDescript updateDescription

	if err := c.ShouldBindJSON(&audDescript); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	if audDescript.JWTToken == "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		return
	}

	err = h.repository.Update(audDescript.Description, audDescript.AuditoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"seccuess": true})
}
