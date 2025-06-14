package response

import (
	"time"

	"example.com/rental/models"
)


type OrderEvent struct {
	OrderId   string     `json:"order_id"`
	PickupAddress     string    `json:"pickup_address"`
	DropoffAddress    string    `json:"dropoff_address"`
	PickupDateTime    time.Time `json:"pickup_date_time"`
	DropoffDateTime   time.Time `json:"dropoff_date_time"`
	TotalPrice        float64   `json:"total_price"`
	PaymentStatus     int8      `json:"payment_status"` 
	PaymentMethod     int8      `json:"payment_method"` 
	Status            int8      `json:"status"`  

}


func SendOrderResponse(order models.Orders) OrderEvent {
	
	oderinfo := OrderEvent{
		OrderId:          order.OrderId,
		PickupAddress:    order.PickupAddress,
		DropoffAddress:   order.DropoffAddress,
		PickupDateTime:   order.PickupDateTime,
		DropoffDateTime:  order.DropoffDateTime,
		TotalPrice:       order.TotalPrice,
		PaymentStatus:    order.PaymentStatus,
		PaymentMethod:    order.PaymentMethod,
		Status:           order.Status,
	}	

	return oderinfo				
}

