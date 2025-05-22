package models

import "gorm.io/gorm"

type Vehical struct {
	gorm.Model
	UserID            int64      `json:"user_id"`
	name 	          string     `json:"name"`
	license_number    string     `json:"license"`
	vehicle_number    string     `json:"vehicle_number"`
	vehicle_type      int64      `json:"type"`

}






