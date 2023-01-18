package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Logout(c *gin.Context) {
	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Session values:")
	fmt.Println(session.Values)

	session.Values["authenticated"] = false

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout"})
}
