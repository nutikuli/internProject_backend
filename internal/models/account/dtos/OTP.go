package dtos

import "time"

type OTPres struct {
	OTP       string    `json:"otp"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"time"`
}