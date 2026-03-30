package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductCategoryRepository struct {
	db *pgxpool.Pool
}

func NewProductCategoryRepository(db *pgxpool.Pool) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		db: db,
	}
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

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductCategory])
}

func (r *ProductCategoryRepository) FindByID(ctx context.Context, prodID int, catID int) (*models.ProductCategory, error) {
	query := `SELECT product_id, category_id FROM products_category WHERE product_id = $1 AND category_id = $2`
	rows, err := r.db.Query(ctx, query, prodID, catID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pc, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.ProductCategory])
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
