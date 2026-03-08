package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductImageService struct {
	repo *repository.ProductImageRepository
}

func NewProductImageService(repo *repository.ProductImageRepository) *ProductImageService {
	return &ProductImageService{repo: repo}
}

func (s *ProductImageService) Create(ctx context.Context, img models.ProductImage) error {
	return s.repo.Create(ctx, img)
}

func (s *ProductImageService) GetAll(ctx context.Context) ([]models.ProductImage, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductImageService) GetByID(ctx context.Context, id int) (*models.ProductImage, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductImageService) Update(ctx context.Context, id int, req models.UpdateProductImageRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Image not found")
	}
	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.Path != nil {
		existing.Path = *req.Path
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *ProductImageService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
