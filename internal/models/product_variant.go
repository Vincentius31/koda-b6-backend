package models

type ProductVariant struct {
	IDVariant       int    `json:"id_variant"`
	ProductID       int    `json:"product_id"`
	VariantName     string `json:"variant_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type CreateProductVariantRequest struct {
	ProductID       int    `json:"product_id" binding:"required"`
	VariantName     string `json:"variant_name" binding:"required"`
	AdditionalPrice int    `json:"additional_price"`
}

type UpdateProductVariantRequest struct {
	ProductID       *int    `json:"product_id"`
	VariantName     *string `json:"variant_name"`
	AdditionalPrice *int    `json:"additional_price"`
}