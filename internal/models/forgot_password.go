package models

import "time"

type ForgotPassword struct {
	IDRequest int       `db:"id_request" json:"id_request"`
	Email     string    `db:"email" json:"email"`
	OTPCode   int       `db:"otp_code" json:"otp_code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type VerifyOTPRequest struct {
	Email   string `json:"email" binding:"required,email"`
	OTPCode int    `json:"otp_code" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}