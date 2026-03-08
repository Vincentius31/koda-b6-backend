package repository

import (
	"context"
	"koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
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

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.IDProduct, &p.Name, &p.Desc, &p.Price, &p.Quantity, &p.IsActive); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(ctx context.Context, id int) (*models.Product, error) {
	query := `SELECT id_product, name, "desc", price, quantity, is_active FROM products WHERE id_product = $1`
	var p models.Product
	err := r.db.QueryRow(ctx, query, id).Scan(&p.IDProduct, &p.Name, &p.Desc, &p.Price, &p.Quantity, &p.IsActive)
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