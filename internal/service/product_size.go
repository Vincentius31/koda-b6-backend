package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductSizeService struct {
	repo *repository.ProductSizeRepository
}

func NewProductSizeService(repo *repository.ProductSizeRepository) *ProductSizeService {
	return &ProductSizeService{repo: repo}
}

func (s *ProductSizeService) Create(ctx context.Context, req models.CreateProductSizeRequest) error {
	size := models.ProductSize{
		ProductID:       req.ProductID,
		SizeName:        req.SizeName,
		AdditionalPrice: req.AdditionalPrice,
	}
	return s.repo.Create(ctx, size)
}

func (s *ProductSizeService) GetAll(ctx context.Context) ([]models.ProductSize, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductSizeService) GetByID(ctx context.Context, id int) (*models.ProductSize, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductSizeService) Update(ctx context.Context, id int, req models.UpdateProductSizeRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("size not found")
	}
	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.SizeName != nil {
		existing.SizeName = *req.SizeName
	}
	if req.AdditionalPrice != nil {
		existing.AdditionalPrice = *req.AdditionalPrice
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *ProductSizeService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
