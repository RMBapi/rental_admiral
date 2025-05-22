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

}

func SaveUser(user *models.User) error {
  result := DB.Create(user)
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








