package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware to measure time
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)
		log.Printf("Gin: %s %s took %s\n", c.Request.Method, c.Request.URL.Path, elapsed)
	})

	//r.GET("/hello", func(c *gin.Context){
	//	c.JSON(200,gin.H{"message":"Hello, World!"})
	//})

	//http://localhost:8080/hello/michael
	// Define a route with a path parameter
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello, " + name + "!",
		})
	})
	//http://localhost:8080/search?q=michael
	// Define a route that accepts query parameters
	r.GET("/search", func(c *gin.Context) {
		query := c.DefaultQuery("q", "default query")
		c.JSON(200, gin.H{
			"query": query,
		})
	})

	r.Run(":8080")
}
