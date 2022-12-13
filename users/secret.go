package users

import (
	"github.com/gin-gonic/gin"
)

// TODO: Delete this file
func (h handler) Secret(c *gin.Context) {
	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.JSON(401, gin.H{"error": "not authenticated"})
		return
	}

	c.JSON(200, gin.H{"message": "session authenticated"})
}
