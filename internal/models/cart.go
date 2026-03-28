package models

type Cart struct {
	IDCart    int  `json:"id_cart" db:"id_cart"`
	UserID    int  `json:"user_id" db:"user_id"`
	ProductID int  `json:"product_id" db:"product_id"`
	VariantID *int `json:"variant_id" db:"variant_id"`
	SizeID    *int `json:"size_id" db:"size_id"`
	Quantity  int  `json:"quantity" db:"quantity"`
}

type CreateCartRequest struct {
	UserID    int  `json:"user_id"`
	ProductID int  `json:"product_id" binding:"required"`
	VariantID *int `json:"variant_id"`
	SizeID    *int `json:"size_id"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type UpdateCartRequest struct {
	Quantity *int `json:"quantity" binding:"required,min=1"`
}

type CartItemResponse struct {
	IDCart       int     `json:"id_cart" db:"id_cart"`
	ProductID    int     `json:"product_id" db:"product_id"`
	ProductName  string  `json:"product_name" db:"product_name"`
	ProductImage string  `json:"product_image" db:"product_image"`
	BasePrice    int     `json:"base_price" db:"base_price"`
	VariantID    *int    `json:"variant_id" db:"variant_id"`
	VariantName  *string `json:"variant_name" db:"variant_name"`
	VariantPrice int     `json:"variant_price" db:"variant_price"`
	SizeID       *int    `json:"size_id" db:"size_id"`
	SizeName     *string `json:"size_name" db:"size_name"`
	SizePrice    int     `json:"size_price" db:"size_price"`
	Quantity     int     `json:"quantity" db:"quantity"`
}