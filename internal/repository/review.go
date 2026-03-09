package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
)

type ReviewRepository struct {
	db *pgx.Conn
}

func NewReviewRepository(db *pgx.Conn) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(ctx context.Context, rev models.Review) error {
	query := `INSERT INTO review (user_id, product_id, messages, rating) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, rev.UserID, rev.ProductID, rev.Messages, rev.Rating)
	return err
}

func (r *ReviewRepository) FindAll(ctx context.Context) ([]models.Review, error) {
	query := `SELECT id_review, user_id, product_id, messages, rating FROM review`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var rev models.Review
		err := rows.Scan(&rev.IDReview, &rev.UserID, &rev.ProductID, &rev.Messages, &rev.Rating)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, rev)
	}
	return reviews, nil
}

func (r *ReviewRepository) FindByID(ctx context.Context, id int) (*models.Review, error) {
	var rev models.Review
	query := `SELECT id_review, user_id, product_id, messages, rating FROM review WHERE id_review = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&rev.IDReview, &rev.UserID, &rev.ProductID, &rev.Messages, &rev.Rating)
	if err != nil {
		return nil, err
	}
	return &rev, nil
}

func (r *ReviewRepository) Update(ctx context.Context, id int, rev models.Review) error {
	query := `UPDATE review SET user_id=$1, product_id=$2, messages=$3, rating=$4 WHERE id_review=$5`
	_, err := r.db.Exec(ctx, query, rev.UserID, rev.ProductID, rev.Messages, rev.Rating, id)
	return err
}

func (r *ReviewRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM review WHERE id_review = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
