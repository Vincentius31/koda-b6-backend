package service

import (
	"context"
	"errors"
	"koda-b6-backend/internal/models"
	"koda-b6-backend/internal/repository"
)

type ProductCategoryService struct {
	repo *repository.ProductCategoryRepository
}

func NewProductCategoryService(repo *repository.ProductCategoryRepository) *ProductCategoryService {
	return &ProductCategoryService{repo: repo}
}

func (s *ProductCategoryService) Create(ctx context.Context, req models.CreateProductCategoryRequest) error {
	pc := models.ProductCategory{
		ProductID: req.ProductID, 
		CategoryID: req.CategoryID,
	}
	return s.repo.Create(ctx, pc)
}

func (s *ProductCategoryService) GetAll(ctx context.Context) ([]models.ProductCategory, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductCategoryService) GetByID(ctx context.Context, prodID int, catID int) (*models.ProductCategory, error) {
	return s.repo.FindByID(ctx, prodID, catID)
}

func (s *ProductCategoryService) Update(ctx context.Context, oldP int, oldC int, req models.UpdateProductCategoryRequest) error {
	existing, err := s.repo.FindByID(ctx, oldP, oldC)
	if err != nil {
		return errors.New("Data not found")
	}

	newProductID := existing.ProductID
	if req.ProductID != nil {
		newProductID = *req.ProductID
	}

	newCategoryID := existing.CategoryID
	if req.CategoryID != nil {
		newCategoryID = *req.CategoryID
	}

	updatedData := models.ProductCategory{
		ProductID:  newProductID,
		CategoryID: newCategoryID,
	}
	
	return s.repo.Update(ctx, oldP, oldC, updatedData)
}

func (s *ProductCategoryService) Delete(ctx context.Context, prodID int, catID int) error {
	return s.repo.Delete(ctx, prodID, catID)
}