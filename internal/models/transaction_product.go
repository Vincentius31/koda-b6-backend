package models

type TransactionProduct struct {
	IDTransProd   int    `json:"id_trans_prod"`
	TransactionID int    `json:"transaction_id"`
	ProductID     *int   `json:"product_id"` 
	Quantity      int    `json:"quantity"`
	Size          string `json:"size"`
	Variant       string `json:"variant"`
	Price         int    `json:"price"`
}

type CreateTransactionProductRequest struct {
	TransactionID int    `json:"transaction_id" binding:"required"`
	ProductID     *int   `json:"product_id"`
	Quantity      int    `json:"quantity" binding:"required,min=1"`
	Size          string `json:"size"`
	Variant       string `json:"variant"`
	Price         int    `json:"price" binding:"required"`
}

type UpdateTransactionProductRequest struct {
	TransactionID *int    `json:"transaction_id"`
	ProductID     *int    `json:"product_id"`
	Quantity      *int    `json:"quantity"`
	Size          *string `json:"size"`
	Variant       *string `json:"variant"`
	Price         *int    `json:"price"`
}
