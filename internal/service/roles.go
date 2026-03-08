package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"strings"
)

type RoleService struct{
	repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService{
	return &RoleService{
		repo: repo,
	}
}

func (s *RoleService) Create(ctx context.Context, req models.CreateRoleRequest) error{
	if strings.TrimSpace(req.NameRoles) == "" {
		return errors.New("Role name cannot be empty")
	}

	role := models.Role{
		NameRoles: req.NameRoles,
	}

	return s.repo.Create(ctx, role)
}

func (s *RoleService) GetAll(ctx context.Context)([]models.Role, error){
	return s.repo.FindAll(ctx)
}

func (s *RoleService) GetByID(ctx context.Context, id int) (*models.Role, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("role not found")
	}
	return role, nil
}

func (s *RoleService) Update(ctx context.Context, id int, req models.CreateRoleRequest) error {
	if strings.TrimSpace(req.NameRoles) == "" {
		return errors.New("role name cannot be empty")
	}

	// Cek apakah role ada
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("role not found")
	}

	role := models.Role{NameRoles: req.NameRoles}
	return s.repo.Update(ctx, id, role)
}

func (s *RoleService) Delete(ctx context.Context, id int) error {
	// Cek apakah role ada
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("role not found")
	}

	return s.repo.Delete(ctx, id)
}