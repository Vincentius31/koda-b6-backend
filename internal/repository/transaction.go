package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
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

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.IDTransaction, &t.UserID, &t.TransactionNumber, &t.DeliveryMethod, &t.Subtotal, &t.Total, &t.Status, &t.PaymentMethod, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}

func (r *TransactionRepository) FindByID(ctx context.Context, id int) (*models.Transaction, error) {
	var t models.Transaction
	query := `SELECT id_transaction, user_id, transaction_number, delivery_method, subtotal, total, status, payment_method, created_at 
	          FROM "transaction" WHERE id_transaction = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&t.IDTransaction, &t.UserID, &t.TransactionNumber, &t.DeliveryMethod, &t.Subtotal, &t.Total, &t.Status, &t.PaymentMethod, &t.CreatedAt)
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
