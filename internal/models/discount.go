package models

type Discount struct {
    IDDiscount   int     `json:"id_discount" db:"id_discount"`
    ProductID    int     `json:"product_id" db:"product_id"`
    DiscountRate float64 `json:"discount_rate" db:"discount_rate"`
    Description  string  `json:"description" db:"description"`
    IsFlashSale  bool    `json:"is_flash_sale" db:"is_flash_sale"`
}

type CreateDiscountRequest struct {
    ProductID    int     `json:"product_id" binding:"required"`
    DiscountRate float64 `json:"discount_rate" binding:"required"`
    Description  string  `json:"description"`
    IsFlashSale  bool    `json:"is_flash_sale"`
}

type UpdateDiscountRequest struct {
    ProductID    *int     `json:"product_id"`
    DiscountRate *float64 `json:"discount_rate"`
    Description  *string  `json:"description"`
    IsFlashSale  *bool    `json:"is_flash_sale"`
}