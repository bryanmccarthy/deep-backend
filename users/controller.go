package users

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

var (
	store = sessions.NewCookieStore([]byte("set-this-to-secret")) // TODO: set this to secret
)

func Init() {
	store.Options = &sessions.Options{
		MaxAge:   3600 * 10, // 10 hours
		HttpOnly: true,
	}
}

func Routes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.POST("/users/register", h.Register)
	r.POST("/users/login", h.Login)
	r.GET("/users/logout", h.Logout)
	r.GET("/users/secret", h.Secret)
}
