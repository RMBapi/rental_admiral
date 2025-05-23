package db

import (
	"example.com/rental/models"
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

}

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

func DriveruserId(number string ) (int64,error) {
  var id int64
	err := DB.
		Model(&models.User{}).
		Select("id").
		Where("number = ? AND u_type = ?", number, 6).
		Scan(&id).Error

	return id, err
}








