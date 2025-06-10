package vehicle

import (
	"net/http"

	"example.com/rental/db"
	"example.com/rental/models"

	"github.com/gin-gonic/gin"
)

func CreateVehicle(context *gin.Context){

	var vehicle models.Vehical

	err := context.ShouldBindJSON(&vehicle)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse vehicle data"})
		return
	}

	err = db.SaveVehicle(&vehicle)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the vehicle info"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Vehicle Created successfully"})
}	
