package storage

import (
	"context"

	"github.com/asadbekGo/market_system/models"
)

type StorageI interface {
	Category() CategoryRepoI
	Product() ProductRepoI
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *models.CreateCategory) (*models.Category, error)
	GetByID(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(ctx context.Context, req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(ctx context.Context, req *models.CreateProduct) (*models.Product, error)
	GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) error
}
