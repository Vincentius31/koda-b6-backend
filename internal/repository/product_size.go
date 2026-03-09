package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
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

	var sizes []models.ProductSize
	for rows.Next() {
		var s models.ProductSize
		rows.Scan(&s.IDSize, &s.ProductID, &s.SizeName, &s.AdditionalPrice)
		sizes = append(sizes, s)
	}
	return sizes, nil
}

func (r *ProductSizeRepository) FindByID(ctx context.Context, id int) (*models.ProductSize, error) {
	var s models.ProductSize
	query := `SELECT id_size, product_id, size_name, additional_price FROM product_size WHERE id_size = $1`
	err := r.db.QueryRow(ctx, query, id).Scan(&s.IDSize, &s.ProductID, &s.SizeName, &s.AdditionalPrice)
	if err != nil {
		return nil, err
	}
	return &s, nil
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
