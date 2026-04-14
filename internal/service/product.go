package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetPromos(ctx context.Context) ([]string, error) {
	return s.repo.GetAvailablePromos(ctx)
}

func (s *ProductService) Create(ctx context.Context, req models.AdminProductPayload) error {
	if req.PriceProduct <= 0 {
		return errors.New("price must be greater than zero")
	}
	return s.repo.Create(ctx, req)
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.AdminProductPayload, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id int) (*models.AdminProductPayload, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, id int, req models.AdminProductPayload) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.repo.Update(ctx, id, req)
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

func (s *ProductService) UpdateImages(ctx context.Context, id int, paths []string) error {
	_, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("product not found")
	}
	return s.repo.UpdateImages(ctx, id, paths)
}
