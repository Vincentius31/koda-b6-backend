package models

import "time"

type Review struct {
	IDReview  int     `json:"id_review"`
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	Messages  string  `json:"messages"`
	Rating    float64 `json:"rating"`
}

type CreateReviewRequest struct {
	UserID    int     `json:"user_id" binding:"required"`
	ProductID int     `json:"product_id" binding:"required"`
	Messages  string  `json:"messages"`
	Rating    float64 `json:"rating" binding:"required,min=1,max=5"`
}

type UpdateReviewRequest struct {
	UserID    *int     `json:"user_id"`
	ProductID *int     `json:"product_id"`
	Messages  *string  `json:"messages"`
	Rating    *float64 `json:"rating"`
}

type ReviewLanding struct {
	Fullname  string    `db:"fullname" json:"fullname"`
	Messages  string    `db:"messages" json:"messages"`
	Rating    float64   `db:"rating" json:"rating"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}