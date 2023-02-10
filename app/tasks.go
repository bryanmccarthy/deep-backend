package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

func (h handler) tasks(c *gin.Context) {
	var tasks []models.Task

	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.DB.Where("user_id = ?", session.Values["user_id"]).Order("created_at").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
