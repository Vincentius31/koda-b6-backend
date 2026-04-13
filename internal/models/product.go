package models

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

type AdminProductPayload struct {
	ID            int      `json:"id"` // React menggunakan "id", bukan "id_product"
	NameProduct   string   `json:"nameProduct" binding:"required"`
	PriceProduct  int      `json:"priceProduct" binding:"required"`
	PriceDiscount int      `json:"priceDiscount"`
	Description   string   `json:"description"`
	Stock         int      `json:"stock"`
	Size          []string `json:"size"`
	Temp          []string `json:"temp"`
	Method        []string `json:"method"`
	ImageProduct  []string `json:"imageProduct"`
	Category      string   `json:"category"`
	PromoType     string   `json:"promoType"`
}