package handler

import (
	"apiInterview/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	srvc *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		srvc: service,
	}
}

func (handler *Handler) Init() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	user := router.Group("/users")
	{
		user.POST("/", handler.CreateUser)
		user.PUT("/:username", handler.UpdateUser)
		user.DELETE("/:username", handler.DeleteUser)
		user.GET("/:username", handler.GetUser)
		user.GET("/", handler.GetAllUsers)
	}
	return router
}
