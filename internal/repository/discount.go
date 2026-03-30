package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DiscountRepository struct {
	db *pgxpool.Pool
}

func NewDiscountRepository(db *pgxpool.Pool) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

func (r *DiscountRepository) Create(ctx context.Context, d models.Discount) error {
	query := `INSERT INTO discount (product_id, discount_rate, description, is_flash_sale) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, d.ProductID, d.DiscountRate, d.Description, d.IsFlashSale)
	return err
}

func (r *DiscountRepository) FindAll(ctx context.Context) ([]models.Discount, error) {
	query := `SELECT id_discount, product_id, discount_rate, description, is_flash_sale FROM discount`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promos []models.Discount
	for rows.Next() {
		var d models.Discount

		err := rows.Scan(
			&d.IDDiscount,
			&d.ProductID,
			&d.DiscountRate,
			&d.Description,
			&d.IsFlashSale,
		)

		if err != nil {
			return nil, err
		}

		promos = append(promos, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return promos, nil
}

func (r *DiscountRepository) FindByID(ctx context.Context, id int) (*models.Discount, error) {
	query := `SELECT id_discount, product_id, discount_rate, description, is_flash_sale FROM discount WHERE id_discount = $1`
	var d models.Discount

	err := r.db.QueryRow(ctx, query, id).Scan(
		&d.IDDiscount,
		&d.ProductID,
		&d.DiscountRate,
		&d.Description,
		&d.IsFlashSale,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DiscountRepository) Update(ctx context.Context, id int, d models.Discount) error {
	query := `UPDATE discount SET product_id=$1, discount_rate=$2, description=$3, is_flash_sale=$4 WHERE id_discount=$5`
	_, err := r.db.Exec(ctx, query, d.ProductID, d.DiscountRate, d.Description, d.IsFlashSale, id)
	return err
}

func (r *DiscountRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM discount WHERE id_discount = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
