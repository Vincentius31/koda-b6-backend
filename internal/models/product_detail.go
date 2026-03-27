package models

type ProductDetail struct {
	IDProduct     int      `json:"id_product"`
	Name          string   `json:"name"`
	Desc          string   `json:"desc"`
	Price         int      `json:"price"`
	DiscountRate  float64  `json:"discount_rate"`
	DiscountPrice int      `json:"discount_price"`
	Rating        float64  `json:"rating"`
	TotalReview   int      `json:"total_review"`
	Images        []string `json:"images"`
	Sizes         []string `json:"sizes"`
	Variants      []string `json:"variants"`
}

type ProductDetailResponse struct {
	Product     ProductDetail    `json:"product"`
	Recommended []ProductCatalog `json:"recommended"`
}