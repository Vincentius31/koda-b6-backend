package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct{
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

// Get all user
func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture 
	          FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.IDUser, &u.RolesID, &u.Fullname, &u.Email, &u.Password, &u.Address, &u.Phone, &u.ProfilePicture)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

// Get User by Id
func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture 
			  FROM users 
			  WHERE id_user = $1`
	var u models.User
	err := r.db.QueryRow(ctx, query, id).Scan(&u.IDUser, &u.RolesID, &u.Fullname, &u.Email, &u.Password, &u.Address, &u.Phone, &u.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Get User by Email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `SELECT id_user, roles_id, fullname, email, password, address, phone, profile_picture FROM users WHERE email = $1`
	var u models.User
	err := r.db.QueryRow(ctx, query, email).Scan(&u.IDUser, &u.RolesID, &u.Fullname, &u.Email, &u.Password, &u.Address, &u.Phone, &u.ProfilePicture)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// Create User
func (r *UserRepository) Create(ctx context.Context, u models.User) error {
	query := `INSERT INTO users (fullname, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, u.Fullname, u.Email, u.Password)
	return err
}

// Update User
func (r *UserRepository) Update(ctx context.Context, id int, u models.User) error {
	query := `UPDATE users SET fullname=$1, email=$2, password=$3, address=$4, phone=$5, profile_picture=$6 WHERE id_user=$7`
	_, err := r.db.Exec(ctx, query, u.Fullname, u.Email, u.Password, u.Address, u.Phone, u.ProfilePicture, id)
	return err
}

// Delete User
func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id_user = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}