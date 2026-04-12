package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionRepository struct {
	db *pgxpool.Pool
}

func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(ctx context.Context, t models.Transaction) error {
	query := `INSERT INTO "transaction" (user_id, transaction_number, delivery_method, subtotal, total, status, payment_method) 
              VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.db.Exec(ctx, query, t.UserID, t.TransactionNumber, t.DeliveryMethod, t.Subtotal, t.Total, t.Status, t.PaymentMethod)
	return err
}

func (r *TransactionRepository) CreateReturningID(ctx context.Context, t models.Transaction) (int, string, error) {
	query := `INSERT INTO "transaction" (user_id, transaction_number, delivery_method, subtotal, total, status, payment_method) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id_transaction, transaction_number`
	var id int
	var transNum string
	err := r.db.QueryRow(ctx, query, t.UserID, t.TransactionNumber, t.DeliveryMethod, t.Subtotal, t.Total, t.Status, t.PaymentMethod).Scan(&id, &transNum)
	return id, transNum, err
}

func (r *TransactionRepository) CreateWithTx(ctx context.Context, tx pgx.Tx, t models.Transaction) (int, string, error) {
	query := `INSERT INTO "transaction" (user_id, transaction_number, delivery_method, subtotal, total, status, payment_method) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) 
              RETURNING id_transaction, transaction_number`
	var id int
	var transNum string
	err := tx.QueryRow(ctx, query, t.UserID, t.TransactionNumber, t.DeliveryMethod, t.Subtotal, t.Total, t.Status, t.PaymentMethod).Scan(&id, &transNum)
	return id, transNum, err
}

func (r *TransactionRepository) FindAll(ctx context.Context) ([]models.Transaction, error) {
	query := `SELECT id_transaction, user_id, transaction_number, delivery_method, subtotal, total, status, payment_method, created_at 
              FROM "transaction" ORDER BY created_at DESC`
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

func (r *TransactionRepository) UpdateStatus(ctx context.Context, id int, status string) error {
	query := `UPDATE "transaction" SET status = $1 WHERE id_transaction = $2`
	_, err := r.db.Exec(ctx, query, status, id)
	return err
}

func (r *TransactionRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM "transaction" WHERE id_transaction = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// PERBAIKAN DI SINI: Menambahkan ::date di belakang $1
func (r *TransactionRepository) CountByDate(ctx context.Context, date string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM "transaction" WHERE DATE(created_at) = $1::date`
	err := r.db.QueryRow(ctx, query, date).Scan(&count)
	return count, err
}

func (r *TransactionRepository) FindByUserID(ctx context.Context, userID int) ([]models.TransactionListResponse, error) {
	query := `
        SELECT 
            t.id_transaction,
            t.transaction_number,
            t.total,
            t.status,
            t.created_at,
            COALESCE((SELECT pi.path FROM transaction_product tp 
                JOIN product_images pi ON pi.product_id = tp.product_id 
                WHERE tp.transaction_id = t.id_transaction LIMIT 1), '') as first_item_image
        FROM "transaction" t
        WHERE t.user_id = $1
        ORDER BY t.created_at DESC
    `
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionListResponse])
}

func (r *TransactionRepository) FindDetailByID(ctx context.Context, transactionID int, userID int) (*models.TransactionDetailResponse, error) {
	query := `
        SELECT 
            t.id_transaction,
            t.transaction_number,
            t.delivery_method,
            t.total,
            t.status,
            t.payment_method,
            t.created_at,
            u.fullname,
            u.email,
            u.address
        FROM "transaction" t
        JOIN users u ON t.user_id = u.id_user
        WHERE t.id_transaction = $1 AND t.user_id = $2
    `
	row := r.db.QueryRow(ctx, query, transactionID, userID)

	var resp models.TransactionDetailResponse
	var customer models.CustomerInfo
	err := row.Scan(
		&resp.IDTransaction,
		&resp.TransactionNumber,
		&resp.DeliveryMethod,
		&resp.Total,
		&resp.Status,
		&resp.PaymentMethod,
		&resp.CreatedAt,
		&customer.Fullname,
		&customer.Email,
		&customer.Address,
	)
	if err != nil {
		return nil, err
	}
	resp.Customer = customer

	itemsQuery := `
        SELECT 
            tp.product_id,
            p.name as product_name,
            COALESCE((SELECT pi.path FROM product_images pi WHERE pi.product_id = tp.product_id LIMIT 1), '') as image,
            tp.quantity,
            tp.size,
            tp.variant,
            tp.price
        FROM transaction_product tp
        JOIN products p ON tp.product_id = p.id_product
        WHERE tp.transaction_id = $1
    `
	rows, err := r.db.Query(ctx, itemsQuery, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.TransactionItemResponse])
	if err != nil {
		return nil, err
	}
	resp.Items = items

	return &resp, nil
}

func (r *TransactionRepository) CreateTransactionProduct(ctx context.Context, tx pgx.Tx, tp models.TransactionProduct) error {
	query := `INSERT INTO transaction_product (transaction_id, product_id, quantity, size, variant, price) 
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := tx.Exec(ctx, query, tp.TransactionID, tp.ProductID, tp.Quantity, tp.Size, tp.Variant, tp.Price)
	return err
}
