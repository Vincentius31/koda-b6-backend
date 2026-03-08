package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
)

type ProductImageRepository struct {
	db *pgx.Conn
}

func NewProductImageRepository(db *pgx.Conn) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

func (r *ProductImageRepository) Create(ctx context.Context, img models.ProductImage) error {
	query := `INSERT INTO product_images (product_id, path) VALUES ($1, $2)`
	_, err := r.db.Exec(ctx, query, img.ProductID, img.Path)
	return err
}

func (r *ProductImageRepository) FindAll(ctx context.Context) ([]models.ProductImage, error) {
	query := `SELECT id_image, product_id, path FROM product_images`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.ProductImage
	for rows.Next() {
		var img models.ProductImage
		if err := rows.Scan(&img.IDImage, &img.ProductID, &img.Path); err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}

func (r *ProductImageRepository) FindByID(ctx context.Context, id int) (*models.ProductImage, error) {
	query := `SELECT id_image, product_id, path FROM product_images WHERE id_image = $1`
	var img models.ProductImage
	err := r.db.QueryRow(ctx, query, id).Scan(&img.IDImage, &img.ProductID, &img.Path)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func (r *ProductImageRepository) Update(ctx context.Context, id int, img models.ProductImage) error {
	query := `UPDATE product_images SET product_id = $1, path = $2 WHERE id_image = $3`
	_, err := r.db.Exec(ctx, query, img.ProductID, img.Path, id)
	return err
}

func (r *ProductImageRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM product_images WHERE id_image = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
