package models

import (
	"time"

	"gorm.io/gorm"
)


type Orders struct {
	gorm.Model
	OrderId          string     `json:"orderd"`
	CustomerID       int        `json:"customerId"` // FK to users.id
	DriverID         *int       `json:"driverId"`   // FK to users.id
	VehicleID        *int       `json:"vehicleId"`  // FK to vehicles.id
	CreatedBy	     *int       `json:"createdBy"`  // FK to users.id
	DispatcherID     *int       `json:"dispatcherId"` // FK to users.id
	VehicalNumber    *string    `json:"vehicleNumber"` 
	VehicleType       int64     `json:"vehicleType"` 
	PickUpLatitude    float64   `json:"pickupLatitude"`
	PickUpLongitude   float64   `json:"pickupLongitude"`
	PickupAddress     string    `json:"pickupAddress"`
	DropOffLatitude   float64   `json:"dropoffLatitude"`	
	DropOffLongitude  float64   `json:"dropoffLongitude"`
	DropoffAddress    string    `json:"dropoffAddress"`
	PickupDateTime    time.Time `json:"pickupDateTime"`
	DropoffDateTime   time.Time `json:"dropoffDateTime"`
	TotalPrice        float64   `json:"totalPrice"`
	PaymentStatus     int8      `json:"paymentStatus"` // 1 = paid, 2 = unpaid
	PaymentMethod     int8      `json:"paymentMethod"` // 1 = cash, 2 = card
	Status            int8      `json:"status"`        // 1 = pending, 2 = accepted, 3 = completed, 4 = cancelled

}