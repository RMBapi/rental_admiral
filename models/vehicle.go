package models

import "gorm.io/gorm"

type Vehical struct {
	gorm.Model
	UserID            *uint      `json:"user_id"` // FK to users.id
	User              User       `gorm:"foreignKey:UserID"` // GORM relation
	Name 	          string     `json:"name"`
	LicenseNumber     string     `json:"license"`
	VehicleNumber     string     `json:"vehicle_number"`
	VehicleType       int64      `json:"type"`

}


const (
    Sadan  = iota + 1   // 1
    Premium             // 2
    Noah                // 3
    Hiece               // 4
                         
)





