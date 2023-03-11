package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type updateTaskProgressRequest struct {
	TaskID   uint `json:"task_id"`
	Progress uint `json:"progress"`
}

func (h handler) updateTaskProgress(c *gin.Context) {
	var req updateTaskProgressRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var task models.Task

	if err := h.DB.First(&task, req.TaskID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task.Progress = req.Progress

	if err := h.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &task)
}
