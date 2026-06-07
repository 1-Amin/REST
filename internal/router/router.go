package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yourorg/tasks-api/internal/handlers"
	"github.com/yourorg/tasks-api/internal/middleware"
)

func New(t *handlers.TaskHandler, a *handlers.AuthHandler, secret []byte) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://app.example.com"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		auth.POST("/register", a.Register)
		auth.POST("/login", a.Login)

		secured := v1.Group("")
		secured.Use(middleware.JWT(secret))
		secured.GET("/tasks", t.List)
		secured.POST("/tasks", t.Create)
		secured.GET("/tasks/:id", t.Get)
		secured.PATCH("/tasks/:id", t.Update)
		secured.DELETE("/tasks/:id", t.Delete)
	}
	return r
}
