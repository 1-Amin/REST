package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"tasks-api/internal/config"
	"tasks-api/internal/models"
)

func main() {
	_ = godotenv.Load()

	db, err := config.NewDB(os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal("db connect: ", err)
	}
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("migrate: ", err)
	}
	log.Println("connected and migrated")

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "time": time.Now().UTC().Format(time.RFC3339)})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
