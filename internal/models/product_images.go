package models

type ProductImage struct {
	IDImage   int    `json:"id_image"`
	ProductID int    `json:"product_id"`
	Path      string `json:"path"`
}

type CreateProductImageRequest struct {
	ProductID int    `json:"product_id" binding:"required"`
	Path      string `json:"path" binding:"required"`
}

type UpdateProductImageRequest struct {
	ProductID *int    `json:"product_id"`
	Path      *string `json:"path"`
}