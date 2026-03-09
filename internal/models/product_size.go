package models

type ProductSize struct {
	IDSize          int    `json:"id_size"`
	ProdudctID      int    `json:"product_id"`
	SizeName        string `json:"size_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type CreateProductSizeRequest struct {
	ProductID       int    `json:"product_id" binding:"required"`
	SizeName        string `json:"size_name" binding:"required"`
	AdditionalPrice int    `json:"additional_price"`
}

type UpdateProductSizeRequest struct {
	ProductID       *int    `json:"product_id" binding:"required"`
	SizeName        *string `json:"size_name" binding:"required"`
	AdditionalPrice *int    `json:"additional_price"`
}
