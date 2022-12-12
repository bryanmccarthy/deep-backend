package users

import (
	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (h handler) Login(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := h.DB.Where("Email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)
	c.JSON(200, &user)
}

// Test with:
// Email: testemail@gmail.com
// Password: testpassword
