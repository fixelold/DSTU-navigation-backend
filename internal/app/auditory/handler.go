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

type readRequest struct {
	Number string `form:"number" binding:"required"`
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
	ID         int    `json:"id"`
	Number     string `json:"number"`
	AuditoryID int    `json:"auditory_id"`
}

func (h *handler) read(c *gin.Context) {
	var r readRequest
	var auditory auditory
	var err error

	// Связывание данных, полученных от клиентской стороны, со структурой readRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	// Получение аудитории по ее номеру
	auditory, err = h.repository.Read(r.Number)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// Возврат ответа на клиентскую сторону
	c.JSON(http.StatusOK, auditory)
}

// Получение описания начальной и конечной аудитории.
func (h *handler) getDescription(c *gin.Context) {
	var auditorys audNumber
	var response response
	var err error

	// Связывание данных, полученных от клиентской стороны, со структурой audNumber
	if err := c.ShouldBindQuery(&auditorys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode query"})
		return
	}

	// Получение описание начальной аудитории
	response.Start, err = h.repository.GetDescription(auditorys.Start)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// Получение описание конечной адуитории
	response.End, err = h.repository.GetDescription(auditorys.End)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// Возврат ответа на клиентскую сторону
	c.JSON(http.StatusOK, response)
}

// Обновление описания аудитории
func (h *handler) update(c *gin.Context) {
	var err error
	var audDescript updateDescription

	// Связывание данных, полученных от клиентской стороны, со структурой updateDescription
	if err := c.ShouldBindJSON(&audDescript); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't decode json"})
		return
	}

	// Проверка на существования аудитории
	_, err = h.repository.GetDescription(audDescript.AuditoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Обновление описания аудитории по ее id
	err = h.repository.Update(audDescript.Description, audDescript.AuditoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	// Возврат ответа на клиентскую сторону
	c.JSON(http.StatusOK, gin.H{"seccuess": true})
}