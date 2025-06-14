package userHandlers

import (
	"fmt"
	"net/http"

	"example.com/rental/db"
	"example.com/rental/models"
	"example.com/rental/response"
	"example.com/rental/util"
	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context){

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse user data"})
		return
	}

	user.HashID , err = util.UserHashId()
	if err != nil {
		return
	}

    var temp_password string 

	if user.UType == models.Admin || user.UType == models.SystemManager || user.UType == models.Dispatcher || user.UType == models.Agent{
		password, pass , err := util.UserPassword()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not generate password"})
			return
		}
		user.Password = &pass
		temp_password = password
		
	}else if user.UType == models.Customer || user.UType == models.Driver  {
		user.Password = nil
	}

	err = db.SaveUser(&user)

    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	UEvent := response.SendUserResponse(user,temp_password)
	
	context.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("The %s account is successfully created", UEvent.UType), "event": UEvent})
}




