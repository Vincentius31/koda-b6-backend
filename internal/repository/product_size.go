package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type ProductSizeRepository struct {
	db *pgx.Conn
}

func NewProductSizeRepository(db *pgx.Conn) *ProductSizeRepository {
	return &ProductSizeRepository{db: db}
}

func (r *ProductSizeRepository) Create(ctx context.Context, s models.ProductSize) error {
	query := `INSERT INTO product_size (product_id, size_name, additional_price) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, s.ProductID, s.SizeName, s.AdditionalPrice)
	return err
}

func (r *ProductSizeRepository) FindAll(ctx context.Context) ([]models.ProductSize, error) {
	query := `SELECT id_size, product_id, size_name, additional_price FROM product_size`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductSize])
}

func (r *ProductSizeRepository) FindByID(ctx context.Context, id int) (*models.ProductSize, error) {
	query := `SELECT id_size, product_id, size_name, additional_price FROM product_size WHERE id_size = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	size, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductSize])
	if err != nil {
		return nil, err
	}
	return &size, nil
}

func (r *ProductSizeRepository) Update(ctx context.Context, id int, s models.ProductSize) error {
	query := `UPDATE product_size SET product_id = $1, size_name = $2, additional_price = $3 WHERE id_size = $4`
	_, err := r.db.Exec(ctx, query, s.ProductID, s.SizeName, s.AdditionalPrice, id)
	return err
}

func (r *ProductSizeRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM product_size WHERE id_size = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}