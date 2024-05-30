package usecase

import (
	"context"
	"net/http"

	"github.com/nutikuli/internProject_backend/internal/models/product-category"
	"github.com/nutikuli/internProject_backend/internal/models/product-category/entities"
)

type productCateUsecase struct {
	productCateRepo product_category.ProductCategoryRepository
}

func NewProductCategoryUsecase(prodCate product_category.ProductCategoryRepository) product_category.ProductCategoryUsecase {
	return &productCateUsecase{
		productCateRepo: prodCate,
	}
}

// OnCreateProductCategory implements product_category.ProductCategoryUsecase.

func (p *productCateUsecase) OnCreateProductCategory(ctx context.Context, req entities.ProductCategoryCreatedReq) (*int64, int, error) {
	created, err := p.productCateRepo.CreateProductCategory(ctx, req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return created, http.StatusOK, nil
}

// OnDeleteProductCategoryById implements product_category.ProductCategoryUsecase.
func (p *productCateUsecase) OnDeleteProductCategoryById(ctx context.Context, categoryId *int64) (int, error) {
	err := p.productCateRepo.DeleteProductCategoryById(ctx, categoryId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// OnGetProductCategoriesByStoreId implements product_category.ProductCategoryUsecase.
func (p *productCateUsecase) OnGetProductCategoriesByStoreId(ctx context.Context, storeId *int64) ([]*entities.ProductCategory, int, error) {
	categories, err := p.productCateRepo.GetProductCategoriesByStoreId(ctx, storeId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return categories, http.StatusOK, nil
}

// OnGetProductCategoryById implements product_category.ProductCategoryUsecase.
func (p *productCateUsecase) OnGetProductCategoryById(ctx context.Context, categoryId *int64) (*entities.ProductCategory, int, error) {
	category, err := p.productCateRepo.GetProductCategoryById(ctx, categoryId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return category, http.StatusOK, nil
}
