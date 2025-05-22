package routes

import (
	"example.com/rental/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/v1/createUser", handlers.CreateUser) 
}
