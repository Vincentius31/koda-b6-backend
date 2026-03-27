package models

type DetailSize struct {
	SizeName        string `json:"size_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type DetailVariant struct {
	VariantName     string `json:"variant_name"`
	AdditionalPrice int    `json:"additional_price"`
}

type ProductDetail struct {
	IDProduct     int             `json:"id_product"`
	Name          string          `json:"name"`
	Desc          string          `json:"desc"`
	Price         int             `json:"price"`
	DiscountRate  float64         `json:"discount_rate"`
	DiscountPrice int             `json:"discount_price"`
	Rating        float64         `json:"rating"`
	TotalReview   int             `json:"total_review"`
	Images        []string        `json:"images"`
	Sizes         []DetailSize    `json:"sizes"`
	Variants      []DetailVariant `json:"variants"`
}

type ProductDetailResponse struct {
	Product     ProductDetail    `json:"product"`
	Recommended []ProductCatalog `json:"recommended"`
}