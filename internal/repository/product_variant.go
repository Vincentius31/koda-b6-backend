package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
)

type ProductVariantRepository struct {
	db *pgx.Conn
}

func NewProductVariantRepository(db *pgx.Conn) *ProductVariantRepository {
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

	var variants []models.ProductVariant
	for rows.Next() {
		var v models.ProductVariant
		rows.Scan(&v.IDVariant, &v.ProductID, &v.VariantName, &v.AdditionalPrice)
		variants = append(variants, v)
	}
	return variants, nil
}

func (r *ProductVariantRepository) FindByID(ctx context.Context, id int) (*models.ProductVariant, error) {
	var v models.ProductVariant
	query := `SELECT id_variant, product_id, variant_name, additional_price FROM product_variant WHERE id_variant = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&v.IDVariant, &v.ProductID, &v.VariantName, &v.AdditionalPrice)
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
