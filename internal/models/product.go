package models

type Product struct {
	IDProduct int    `json:"id_product"`
	Name      string `json:"name"`
	Desc      string `json:"desc"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	IsActive  bool   `json:"is_active"`
}

type CreateProductRequest struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"desc"`
	Price    int    `json:"price" binding:"required,min=1"`
	Quantity int    `json:"quantity" binding:"min=0"`
	IsActive bool   `json:"is_active"`
}

type UpdateProductRequest struct {
	Name     *string `json:"name"`
	Desc     *string `json:"desc"`
	Price    *int    `json:"price"`
	Quantity *int    `json:"quantity"`
	IsActive *bool   `json:"is_active"`
}

type ProductLanding struct {
    IDProduct   int    `db:"id_product" json:"id_product"`
    Name        string `db:"name" json:"name"`
    Desc        string `db:"desc" json:"desc"`
    Price       int    `db:"price" json:"price"`
    ImagePath   string `db:"image_path" json:"image_path"`
    TotalReview int    `db:"total_review" json:"total_review"`
}

type ProductCatalog struct {
	IDProduct     int     `db:"id_product" json:"id_product"`
	Name          string  `db:"name" json:"name"`
	Desc          string  `db:"desc" json:"desc"`
	Price         int     `db:"price" json:"price"`
	DiscountRate  float64 `db:"discount_rate" json:"discount_rate"`
	DiscountPrice int     `db:"discount_price" json:"discount_price"`
	Rating        float64 `db:"rating" json:"rating"`
	ImagePath     string  `db:"image_path" json:"image_path"`
}

type PagingMeta struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
}

type ProductCatalogResponse struct {
	Items []ProductCatalog `json:"items"`
	Meta  PagingMeta       `json:"meta"`
}