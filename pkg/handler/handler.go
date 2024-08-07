package handler

import (
	"github.com/TrackTasks/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.createPeople)
			users.DELETE("/:id", h.deletePeople)
			users.PUT("/:id", h.updatePeople)
			users.GET("/", h.getPeople)
			tasks := users.Group(":id/tasks/:id")
			{
				tasks.POST("/:start", h.startTask)
				tasks.POST("/stop", h.stopTask)

			}
		}
	}

	return router
}
