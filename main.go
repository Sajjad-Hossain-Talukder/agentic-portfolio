package main

import (
	"agentic-portfolio/agent"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func main() {
	cvPath := "./portfolio/docs/cv.json"
	imagesPath := "./portfolio/images"
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Agentic Portfolio API is running 🚀")
	})

	r.GET("/api/v1/profile", func(c *gin.Context) {
		html, err := agent.GeneratePortfolioHTML(cvPath, imagesPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
	})

	// Start server
	r.Run(":8080")
}
