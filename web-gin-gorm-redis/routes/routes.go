package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/langyan/web-gin-gorm-redis/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users/:id", controllers.GetUser)
}
