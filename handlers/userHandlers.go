package handlers

import (
	"net/http"

	"example.com/rental/db"
	"example.com/rental/models"
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

	if user.UType == models.Admin || user.UType == models.SystemManager || user.UType == models.Dispatcher || user.UType == models.Agent {
		pass , err := util.UserPassword()
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not generate password"})
			return
		}
		user.Password = &pass
		
		
	}else if user.UType == models.Customer {
		user.Password = nil
	}

	err = db.SaveUser(&user)

    if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}


    if user.UType == models.Admin || user.UType == models.SystemManager || user.UType == models.Dispatcher || user.UType == models.Agent {
		context.JSON(http.StatusCreated, gin.H{"message": "Created", "event": user})
	}else if user.UType == models.Customer {
		context.JSON(http.StatusCreated, gin.H{"message": "Customer Created", "event": user})	
	}

}	


