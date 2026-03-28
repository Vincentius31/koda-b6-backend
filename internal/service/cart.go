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

func (s *CartService) GetUserCart(ctx context.Context, userID int) ([]models.CartItemResponse, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *CartService) UpdateQty(ctx context.Context, id int, req models.UpdateCartRequest) error {
	if req.Quantity == nil || *req.Quantity < 1 {
		return errors.New("Quantity must be at least 1")
	}
	return s.repo.UpdateQty(ctx, id, *req.Quantity)
}

func (s *CartService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}