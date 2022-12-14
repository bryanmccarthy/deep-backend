package app

import (
	"net/mail"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

func (h handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Validate email
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid email"})
		return
	}

	var user models.User

	user.Firstname = req.Firstname
	user.Lastname = req.Lastname
	user.Email = req.Email

	encryptedPassword, _ := EncryptPassword(req.Password)
	user.Password = encryptedPassword

	if err := h.DB.Where("Email = ?", req.Email).First(&user).Error; err == nil {
		c.JSON(400, gin.H{"error": "email already exists"})
		return
	}

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
