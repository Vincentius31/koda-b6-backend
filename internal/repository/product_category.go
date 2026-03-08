package repository

import (
	"context"
	"koda-b6-backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ProductCategoryRepository struct {
	db *pgx.Conn
}

func NewProductCategoryRepository(db *pgx.Conn) *ProductCategoryRepository {
	return &ProductCategoryRepository{db: db}
}

func (r *ProductCategoryRepository) Create(ctx context.Context, pc models.ProductCategory) error {
	query := `INSERT INTO products_category (product_id, category_id) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, pc.ProductID, pc.CategoryID)
	return err
}

func (r *ProductCategoryRepository) FindAll(ctx context.Context) ([]models.ProductCategory, error) {
	query := `SELECT product_id, category_id FROM products_category`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ProductCategory
	for rows.Next() {
		var pc models.ProductCategory
		if err := rows.Scan(&pc.ProductID, &pc.CategoryID); err != nil {
			return nil, err
		}
		results = append(results, pc)
	}
	return results, nil
}

func (r *ProductCategoryRepository) FindByID(ctx context.Context, prodID int, catID int) (*models.ProductCategory, error) {
	query := `SELECT product_id, category_id FROM products_category WHERE product_id = $1 AND category_id = $2`
	var pc models.ProductCategory
	err := r.db.QueryRow(ctx, query, prodID, catID).Scan(&pc.ProductID, &pc.CategoryID)
	if err != nil {
		return nil, err
	}
	return &pc, nil
}

func (r *ProductCategoryRepository) Update(ctx context.Context, oldProdID int, oldCatID int, pc models.ProductCategory) error {
	query := `UPDATE products_category SET product_id = $1, category_id = $2 WHERE product_id = $3 AND category_id = $4`
	_, err := r.db.Exec(ctx, query, pc.ProductID, pc.CategoryID, oldProdID, oldCatID)
	return err
}

func (r *ProductCategoryRepository) Delete(ctx context.Context, prodID int, catID int) error {
	query := `DELETE FROM products_category WHERE product_id = $1 AND category_id = $2`
	_, err := r.db.Exec(ctx, query, prodID, catID)
	return err
}