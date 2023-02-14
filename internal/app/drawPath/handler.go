package drawPath

import (
	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"

	"github.com/gin-gonic/gin"
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

}
