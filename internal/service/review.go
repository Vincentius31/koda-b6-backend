package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ReviewService struct {
	repo *repository.ReviewRepository
}

func NewReviewService(repo *repository.ReviewRepository) *ReviewService {
	return &ReviewService{repo: repo}
}

func (s *ReviewService) Create(ctx context.Context, req models.CreateReviewRequest) error {
	rev := models.Review{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Messages:  req.Messages,
		Rating:    req.Rating,
	}
	return s.repo.Create(ctx, rev)
}

func (s *ReviewService) GetAll(ctx context.Context) ([]models.Review, error) {
	return s.repo.FindAll(ctx)
}

func (s *ReviewService) GetByID(ctx context.Context, id int) (*models.Review, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ReviewService) Update(ctx context.Context, id int, req models.UpdateReviewRequest) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("Review not found")
	}

	if req.UserID != nil {
		existing.UserID = *req.UserID
	}
	if req.ProductID != nil {
		existing.ProductID = *req.ProductID
	}
	if req.Messages != nil {
		existing.Messages = *req.Messages
	}
	if req.Rating != nil {
		existing.Rating = *req.Rating
	}

	return s.repo.Update(ctx, id, *existing)
}

func (s *ReviewService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
