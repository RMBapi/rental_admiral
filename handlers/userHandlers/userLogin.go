package userHandlers

import (
	"fmt"
	"net/http"
	"net/url"

	"example.com/rental/db"
	"example.com/rental/util"
	"github.com/gin-gonic/gin"
)

func UserLogin(context *gin.Context) {
	emailEncoded := context.Query("email")
	fmt.Println("Encoded Email:", emailEncoded)
	email, err := url.QueryUnescape(emailEncoded)
	fmt.Println("Email:", email)
	if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
    password := context.Query("password")

    if email == "" || password == "" {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
        return
    }


	hashPass,err := db.GetUserPassword(email)

    if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Can't fatch the user password"})
		return
	}

	authUser := util.CheckPasswordHash(password, hashPass)

	if authUser == true{
		context.JSON(http.StatusOK, gin.H{"message": "Successfully logged in"})

	}else{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}
    
}