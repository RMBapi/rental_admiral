package userHandlers

import (
	"net/http"

	"example.com/rental/util"
	"github.com/gin-gonic/gin"
)

func ForgetPassword(context *gin.Context){
	otp,err := util.GenerateOTP()
	
	  if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create the user"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "OTP generated successfully", "otp": otp})

	
}