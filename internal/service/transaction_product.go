package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type TransactionProductService struct {
	repo *repository.TransactionProductRepository
}

func NewTransactionProductService(repo *repository.TransactionProductRepository) *TransactionProductService {
	return &TransactionProductService{repo: repo}
}

func (s *TransactionProductService) Create(ctx context.Context, req models.CreateTransactionProductRequest) error {
	tp := models.TransactionProduct{
		TransactionID: req.TransactionID,
		ProductID:     req.ProductID,
		Quantity:      req.Quantity,
		Size:          req.Size,
		Variant:       req.Variant,
		Price:         req.Price,
	}
	return s.repo.Create(ctx, tp)
}

func (s *TransactionProductService) GetAll(ctx context.Context) ([]models.TransactionProduct, error) {
	return s.repo.FindAll(ctx)
}

func (s *TransactionProductService) GetByID(ctx context.Context, id int) (*models.TransactionProduct, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TransactionProductService) Update(ctx context.Context, id int, req models.UpdateTransactionProductRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("transaction product detail not found")
	}
	if req.TransactionID != nil {
		existing.TransactionID = *req.TransactionID
	}
	if req.ProductID != nil {
		existing.ProductID = req.ProductID
	}
	if req.Quantity != nil {
		existing.Quantity = *req.Quantity
	}
	if req.Size != nil {
		existing.Size = *req.Size
	}
	if req.Variant != nil {
		existing.Variant = *req.Variant
	}
	if req.Price != nil {
		existing.Price = *req.Price
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *TransactionProductService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
