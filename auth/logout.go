package auth

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	session.Values["authenticated"] = false

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "logout"})
}
