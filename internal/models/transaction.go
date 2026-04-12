package models

import "time"

type Transaction struct {
	IDTransaction     int       `json:"id_transaction"`
	UserID           *int      `json:"user_id"`
	TransactionNumber string    `json:"transaction_number"`
	DeliveryMethod    string    `json:"delivery_method"`
	Subtotal          int       `json:"subtotal"`
	Total             int       `json:"total"`
	Status            string    `json:"status"`
	PaymentMethod     string    `json:"payment_method"`
	CreatedAt         time.Time `json:"created_at"`
}

type TransactionListResponse struct {
	IDTransaction     int       `json:"id_transaction"`
	TransactionNumber string    `json:"transaction_number"`
	Total             int       `json:"total"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	FirstItemImage    string    `json:"first_item_image"`
}

type TransactionDetailResponse struct {
	IDTransaction     int                       `json:"id_transaction"`
	TransactionNumber string                    `json:"transaction_number"`
	DeliveryMethod    string                    `json:"delivery_method"`
	Total             int                       `json:"total"`
	Status            string                    `json:"status"`
	PaymentMethod     string                    `json:"payment_method"`
	CreatedAt         time.Time                 `json:"created_at"`
	Customer          CustomerInfo              `json:"customer"`
	Items             []TransactionItemResponse `json:"items"`
}

type CustomerInfo struct {
	Fullname string  `json:"fullname"`
	Email    string  `json:"email"`
	Address  *string `json:"address"`
}

type TransactionItemResponse struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Image       string `json:"image"`
	Quantity    int    `json:"quantity"`
	Size        string `json:"size"`
	Variant     string `json:"variant"`
	Price       int    `json:"price"`
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

type CheckoutItem struct {
	ProductID int    `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	Size      string `json:"size"`
	Variant   string `json:"variant"`
	Price     int    `json:"price" binding:"required"`
}

type CheckoutRequest struct {
	DeliveryMethod string         `json:"delivery_method" binding:"required"`
	Subtotal       int            `json:"subtotal" binding:"required"`
	Total          int            `json:"total" binding:"required"`
	PaymentMethod  string         `json:"payment_method" binding:"required"`
	Items          []CheckoutItem `json:"items" binding:"required,min=1"`
}

type CheckoutResponse struct {
	IDTransaction     int    `json:"id_transaction"`
	TransactionNumber string `json:"transaction_number"`
}