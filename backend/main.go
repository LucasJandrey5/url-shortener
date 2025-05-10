package main

import (
	"log"
	"os"

	"github.com/LucasJandrey5/url-shortener/backend/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.New()
	routes.InitRoutes(&router.RouterGroup)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
