package main

import (
	"github.com/gin-gonic/gin"
	"github.com/langyan/web-gin-gorm-redis/cache"
	"github.com/langyan/web-gin-gorm-redis/config"
	"github.com/langyan/web-gin-gorm-redis/routes"
)

func main() {
	// Load configurations
	config.LoadConfig()

	// Connect to database
	config.ConnectDatabase()

	// Initialize Redis
	cache.InitRedis()

	// Set up router
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
