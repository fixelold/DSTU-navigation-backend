package user

import (
	"github.com/gin-gonic/gin"

	"navigation/internal/logging"
	"navigation/internal/transport/rest/handlers"
	"navigation/internal/transport/rest/middleware"
)

const userPathURL = "/user"

type handler struct {
	logger         *logging.Logger
	userMiddleware middleware.UserMiddleware
}

func NewHandler(logger *logging.Logger, userMiddleware middleware.UserMiddleware) handlers.Handler {
	return &handler{
		logger:         logger,
		userMiddleware: userMiddleware,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	jwtMiddleware := h.userMiddleware.JwtMiddleware()
	user := router.Group(userPathURL)
	user.Use(middleware.CORSMiddleware())
	user.POST("/signin", jwtMiddleware.LoginHandler)
}
