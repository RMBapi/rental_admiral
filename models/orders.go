package models

import (
	"time"

	"gorm.io/gorm"
)


type Orders struct {
	gorm.Model
	OrderId          string    `json:"order_id"`
	CustomerID       uint      `json:"customer_id"` // FK to users.id
	DriverID         *uint      `json:"driver_id"`   // FK to users.id
	VehicleID        *uint      `json:"vehicle_id"`  // FK to vehicles.id
	CreatedBy	     uint      `json:"created_by"`  // FK to users.id
	DispatcherID     *uint      `json:"dispatcher_id"` // FK to users.id
	VehicalNumber     string    `json:"vehicle_number"` 
	VehicleType       int64     `json:"vehicle_type"` 
	PickUpLatitude    float64   `json:"pickup_latitude"`
	PickUpLongitude   float64   `json:"pickup_longitude"`
	PickupAddress     string    `json:"pickup_address"`
	DropOffLatitude   float64   `json:"dropoff_latitude"`	
	DropOffLongitude  float64   `json:"dropoff_longitude"`
	DropoffAddress    string    `json:"dropoff_address"`
	PickupDateTime    time.Time `json:"pickup_date_time"`
	DropoffDateTime   time.Time `json:"dropoff_date_time"`
	TotalPrice        float64   `json:"total_price"`
	PaymentStatus     int8      `json:"payment_status"` // 0 = unpaid, 1 = paid
	PaymentMethod     int8      `json:"payment_method"` // 0 = cash, 1 = card
	Status            int8      `json:"status"`        // 0 = pending, 1 = accepted, 2 = completed, 3 = cancelled

}