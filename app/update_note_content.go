package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type updateNoteContentRequest struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func (h handler) updateNoteContent(c *gin.Context) {
	var req updateNoteContentRequest

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

	if err := h.DB.Where("id = ? AND user_id = ?", req.ID, session.Values["user_id"]).First(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	note.Content = req.Content

	if err := h.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &note)
}
