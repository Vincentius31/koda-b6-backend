package models

type Cart struct {
	IDCart    int  `json:"id_cart"`
	UserID    int  `json:"user_id"`
	ProductID int  `json:"product_id"`
	VariantID *int `json:"variant_id"` 
	SizeID    *int `json:"size_id"`   
	Quantity  int  `json:"quantity"`
}

type CreateCartRequest struct {
	UserID    int  `json:"user_id" binding:"required"`
	ProductID int  `json:"product_id" binding:"required"`
	VariantID *int `json:"variant_id"`
	SizeID    *int `json:"size_id"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartRequest struct {
	UserID    *int `json:"user_id"`
	ProductID *int `json:"product_id"`
	VariantID *int `json:"variant_id"`
	SizeID    *int `json:"size_id"`
	Quantity  *int `json:"quantity"`
}