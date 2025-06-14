package routes

import (
	"example.com/rental/handlers/userHandlers"
	"example.com/rental/handlers/vehicle"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/v1/createVehicle", vehicle.CreateVehicle)
	server.POST("/api/v1/createUser", userHandlers.CreateUser)
	server.GET("/api/v1/profile", userHandlers.FindUser)
	server.GET("/api/v1/login", userHandlers.UserLogin)
	server.GET("/api/v1/otp", userHandlers.ForgetPassword) 

}