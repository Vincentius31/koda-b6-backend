package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"strings"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(ctx context.Context, req models.CreateCategoryRequest) error {
	if strings.TrimSpace(req.NameCategory) == "" {
		return errors.New("category name cannot be empty")
	}
	cat := models.Category{NameCategory: req.NameCategory}
	return s.repo.Create(ctx, cat)
}

func (s *CategoryService) GetAll(ctx context.Context) ([]models.Category, error) {
	return s.repo.FindAll(ctx)
}

func (s *CategoryService) GetByID(ctx context.Context, id int) (*models.Category, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CategoryService) Update(ctx context.Context, id int, req models.CreateCategoryRequest) error {
	cat := models.Category{NameCategory: req.NameCategory}
	return s.repo.Update(ctx, id, cat)
}

func (s *CategoryService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
