package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main(){
	r:=gin.Default()

	 // Middleware to measure time
	 r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)
		log.Printf("Gin: %s %s took %s\n", c.Request.Method, c.Request.URL.Path, elapsed)
	   })

	r.GET("/hello", func(c *gin.Context){
		c.JSON(200,gin.H{"message":"Hello, World!"})
	})
	r.Run(":8080")
}
