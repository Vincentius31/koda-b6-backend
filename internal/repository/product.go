package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
)

type ProductRepository struct {
	db *pgx.Conn
}

func NewProductRepository(db *pgx.Conn) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, p models.Product) error {
	query := `INSERT INTO products (name, "desc", price, quantity, is_active) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, p.Name, p.Desc, p.Price, p.Quantity, p.IsActive)
	return err
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]models.Product, error) {
	query := `SELECT id_product, name, "desc", price, quantity, is_active FROM products`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
}

func (r *ProductRepository) FindByID(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id_product, name, "desc", price, quantity, is_active FROM products WHERE id_product = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Update(ctx context.Context, id int, p models.Product) error {
	query := `UPDATE products SET name=$1, "desc"=$2, price=$3, quantity=$4, is_active=$5 WHERE id_product=$6`
	_, err := r.db.Exec(ctx, query, p.Name, p.Desc, p.Price, p.Quantity, p.IsActive, id)
	return err
}

func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id_product = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *ProductRepository) GetRecommended(ctx context.Context) ([]models.ProductLanding, error) {
	query := `
        SELECT 
            p.id_product, 
            p.name, 
            p.desc, 
            p.price,
            (SELECT pi.path FROM product_images pi WHERE pi.product_id = p.id_product LIMIT 1) as image_path,
            COUNT(rv.id_review) as total_review
        FROM products p
        LEFT JOIN review rv ON p.id_product = rv.product_id
        WHERE p.is_active = TRUE
        GROUP BY p.id_product
        ORDER BY total_review DESC
        LIMIT 4
    `
	
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductLanding])
}
