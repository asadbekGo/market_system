package storage

import "github.com/asadbekGo/market_system/models"

type StorageI interface {
	Category() CategoryRepoI
	Product() ProductRepoI
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	GetByID(req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (int64, error)
	Delete(req *models.CategoryPrimaryKey) error
}

type ProductRepoI interface {
	Create(req *models.CreateProduct) (*models.Product, error)
	GetByID(req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(req *models.UpdateProduct) (int64, error)
	Delete(req *models.ProductPrimaryKey) error
}
