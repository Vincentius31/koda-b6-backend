package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type TransactionRepository struct {
	db *pgx.Conn
}

func NewTransactionRepository(db *pgx.Conn) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, t models.Transaction) error {
	query := `INSERT INTO "transaction" (user_id, transaction_number, delivery_method, subtotal, total, status, payment_method) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, t.UserID, t.TransactionNumber, t.DeliveryMethod, t.Subtotal, t.Total, t.Status, t.PaymentMethod)
	return err
}

func (r *TransactionRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	query := `SELECT id_transaction, user_id, transaction_number, delivery_method, subtotal, total, status, payment_method, created_at 
	          FROM "transaction"`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Transaction])
}

func (r *TransactionRepository) FindByID(ctx context.Context, id int) (*models.Transaction, error) {
	query := `SELECT id_transaction, user_id, transaction_number, delivery_method, subtotal, total, status, payment_method, created_at 
	          FROM "transaction" WHERE id_transaction = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	t, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Transaction])
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TransactionRepository) Update(ctx context.Context, id int, t models.Transaction) error {
	query := `UPDATE "transaction" SET user_id=$1, transaction_number=$2, delivery_method=$3, subtotal=$4, total=$5, status=$6, payment_method=$7 
	          WHERE id_transaction=$8`
	_, err := r.db.Exec(ctx, query, t.UserID, t.TransactionNumber, t.DeliveryMethod, t.Subtotal, t.Total, t.Status, t.PaymentMethod, id)
	return err
}

func (r *TransactionRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM "transaction" WHERE id_transaction = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
