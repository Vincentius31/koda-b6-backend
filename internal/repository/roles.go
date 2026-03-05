package repository

import (
	"context"
	"koda-b6-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type RoleRepository struct{
	db *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) *RoleRepository{
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) Create(ctx context.Context, role models.Role) error {
	query := `INSERT INTO roles (name_roles) VALUES ($1)`
	_, err := r.db.Exec(ctx, query, role.NameRoles)
	return err
}

func (r *RoleRepository) FindAll(ctx context.Context)([]models.Role, error){
	query := `SELECT id_roles, name_roles FROM roles`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next(){
		var role models.Role
		if err := rows.Scan(&role.IDRoles, &role.NameRoles); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}