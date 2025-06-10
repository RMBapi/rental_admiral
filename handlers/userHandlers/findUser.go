package userHandlers

import (
	"net/http"

	"example.com/rental/db"
	"github.com/gin-gonic/gin"
)

func FindUser(context *gin.Context){
    DriverNumber := context.Query("driver_number")
	// UserNumber := context.Query("user_number")

	driver_id,err := db.DriverProfile(DriverNumber)
	// customer_id,err := db.CustomerProfile(UserNumber)
	

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


// func GetUserTypeName(userType int8) string {
// 	name := models.UserTypeMap[userType];
// 	if name != "" {
// 		return name
// 	}else{
// 		return "Unknown"
// 	}
// }



		