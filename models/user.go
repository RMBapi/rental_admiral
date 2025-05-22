package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	HashID       string     `json:"hash_id"`
    Name         string     `json:"name"`
    Email 	     *string    `json:"email"`
	Password     *string    `json:"password"`
    Number       string     `json:"phone"`
	UType		 int8       `json:"type"`
}


const (
    Customer  = iota + 1      // 1
    SystemManager             // 2
    Dispatcher                // 3
    Agent                     // 4
    Admin                     // 5
    Driver                    // 6
)


