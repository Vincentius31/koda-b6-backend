package models

type ProductCategory struct {
	ProductID  int `json:"product_id"`
	CategoryID int `json:"category_id"`
}

type CreateProductCategoryRequest struct {
	ProductID  int `json:"product_id" binding:"required"`
	CategoryID int `json:"category_id" binding:"required"`
}

type UpdateProductCategoryRequest struct {
	ProductID  *int `json:"product_id"`
	CategoryID *int `json:"category_id"`
}