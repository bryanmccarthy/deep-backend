package users

import (
	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

func (h handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user.Username = req.Username
	user.Email = req.Email

	encryptedPassword, _ := EncryptPassword(req.Password)
	user.Password = encryptedPassword

	if err := h.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, &user)
}

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}
