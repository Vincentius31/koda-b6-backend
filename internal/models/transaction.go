package models

import "time"

type Transaction struct {
	IDTransaction     int       `json:"id_transaction"`
	UserID            *int      `json:"user_id"`
	TransactionNumber string    `json:"transaction_number"`
	DeliveryMethod    string    `json:"delivery_method"`
	Subtotal          int       `json:"subtotal"`
	Total             int       `json:"total"`
	Status            string    `json:"status"`
	PaymentMethod     string    `json:"payment_method"`
	CreatedAt         time.Time `json:"created_at"`
}

type CreateTransactionRequest struct {
	UserID            *int   `json:"user_id"`
	TransactionNumber string `json:"transaction_number" binding:"required"`
	DeliveryMethod    string `json:"delivery_method" binding:"required"`
	Subtotal          int    `json:"subtotal" binding:"required"`
	Total             int    `json:"total" binding:"required"`
	Status            string `json:"status"`
	PaymentMethod     string `json:"payment_method" binding:"required"`
}

type UpdateTransactionRequest struct {
	UserID            *int    `json:"user_id"`
	TransactionNumber *string `json:"transaction_number"`
	DeliveryMethod    *string `json:"delivery_method"`
	Subtotal          *int    `json:"subtotal"`
	Total             *int    `json:"total"`
	Status            *string `json:"status"`
	PaymentMethod     *string `json:"payment_method"`
}
