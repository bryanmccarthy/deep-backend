package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type updateTaskRequest struct {
	ID         uint
	Title      string
	TimeSpent  uint
	Difficulty uint
	Completed  bool
}

func (h handler) updateTask(c *gin.Context) {
	var req updateTaskRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var task models.Task

	if err := h.DB.Where("id = ? AND user_id = ?", req.ID, session.Values["user_id"]).First(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task.Title = req.Title
	task.TimeSpent = req.TimeSpent
	task.Difficulty = req.Difficulty
	task.Completed = req.Completed

	if err := h.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &task)
}
