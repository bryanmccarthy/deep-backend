package users

import (
	"github.com/gin-gonic/gin"
)

// TODO
func (h handler) Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "logout"})
}
