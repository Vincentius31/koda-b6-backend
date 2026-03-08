package models

type Category struct {
	IDCategory   int    `json:"id_category"`
	NameCategory string `json:"name_category"`
}

type CreateCategoryRequest struct {
	NameCategory string `json:"name_category" binding:"required"`
}