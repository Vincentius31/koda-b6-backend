package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductVariantRepository struct {
	db *pgxpool.Pool
}

func NewProductVariantRepository(db *pgxpool.Pool) *ProductVariantRepository {
	return &ProductVariantRepository{db: db}
}

func (r *ProductVariantRepository) Create(ctx context.Context, v models.ProductVariant) error {
	query := `INSERT INTO product_variant (product_id, variant_name, additional_price) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, v.ProductID, v.VariantName, v.AdditionalPrice)
	return err
}

func (r *ProductVariantRepository) FindAll(ctx context.Context) ([]models.ProductVariant, error) {
	query := `SELECT id_variant, product_id, variant_name, additional_price FROM product_variant`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductVariant])
}

func (r *ProductVariantRepository) FindByID(ctx context.Context, id int) (*models.ProductVariant, error) {
	query := `SELECT id_variant, product_id, variant_name, additional_price FROM product_variant WHERE id_variant = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	v, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductVariant])
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *ProductVariantRepository) Update(ctx context.Context, id int, v models.ProductVariant) error {
	query := `UPDATE product_variant SET product_id = $1, variant_name = $2, additional_price = $3 WHERE id_variant = $4`
	_, err := r.db.Exec(ctx, query, v.ProductID, v.VariantName, v.AdditionalPrice, id)
	return err
}

func (r *ProductVariantRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM product_variant WHERE id_variant = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
