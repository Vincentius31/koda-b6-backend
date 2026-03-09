package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type TransactionProductRepository struct {
	db *pgx.Conn
}

func NewTransactionProductRepository(db *pgx.Conn) *TransactionProductRepository {
	return &TransactionProductRepository{db: db}
}

func (r *TransactionProductRepository) Create(ctx context.Context, tp models.TransactionProduct) error {
	query := `INSERT INTO transaction_product (transaction_id, product_id, quantity, size, variant, price) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, tp.TransactionID, tp.ProductID, tp.Quantity, tp.Size, tp.Variant, tp.Price)
	return err
}

func (r *TransactionProductRepository) FindAll(ctx context.Context) ([]models.TransactionProduct, error) {
	query := `SELECT id_trans_prod, transaction_id, product_id, quantity, size, variant, price FROM transaction_product`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionProduct])
}

func (r *TransactionProductRepository) FindByID(ctx context.Context, id int) (*models.TransactionProduct, error) {
	query := `SELECT id_trans_prod, transaction_id, product_id, quantity, size, variant, price 
              FROM transaction_product WHERE id_trans_prod = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.TransactionProduct])
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *TransactionProductRepository) Update(ctx context.Context, id int, tp models.TransactionProduct) error {
	query := `UPDATE transaction_product SET transaction_id=$1, product_id=$2, quantity=$3, size=$4, variant=$5, price=$6 
              WHERE id_trans_prod=$7`
	_, err := r.db.Exec(ctx, query, tp.TransactionID, tp.ProductID, tp.Quantity, tp.Size, tp.Variant, tp.Price, id)
	return err
}

func (r *TransactionProductRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM transaction_product WHERE id_trans_prod = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
