package service

import (
	"context"
	"errors"
	"fmt"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService struct {
	db       *pgxpool.Pool
	repo     *repository.TransactionRepository
	cartRepo *repository.CartRepository
}

func NewTransactionService(db *pgxpool.Pool, repo *repository.TransactionRepository, cartRepo *repository.CartRepository) *TransactionService {
	return &TransactionService{
		db:       db,
		repo:     repo,
		cartRepo: cartRepo,
	}
}

func (s *TransactionService) generateTransactionNumber(ctx context.Context) (string, error) {
	now := time.Now()
	datePart := now.Format("20060102") 

	dateForQuery := now.Format("2006-01-02")

	count, err := s.repo.CountByDate(ctx, dateForQuery)
	if err != nil {
		return "", err
	}

	indexOrder := fmt.Sprintf("%04d", count+1)

	return fmt.Sprintf("ORD-%s-%s", datePart, indexOrder), nil
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

func (s *TransactionService) Checkout(ctx context.Context, userID int, req models.CheckoutRequest) (*models.CheckoutResponse, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("no items to checkout")
	}

	transNum, err := s.generateTransactionNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to generate transaction number: %w", err)
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	trans := models.Transaction{
		UserID:            &userID,
		TransactionNumber: transNum,
		DeliveryMethod:    req.DeliveryMethod,
		Subtotal:          req.Subtotal,
		Total:             req.Total,
		Status:            "Pending",
		PaymentMethod:     req.PaymentMethod,
	}

	transactionID, _, err := s.repo.CreateWithTx(ctx, tx, trans)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	for _, item := range req.Items {
		tp := models.TransactionProduct{
			TransactionID: transactionID,
			ProductID:     &item.ProductID,
			Quantity:      item.Quantity,
			Size:          item.Size,
			Variant:       item.Variant,
			Price:         item.Price,
		}
		if err := s.repo.CreateTransactionProduct(ctx, tx, tp); err != nil {
			return nil, fmt.Errorf("failed to add transaction product: %w", err)
		}
	}

	if err := s.cartRepo.DeleteByUserIDWithTx(ctx, tx, userID); err != nil {
		return nil, fmt.Errorf("failed to clear cart: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return &models.CheckoutResponse{
		IDTransaction:     transactionID,
		TransactionNumber: transNum,
	}, nil
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

func (s *TransactionService) GetByUserID(ctx context.Context, userID int) ([]models.TransactionListResponse, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *TransactionService) GetDetailByID(ctx context.Context, transactionID int, userID int) (*models.TransactionDetailResponse, error) {
	return s.repo.FindDetailByID(ctx, transactionID, userID)
}
