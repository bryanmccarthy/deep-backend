package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type createTaskRequest struct {
	Title      string `json:"title"`
	Difficulty uint   `json:"difficulty"`
	DueDate    string `json:"due_date"`
	Completed  bool   `json:"completed"`
	Progress   uint   `json:"progress"`
	TimeSpent  uint   `json:"time_spent"`
}

func (h handler) createTask(c *gin.Context) {
	var req createTaskRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var task models.Task

	task.Title = req.Title
	task.Difficulty = req.Difficulty
	task.DueDate = req.DueDate
	task.Completed = req.Completed
	task.Progress = req.Progress
	task.TimeSpent = req.TimeSpent
	task.UserID = session.Values["user_id"].(uint)

	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &task)
}
