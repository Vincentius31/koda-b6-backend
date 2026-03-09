package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type TransactionService struct {
	repo *repository.TransactionRepository
}

func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Create(ctx context.Context, req models.CreateTransactionRequest) error {
	status := req.Status
	if status == "" {
		status = "Pending"
	}

	trans := models.Transaction{
		UserID:            req.UserID,
		TransactionNumber: req.TransactionNumber,
		DeliveryMethod:    req.DeliveryMethod,
		Subtotal:          req.Subtotal,
		Total:             req.Total,
		Status:            status,
		PaymentMethod:     req.PaymentMethod,
	}
	return s.repo.Create(ctx, trans)
}

func (s *TransactionService) GetAll(ctx context.Context) ([]models.Transaction, error) {
	return s.repo.FindAll(ctx)
}

func (s *TransactionService) GetByID(ctx context.Context, id int) (*models.Transaction, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *TransactionService) Update(ctx context.Context, id int, req models.UpdateTransactionRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Transaction not found")
	}

	if req.UserID != nil {
		existing.UserID = req.UserID
	}
	if req.TransactionNumber != nil {
		existing.TransactionNumber = *req.TransactionNumber
	}
	if req.DeliveryMethod != nil {
		existing.DeliveryMethod = *req.DeliveryMethod
	}
	if req.Subtotal != nil {
		existing.Subtotal = *req.Subtotal
	}
	if req.Total != nil {
		existing.Total = *req.Total
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.PaymentMethod != nil {
		existing.PaymentMethod = *req.PaymentMethod
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *TransactionService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
