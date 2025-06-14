package db

import (
	"example.com/rental/models"
	"example.com/rental/response"
	_ "github.com/mattn/go-sqlite3"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DB *gorm.DB
var err error


func InitDB() {
  DB, err = gorm.Open(sqlite.Open("rentals.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  DB.AutoMigrate(&models.User{})
  DB.AutoMigrate(&models.Vehical{})
  DB.AutoMigrate(&models.Orders{})
  DB.AutoMigrate(&models.OTP{})

//   CreateTable()

}

// func CreateTable(){

// 	createOtpTable := `
// 	CREATE TABLE IF NOT EXISTS otps (
// 		email TEXT NOT NULL,
// 		otp TEXT NOT NULL,
// 		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
// 		expires_at DATETIME NOT NULL
// 	);`

// 	err := DB.Exec(createOtpTable)

// 	if err.Error != nil {
// 		panic("Couldn't create otps table: " )
// 	}


// }

func SaveUser(user *models.User) error {
  result := DB.Create(user)
  return result.Error
}

func SaveVehicle(vehicle *models.Vehical) error {
  result := DB.Create(vehicle)
  return result.Error
}


func UniqueHashIdCheck(key string) (int64,error) {
 var count int64
	err := DB.Model(&models.User{}).
		Where("hash_id = ?", key).
		Count(&count).
		Error

    return count,err

}

func DriverProfile(number string) (response.DriverInfo, error) {

	var driver response.DriverInfo

	err := DB.
		Model(&models.User{}).
		Select("id", "name").
		Where("number = ? AND u_type = ?", number, 6).
		Scan(&driver).Error

	return driver, err
}


func FindUserByEmail(email string) (bool, error) {
	var id int64
	err := DB.
		Model(&models.User{}).
		Select("id").
		Where("email = ? ", email).
		Scan(&id).Error

	if id!= 0 {
		return true, nil
	}	

	return false, err

}


func CustomerProfile(number string ) (int64,error) {
  var id int64
	err := DB.
		Model(&models.User{}).
		Select("id").
		Where("number = ? AND u_type = ?", number, 6).
		Scan(&id).Error

	return id, err
}


func GetUserPassword(email string)(string,error){
	var password string
	err := DB.
		Model(&models.User{}).
		Select("password").
		Where("email = ?", email).
		Scan(&password).Error

	if err != nil {
		return "", err
	}

	return password, nil
}










