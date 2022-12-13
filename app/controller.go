package app

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

	// TODO: change status codes
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	r.GET("/auth/logout", h.Logout)
}