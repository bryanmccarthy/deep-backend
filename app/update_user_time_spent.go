package app

import (
  "net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
)

type updateUserTimeSpentRequest struct {
  TimeSpent uint `json:"time_spent"`
}

func (h handler) updateUserTimeSpent(c *gin.Context) {
  var req updateUserTimeSpentRequest

  if err := c.BindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

  var user models.User

  if err := h.DB.Where("id = ?", session.Values["user_id"]).First(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  user.TimeSpent += req.TimeSpent

  if err := h.DB.Save(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, &user.TimeSpent)
}

