package handler

import (
	"github.com/TrackTasks/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		logrus.Error("get task failed: %v", err, http.StatusBadRequest)
		return
	}
	if h.services == nil {
		logrus.Error("services == nil")
	}
	id, err := h.services.CreateTask(newTask)

	if err != nil {
		logrus.Fatalf("create task failed: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
func (h *Handler) GetTask(c *gin.Context) {
	id := c.Query("id")
	taskIdInt, err := strconv.Atoi(id)
	status := c.Query("status")
	//taskId := c.Param("id")
	//taskIdInt, err := strconv.Atoi(taskId)
	//if err != nil {
	//	logrus.Error("convert taskId to int failed: %v", err)
	//}
	//status := c.Param("status")

	task, err := h.services.GetTask(taskIdInt, status)
	if err != nil {
		logrus.Errorf("get task failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get tasks",
		})
		return
	}
	c.JSON(http.StatusOK, task)

}
func (h *Handler) UpdateTask(c *gin.Context) {
	id, _ := c.Params.Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("Error to get id:%v", err, http.StatusBadRequest)
		return
	}
	var updateTask models.Task
	if err := c.BindJSON(&updateTask); err != nil {
		logrus.Error("update task failed: %v", err)
	}
	err = h.services.UpdateTask(intId, updateTask)
	if err != nil {
		logrus.Errorf("update task failed: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{"success update": true})
}
func (h *Handler) DeleteTask(c *gin.Context) {
	id, _ := c.Params.Get("id")
	taskIdInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("convert taskId to int failed: %v", err)
	}
	status := c.Param("status")

	affectedRows, err := h.services.DeleteTask(taskIdInt, status)
	if err != nil {
		logrus.Error("delete task failed: %v", err, http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{
		"affected rows": affectedRows,
	})
}
