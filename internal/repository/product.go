package repository

import (
	"context"
	"fmt"
	"koda-b6-backend/internal/models"
	"strconv"
	"strings"

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

func (r *ProductRepository) buildCatalogFilters(params map[string]string) (string, []interface{}) {
	var conditions []string
	var args []interface{}
	counter := 1

	// 1. Filter Search
	if search := params["search"]; search != "" {
		conditions = append(conditions, fmt.Sprintf("p.name ILIKE $%d", counter))
		args = append(args, "%"+search+"%")
		counter++
	}

	// 2. Filter Category
	if cat := params["category"]; cat != "" {
		conditions = append(conditions, fmt.Sprintf("c.name_category = $%d", counter))
		args = append(args, cat)
		counter++
	}

	// 3. Filter Price Range
	minPrice := params["min_price"]
	maxPrice := params["max_price"]
	if minPrice != "" && maxPrice != "" {
		conditions = append(conditions, fmt.Sprintf("p.price BETWEEN $%d AND $%d", counter, counter+1))
		args = append(args, minPrice, maxPrice)
		counter += 2
	}

	whereSQL := ""
	if len(conditions) > 0 {
		whereSQL = " AND " + strings.Join(conditions, " AND ")
	}
	return whereSQL, args
}

func (r *ProductRepository) GetCatalog(ctx context.Context, params map[string]string) (*models.ProductCatalogResponse, error) {
	page, _ := strconv.Atoi(params["page"])
	if page < 1 {
		page = 1
	}
	limit := 6
	offset := (page - 1) * limit

	whereSQL, args := r.buildCatalogFilters(params)

	// 1. Pagination
	var total int
	countQuery := `
		SELECT COUNT(DISTINCT p.id_product) 
		FROM products p
		LEFT JOIN products_category pc ON p.id_product = pc.product_id
		LEFT JOIN category c ON pc.category_id = c.id_category
		WHERE p.is_active = TRUE ` + whereSQL
	_ = r.db.QueryRow(ctx, countQuery, args...).Scan(&total)

	// 2. Tampilkan data sesuai pagination
	fetchQuery := `
		SELECT 
			p.id_product, p.name, p.desc, p.price,
			COALESCE(d.discount_rate, 0) as discount_rate,
			CAST(p.price - (p.price * COALESCE(d.discount_rate, 0)) AS INT) as discount_price,
			COALESCE(AVG(rv.rating), 0) as rating,
			COALESCE((SELECT path FROM product_images WHERE product_id = p.id_product LIMIT 1), '') as image_path
		FROM products p
		LEFT JOIN discount d ON p.id_product = d.product_id
		LEFT JOIN review rv ON p.id_product = rv.product_id
		LEFT JOIN products_category pc ON p.id_product = pc.product_id
		LEFT JOIN category c ON pc.category_id = c.id_category
		WHERE p.is_active = TRUE ` + whereSQL + `
		GROUP BY p.id_product, d.discount_rate
		ORDER BY p.id_product DESC
		LIMIT $` + fmt.Sprint(len(args)+1) + ` OFFSET $` + fmt.Sprint(len(args)+2)

	finalArgs := append(args, limit, offset)
	rows, err := r.db.Query(ctx, fetchQuery, finalArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductCatalog])

	return &models.ProductCatalogResponse{
		Items: items,
		Meta: models.PagingMeta{
			TotalItems:  total,
			TotalPages:  (total + limit - 1) / limit,
			CurrentPage: page,
		},
	}, err
}

func (r *ProductRepository) GetFullDetailByID(ctx context.Context, id int) (*models.ProductDetail, error) {
	query := `
		SELECT 
			p.id_product, p.name, p.desc, p.price,
			COALESCE(d.discount_rate, 0) as discount_rate,
			CAST(p.price - (p.price * COALESCE(d.discount_rate, 0)) AS INT) as discount_price,
			COALESCE(AVG(rv.rating), 0) as rating,
			COUNT(rv.id_review) as total_review
		FROM products p
		LEFT JOIN discount d ON p.id_product = d.product_id
		LEFT JOIN review rv ON p.id_product = rv.product_id
		WHERE p.id_product = $1
		GROUP BY p.id_product, d.discount_rate`

	var detail models.ProductDetail
	err := r.db.QueryRow(ctx, query, id).Scan(
		&detail.IDProduct, &detail.Name, &detail.Desc, &detail.Price,
		&detail.DiscountRate, &detail.DiscountPrice, &detail.Rating, &detail.TotalReview,
	)
	if err != nil {
		return nil, err
	}

	imgRows, _ := r.db.Query(ctx, "SELECT path FROM product_images WHERE product_id = $1", id)
	defer imgRows.Close()
	for imgRows.Next() {
		var path string
		imgRows.Scan(&path)
		detail.Images = append(detail.Images, path)
	}

	sizeRows, _ := r.db.Query(ctx, "SELECT id_size, size_name, additional_price FROM product_size WHERE product_id = $1", id)
	defer sizeRows.Close()
	for sizeRows.Next() {
		var s models.DetailSize
		sizeRows.Scan(&s.IDSize, &s.SizeName, &s.AdditionalPrice)
		detail.Sizes = append(detail.Sizes, s)
	}

	varRows, _ := r.db.Query(ctx, "SELECT id_variant, variant_name, additional_price FROM product_variant WHERE product_id = $1", id)
	defer varRows.Close()
	for varRows.Next() {
		var v models.DetailVariant
		varRows.Scan(&v.IDVariant, &v.VariantName, &v.AdditionalPrice)
		detail.Variants = append(detail.Variants, v)
	}

	return &detail, nil
}

func (r *ProductRepository) GetRandomRecommended(ctx context.Context, excludeID int, limit int) ([]models.ProductCatalog, error) {
	query := `
		SELECT 
			p.id_product, p.name, p.desc, p.price,
			COALESCE(d.discount_rate, 0) as discount_rate,
			CAST(p.price - (p.price * COALESCE(d.discount_rate, 0)) AS INT) as discount_price,
			COALESCE(AVG(rv.rating), 0) as rating,
			COALESCE((SELECT path FROM product_images WHERE product_id = p.id_product LIMIT 1), '') as image_path
		FROM products p
		LEFT JOIN discount d ON p.id_product = d.product_id
		LEFT JOIN review rv ON p.id_product = rv.product_id
		WHERE p.is_active = TRUE AND p.id_product != $1
		GROUP BY p.id_product, d.discount_rate
		ORDER BY RANDOM() 
		LIMIT $2
	`

	rows, err := r.db.Query(ctx, query, excludeID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[models.ProductCatalog])
}