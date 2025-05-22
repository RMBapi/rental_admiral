package main

import (
	"example.com/rental/db"
	"example.com/rental/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")   	
}
