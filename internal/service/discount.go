package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type DiscountService struct {
	repo *repository.DiscountRepository
}

func NewDiscountService(repo *repository.DiscountRepository) *DiscountService {
	return &DiscountService{repo: repo}
}

func (s *DiscountService) Create(ctx context.Context, req models.CreateDiscountRequest) error {
	d := models.Discount{
		ProductID:    req.ProductID,
		DiscountRate: req.DiscountRate,
		Description:  req.Description,
		IsFlashSale:  req.IsFlashSale,
	}
	return s.repo.Create(ctx, d)
}

func (s *DiscountService) GetAll(ctx context.Context) ([]models.Discount, error) {
	return s.repo.FindAll(ctx)
}

func (s *DiscountService) GetByID(ctx context.Context, id int) (*models.Discount, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *DiscountService) Update(ctx context.Context, id int, req models.UpdateDiscountRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("discount not found")
	}
	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.DiscountRate != nil {
		existing.DiscountRate = *req.DiscountRate
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.IsFlashSale != nil {
		existing.IsFlashSale = *req.IsFlashSale
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *DiscountService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
