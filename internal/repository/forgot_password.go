package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ForgotPasswordRepository struct {
	db *pgxpool.Pool
}

func NewForgotPasswordRepository(db *pgxpool.Pool) *ForgotPasswordRepository {
	return &ForgotPasswordRepository{
		db: db,
	}
}

func (r *ForgotPasswordRepository) CreateForgetRequest(ctx context.Context, email string, code int) error {
	query := `INSERT INTO forgot_password (email, otp_code) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, email, code)
	return err
}

func (r *ForgotPasswordRepository) GetDataByEmailAndCode(ctx context.Context, email string, code int) (*models.ForgotPassword, error) {
	query := `SELECT id_request, email, otp_code, created_at 
              FROM forgot_password 
              WHERE email = $1 AND otp_code = $2`

	rows, err := r.db.Query(ctx, query, email, code)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	data, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ForgotPassword])
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *ForgotPasswordRepository) DeleteByCode(ctx context.Context, code int) error {
	query := `DELETE FROM forgot_password WHERE otp_code = $1`
	_, err := r.db.Exec(ctx, query, code)
	return err
}
