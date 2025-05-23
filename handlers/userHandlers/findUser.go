package userHandlers

import (
	"net/http"

	"example.com/rental/db"
	"github.com/gin-gonic/gin"
)

func FindDriver(context *gin.Context){
    DriverNumber := context.Query("driver_number")

	driver_id,err := db.DriveruserId(DriverNumber)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't find the diver id"})
		return
	}

	if driver_id == 0{
		context.JSON(http.StatusBadRequest, gin.H{"message": "Driver not found"})
		return
	}
		
	context.JSON(http.StatusOK, gin.H{"message": "Successfully Find the driver id", "Driver id": driver_id})

}	
		