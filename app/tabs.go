package app

import (
	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTabs(c *gin.Context) {
	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if session.Values["user_id"] == nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var userId uint = session.Values["user_id"].(uint)
	var tabs []models.Tab

	if err := h.DB.Where("user_id = ?", userId).Find(&tabs).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &tabs)
}
