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

	// Auth routes
	r.POST("/auth/register", h.Register)
	r.POST("/auth/login", h.Login)
	r.GET("/auth/logout", h.Logout)

	// User routes
	r.GET("/user/time_spent", h.userTimeSpent)
	r.PUT("/user/update/time_spent", h.updateUserTimeSpent)

	// Tasks routes
	r.GET("/tasks", h.tasks)
	r.POST("/tasks/create", h.createTask)
	r.PUT("/tasks/update/title", h.updateTaskTitle)
	r.PUT("/tasks/update/difficulty", h.updateTaskDifficulty)
	r.PUT("/tasks/update/completed", h.updateTaskCompleted)
	r.PUT("/tasks/update/due_date", h.updateTaskDueDate)
	r.PUT("/tasks/update/progress", h.updateTaskProgress)
	r.PUT("/tasks/update/time_spent", h.updateTaskTimeSpent)
	r.DELETE("/tasks/delete", h.deleteTask)

	// Notes routes
	r.GET("/notes/:task_id", h.notes)
	r.POST("/notes/create", h.createNote)
	// r.PUT("/notes/update/title", h.updateNoteTitle)
	r.PUT("/notes/update/content", h.updateNoteContent)
	r.DELETE("/notes/delete", h.deleteNote)
}
