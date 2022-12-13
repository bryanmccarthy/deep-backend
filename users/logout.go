package users

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Logout(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")

	session.Values["authenticated"] = false
	session.Save(c.Request, c.Writer)
	c.JSON(200, gin.H{"message": "logout"})
}
