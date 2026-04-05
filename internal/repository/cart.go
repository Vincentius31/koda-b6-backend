package repository

import (
	"context"
	"errors"
	"fmt"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CartRepository struct {
	db *pgxpool.Pool
}

func NewCartRepository(db *pgxpool.Pool) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) Create(ctx context.Context, c models.Cart) error {
	var existingID, existingQty int

	queryCheck := `SELECT id_cart, quantity FROM cart WHERE user_id = $1 AND product_id = $2 `
	args := []interface{}{c.UserID, c.ProductID}
	counter := 3

	if c.VariantID != nil {
		queryCheck += fmt.Sprintf(`AND variant_id = $%d `, counter)
		args = append(args, *c.VariantID)
		counter++
	} else {
		queryCheck += `AND variant_id IS NULL `
	}

	if c.SizeID != nil {
		queryCheck += fmt.Sprintf(`AND size_id = $%d`, counter)
		args = append(args, *c.SizeID)
	} else {
		queryCheck += `AND size_id IS NULL`
	}

	err := r.db.QueryRow(ctx, queryCheck, args...).Scan(&existingID, &existingQty)

	if err == nil {
		newQty := existingQty + c.Quantity
		_, updateErr := r.db.Exec(ctx, `UPDATE cart SET quantity = $1 WHERE id_cart = $2`, newQty, existingID)
		return updateErr
	}

	if errors.Is(err, pgx.ErrNoRows) || err.Error() == "No rows in result set" {
		queryInsert := `INSERT INTO cart (user_id, product_id, variant_id, size_id, quantity) VALUES ($1, $2, $3, $4, $5)`
		_, errInsert := r.db.Exec(ctx, queryInsert, c.UserID, c.ProductID, c.VariantID, c.SizeID, c.Quantity)
		return errInsert
	}

	return err
}

func (r *CartRepository) FindByUserID(ctx context.Context, userID int) ([]models.CartItemResponse, error) {
	query := `
        SELECT 
            c.id_cart,
            c.product_id,
            p.name as product_name,
            COALESCE((SELECT path FROM product_images WHERE product_id = p.id_product LIMIT 1), '') as product_image,
            CAST(p.price - (p.price * COALESCE(d.discount_rate, 0)) AS INT) as base_price,
            c.variant_id,
            pv.variant_name,
            COALESCE(pv.additional_price, 0) as variant_price,
            c.size_id,
            ps.size_name,
            COALESCE(ps.additional_price, 0) as size_price,
            c.quantity
        FROM cart c
        JOIN products p ON c.product_id = p.id_product
        LEFT JOIN discount d ON p.id_product = d.product_id
        LEFT JOIN product_variant pv ON c.variant_id = pv.id_variant
        LEFT JOIN product_size ps ON c.size_id = ps.id_size
        WHERE c.user_id = $1
        ORDER BY c.id_cart DESC
    `

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.CartItemResponse])
}

func (r *CartRepository) UpdateQty(ctx context.Context, id int, qty int) error {
	query := `UPDATE cart SET quantity=$1 WHERE id_cart=$2`
	_, err := r.db.Exec(ctx, query, qty, id)
	return err
}

func (r *CartRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM cart WHERE id_cart = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *CartRepository) DeleteByUserIDWithTx(ctx context.Context, tx pgx.Tx, userID int) error {
	query := `DELETE FROM cart WHERE user_id = $1`
	_, err := tx.Exec(ctx, query, userID)
	return err
}