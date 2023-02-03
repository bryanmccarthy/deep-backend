package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type updateTaskRequest struct {
	Title string // TODO: add more fields
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

	task.Title = req.Title
	task.UserID = session.Values["user_id"].(uint)

	// update task

}
