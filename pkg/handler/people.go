package handler

import (
	"github.com/TrackTasks/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createPeople(c *gin.Context) {

	var newPeople models.People
	if err := c.BindJSON(&newPeople); err != nil {
		logrus.Error("get people failed: %v", err, http.StatusBadRequest)
		return
	}
	if h.services == nil {
		logrus.Error("services == nil")
	}
	id, err := h.services.Create(newPeople)

	if err != nil {
		logrus.Fatalf("create people failed: %v", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) updatePeople(c *gin.Context) {
	id, _ := c.Params.Get("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("Error to get id:%v", err, http.StatusBadRequest)
		return
	}
	var updatePeople models.People
	if err := c.BindJSON(&updatePeople); err != nil {
		logrus.Error("get people failed: %v", err, http.StatusBadRequest)
		return
	}
	err = h.services.Update(intId, updatePeople)
	if err != nil {
		logrus.Error("update people failed: %v", err, http.StatusBadRequest)
		return
	}
	logrus.Printf("success update people %v", updatePeople)

}
func (h *Handler) deletePeople(c *gin.Context) {
	userId, _ := c.Params.Get("id")
	intId, err := strconv.Atoi(userId)
	if err != nil {
		logrus.Error("Error to get id:%v", err, http.StatusBadRequest)
	}
	affectedRows, err := h.services.People.Delete(intId)
	if err != nil {
		logrus.Error("delete people failed: %v", err, http.StatusBadRequest)
	}
	c.JSON(http.StatusOK, gin.H{
		"affected rows": affectedRows,
	})
}
func (h *Handler) getAllPeople(c *gin.Context) {
	//userId := c.Param("id")
	//id, err := strconv.Atoi(userId)
	//if err != nil {
	//	logrus.Error("Error to get id:%v", err, http.StatusBadRequest)
	//	return
	//}
	people, err := h.services.GetAll()
	if err != nil {
		logrus.Fatalf("get people failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get all people",
		})
	}
	c.JSON(http.StatusOK, people)
}
func (h *Handler) getWithFilters(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	name := c.Query("name")
	surname := c.Query("surname")
	address := c.Query("address")
	passportNumber := c.Query("passport_number")
	passportNumberInt, err := strconv.Atoi(passportNumber)

	people, err := h.services.GetWithFilters(idInt, name, surname, address, passportNumberInt)
	if err != nil {
		logrus.Error("get people failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get people"})
		return
	}
	c.JSON(http.StatusOK, people)

}
