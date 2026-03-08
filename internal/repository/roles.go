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

func (r *RoleRepository) FindByID(ctx context.Context, id int) (*models.Role, error) {
	query := `SELECT id_roles, name_roles FROM roles WHERE id_roles = $1`
	var role models.Role
	err := r.db.QueryRow(ctx, query, id).Scan(&role.IDRoles, &role.NameRoles)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) Update(ctx context.Context, id int, role models.Role) error {
	query := `UPDATE roles SET name_roles = $1 WHERE id_roles = $2`
	_, err := r.db.Exec(ctx, query, role.NameRoles, id)
	return err
}

func (r *RoleRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM roles WHERE id_roles = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}