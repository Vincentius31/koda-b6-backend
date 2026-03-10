package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"strings"
)

type RoleService struct {
	repo *repository.RoleRepository
}

func NewRoleService(repo *repository.RoleRepository) *RoleService {
	return &RoleService{repo: repo}
}

func (s *RoleService) Create(ctx context.Context, req models.CreateRoleRequest) error {
	if strings.TrimSpace(req.NameRoles) == "" {
		return errors.New("Role name cannot be empty or just spaces")
	}

	role := models.Role{
		NameRoles: req.NameRoles,
	}
	return s.repo.Create(ctx, role)
}

func (s *RoleService) GetAll(ctx context.Context) ([]models.Role, error) {
	return s.repo.FindAll(ctx)
}

func (s *RoleService) GetByID(ctx context.Context, id int) (*models.Role, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *RoleService) Update(ctx context.Context, id int, req models.UpdateRoleRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Role not found")
	}

	if req.NameRoles != nil {
		if strings.TrimSpace(*req.NameRoles) == "" {
			return errors.New("Role name cannot be empty")
		}
		existing.NameRoles = *req.NameRoles
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *RoleService) Delete(ctx context.Context, id int) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("role not found")
	}
	return s.repo.Delete(ctx, id)
}