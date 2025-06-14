package main

import (
	"time"

	"example.com/rental/db"
	"example.com/rental/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main(){
	db.InitDB()
	server := gin.Default()
	server.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8081"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))
	routes.RegisterRoutes(server)
	server.Run(":8080")   	
}
