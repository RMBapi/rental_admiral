package order

import (
	"fmt"
	"net/http"

	"example.com/rental/db"
	"example.com/rental/models"
	"example.com/rental/response"
	"example.com/rental/util"
	"github.com/gin-gonic/gin"
)

func CreateOrder(context *gin.Context){

	var order models.Orders

	err := context.ShouldBindJSON(&order)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse order data"})
		return
	}

    order.OrderId,err = util.GenerateHexID(6)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not generate order ID"})
		return
	}

	err = db.SaveOrder(&order)

    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the order"})
		return
	}

	fmt.Println(order)
	

	orderResponse := response.SendOrderResponse(order)

	context.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "event": orderResponse})


}


