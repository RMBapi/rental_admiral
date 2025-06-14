package userHandlers

import (
	"net/http"

	"example.com/rental/db"
	"github.com/gin-gonic/gin"
)

func FindUser(context *gin.Context){
    DriverNumber := context.Query("driver_number")
	UserNumber := context.Query("user_number")

	if DriverNumber!= ""{
		
		driver_info,err := db.UserProfile(DriverNumber,6)	

	    if err != nil {
		    context.JSON(http.StatusBadRequest, gin.H{"message": "Can't find the diver"})
		    return
	    }

		if driver_info.ID == 0{
			context.JSON(http.StatusBadRequest, gin.H{"message": "Driver not found"})
			return
		}
		
	    context.JSON(http.StatusOK, gin.H{"message": "Successfully Find the Driver" , "event": driver_info })

	}else{

		user_info,err := db.UserProfile(UserNumber,1)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Can't find the user"})
			return
		}

		if user_info.ID == 0{
			context.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Successfully Find the user", "event": user_info })

	}

	
}	


// func GetUserTypeName(userType int8) string {
// 	name := models.UserTypeMap[userType];
// 	if name != "" {
// 		return name
// 	}else{
// 		return "Unknown"
// 	}
// }



		