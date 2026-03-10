package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type CartService struct {
	repo *repository.CartRepository
}

func NewCartService(repo *repository.CartRepository) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) Create(ctx context.Context, req models.CreateCartRequest) error {
	cart := models.Cart{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		VariantID: req.VariantID,
		SizeID:    req.SizeID,
		Quantity:  req.Quantity,
	}
	return s.repo.Create(ctx, cart)
}

func (s *CartService) GetAll(ctx context.Context) ([]models.Cart, error) {
	return s.repo.FindAll(ctx)
}

func (s *CartService) GetByID(ctx context.Context, id int) (*models.Cart, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CartService) Update(ctx context.Context, id int, req models.UpdateCartRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Cart item not found")
	}

	if req.UserID != nil {
		existing.UserID = *req.UserID
	}
	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.VariantID != nil {
		existing.VariantID = req.VariantID
	}
	if req.SizeID != nil {
		existing.SizeID = req.SizeID
	}
	if req.Quantity != nil {
		if *req.Quantity < 1 {
			return errors.New("Quantity must be at least 1")
		}
		existing.Quantity = *req.Quantity
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *CartService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
