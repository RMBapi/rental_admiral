package models

import "time"

type OTP struct {
	Email     string
	Otp       string
	CreatedAt time.Time
	ExpiresAt time.Time
}


