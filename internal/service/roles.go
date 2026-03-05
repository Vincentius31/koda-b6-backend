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