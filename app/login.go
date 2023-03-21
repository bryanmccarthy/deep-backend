package app

import (
	"net/http"

	"github.com/bryanmccarthy/deep-backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string
	Password string
}

func (h handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := h.DB.Where("Email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	session, err := store.Get(c.Request, "session")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session.Values["authenticated"] = true
	session.Values["user_id"] = user.ID

	if err := session.Save(c.Request, c.Writer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userObj := map[string]string{
		"firstname": user.Firstname,
		"lastname":  user.Lastname,
	}

	c.JSON(http.StatusOK, &userObj)
}
