package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductVariantService struct {
	repo *repository.ProductVariantRepository
}

func NewProductVariantService(repo *repository.ProductVariantRepository) *ProductVariantService {
	return &ProductVariantService{repo: repo}
}

func (s *ProductVariantService) Create(ctx context.Context, req models.CreateProductVariantRequest) error {
	v := models.ProductVariant{
		ProductID:       req.ProductID,
		VariantName:     req.VariantName,
		AdditionalPrice: req.AdditionalPrice,
	}
	return s.repo.Create(ctx, v)
}

func (s *ProductVariantService) GetAll(ctx context.Context) ([]models.ProductVariant, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductVariantService) GetByID(ctx context.Context, id int) (*models.ProductVariant, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductVariantService) Update(ctx context.Context, id int, req models.UpdateProductVariantRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("variant not found")
	}
	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.VariantName != nil {
		existing.VariantName = *req.VariantName
	}
	if req.AdditionalPrice != nil {
		existing.AdditionalPrice = *req.AdditionalPrice
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *ProductVariantService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
