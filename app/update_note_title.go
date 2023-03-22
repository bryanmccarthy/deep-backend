package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type updateNoteTitleRequest struct {
	ID      uint   `json:"id"`
	Title string `json:"title"`
}

func (h handler) updateNoteTitle (c *gin.Context) {
	var req updateNoteTitleRequest

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

	note.Title = req.Title

	if err := h.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &note)
}