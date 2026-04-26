package main

import (
	"news-api/adapter/input/routes"
	"news-api/configuration/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("Application starting...")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":8000"); err != nil {
		logger.Error("Failed to start server: ", err)
		return
	}

	logger.Info("Server started on port 8000")
}
