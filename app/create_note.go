package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type createNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	TaskID  uint   `json:"task_id"`
}

func (h handler) createNote(c *gin.Context) {
	var req createNoteRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var note models.Note

	note.Title = req.Title
	note.Content = req.Content
	note.TaskID = req.TaskID
	note.UserID = session.Values["user_id"].(uint)

	if err := h.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &note)
}
