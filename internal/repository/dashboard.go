package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardRepository struct {
	db *pgxpool.Pool
}

func NewDashboardRepository(db *pgxpool.Pool) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) GetSalesByCategory(ctx context.Context) ([]models.SalesCategory, error) {
	query := `
        SELECT 
            c.name_category AS name, 
            COALESCE(SUM(tp.quantity), 0) AS sales, 
            COALESCE(SUM(tp.quantity * tp.price), 0) AS profit
        FROM category c
        JOIN products_category pc ON c.id_category = pc.category_id
        JOIN transaction_product tp ON pc.product_id = tp.product_id
        JOIN "transaction" t ON tp.transaction_id = t.id_transaction
        WHERE t.status = 'Done' 
        GROUP BY c.name_category
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.SalesCategory])
}

func (r *DashboardRepository) GetBestSellers(ctx context.Context, limit int) ([]models.BestSeller, error) {
	query := `
        SELECT 
            p.name AS product_name, 
            COALESCE(SUM(tp.quantity), 0) AS sold, 
            COALESCE(SUM(tp.quantity * tp.price), 0) AS profit
        FROM products p
        JOIN transaction_product tp ON p.id_product = tp.product_id
        JOIN "transaction" t ON tp.transaction_id = t.id_transaction
        WHERE t.status = 'Done'
        GROUP BY p.id_product, p.name
        ORDER BY sold DESC
        LIMIT $1
    `

	rows, err := r.db.Query(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.BestSeller])
}

func (r *DashboardRepository) GetOrderStats(ctx context.Context) (*models.OrderStats, error) {
	query := `SELECT status, COUNT(id_transaction) FROM "transaction" GROUP BY status`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := models.OrderStats{}

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, err
		}

		switch status {
		case "Pending", "On Progress":
			stats.OnProgress += count
		case "Shipping":
			stats.Shipping += count
		case "Done":
			stats.Done += count
		}
	}

	return &stats, nil
}
