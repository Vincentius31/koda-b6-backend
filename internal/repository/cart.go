package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
)

type CartRepository struct {
	db *pgx.Conn
}

func NewCartRepository(db *pgx.Conn) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) Create(ctx context.Context, c models.Cart) error {
	query := `INSERT INTO cart (user_id, product_id, variant_id, size_id, quantity) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, c.UserID, c.ProductID, c.VariantID, c.SizeID, c.Quantity)
	return err
}

func (r *CartRepository) FindAll(ctx context.Context) ([]models.Cart, error) {
	query := `SELECT id_cart, user_id, product_id, variant_id, size_id, quantity FROM cart`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []models.Cart
	for rows.Next() {
		var c models.Cart
		err := rows.Scan(&c.IDCart, &c.UserID, &c.ProductID, &c.VariantID, &c.SizeID, &c.Quantity)
		if err != nil {
			return nil, err
		}
		carts = append(carts, c)
	}
	return carts, nil
}

func (r *CartRepository) FindByID(ctx context.Context, id int) (*models.Cart, error) {
	var c models.Cart
	query := `SELECT id_cart, user_id, product_id, variant_id, size_id, quantity FROM cart WHERE id_cart = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&c.IDCart, &c.UserID, &c.ProductID, &c.VariantID, &c.SizeID, &c.Quantity)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CartRepository) Update(ctx context.Context, id int, c models.Cart) error {
	query := `UPDATE cart SET user_id=$1, product_id=$2, variant_id=$3, size_id=$4, quantity=$5 WHERE id_cart=$6`
	_, err := r.db.Exec(ctx, query, c.UserID, c.ProductID, c.VariantID, c.SizeID, c.Quantity, id)
	return err
}

func (r *CartRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM cart WHERE id_cart = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
