package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ReviewRepository struct {
	db *pgxpool.Pool
}

func NewReviewRepository(db *pgxpool.Pool) *ReviewRepository {
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

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Review])
}

func (r *ReviewRepository) FindByID(ctx context.Context, id int) (*models.Review, error) {
	query := `SELECT id_review, user_id, product_id, messages, rating FROM review WHERE id_review = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rev, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Review])
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

func (r *ReviewRepository) GetLatestReviews(ctx context.Context) ([]models.ReviewLanding, error) {
	query := `
        SELECT 
            u.fullname AS fullname, 
            u.profile_picture AS profile_picture,
            r.messages AS messages, 
            r.rating AS rating
        FROM review r
        JOIN users u ON r.user_id = u.id_user
        ORDER BY r.id_review DESC
        LIMIT 5
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ReviewLanding])
}
