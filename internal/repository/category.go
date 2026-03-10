package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"koda-b6-backend/internal/models"
)

type CategoryRepository struct {
	db *pgx.Conn
}

func NewCategoryRepository(db *pgx.Conn) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category models.Category) error {
	query := `INSERT INTO category (name_category) VALUES ($1)`
	_, err := r.db.Exec(ctx, query, category.NameCategory)
	return err
}

func (r *CategoryRepository) FindAll(ctx context.Context) ([]models.Category, error) {
	query := `SELECT id_category, name_category FROM category`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Category])
}

func (r *CategoryRepository) FindByID(ctx context.Context, id int) (*models.Category, error) {
	query := `SELECT id_category, name_category FROM category WHERE id_category = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cat, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Category])
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *CategoryRepository) Update(ctx context.Context, id int, category models.Category) error {
	query := `UPDATE category SET name_category = $1 WHERE id_category = $2`
	_, err := r.db.Exec(ctx, query, category.NameCategory, id)
	return err
}

func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM category WHERE id_category = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
