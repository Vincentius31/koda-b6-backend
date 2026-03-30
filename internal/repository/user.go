package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture, created_at, updated_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture, created_at, updated_at FROM users WHERE id_user = $1`
	rows, err := r.db.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture, created_at, updated_at FROM users WHERE email = $1`
	rows, err := r.db.Query(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, u models.User) error {
	query := `INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, u.Fullname, u.Email, u.Password)
	return err
}

func (r *UserRepository) Update(ctx context.Context, id int, u models.User) error {
	query := `UPDATE users SET fullname=$1, email=$2, password=$3, address=$4, phone=$5, profile_picture=$6, updated_at=CURRENT_TIMESTAMP WHERE id_user=$7`
	_, err := r.db.Exec(ctx, query, u.Fullname, u.Email, u.Password, u.Address, u.Phone, u.ProfilePicture, id)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id_user = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
