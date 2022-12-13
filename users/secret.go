package users

import (
	"github.com/gin-gonic/gin"
)

// TODO: Delete this file
func (h handler) Secret(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.JSON(401, gin.H{"error": "not authenticated"})
		return
	}

	c.JSON(200, gin.H{"message": "session authenticated"})
}
