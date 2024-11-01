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
			//users.GET("/", h.getAllPeople)
			users.GET("/", h.getWithFilters)

			//tasks := users.Group(":id/tasks/:id")
			//{
			//	tasks.POST("/:start", h.startTask)
			//	tasks.POST("/stop", h.stopTask)
			//
			//}
		}
		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.CreateTask)
			tasks.DELETE("/", h.DeleteTask)
			tasks.PUT("/:id", h.UpdateTask)
			tasks.GET("/", h.GetTask)
		}
	}

	return router
}
