package storage

import "github.com/asadbekGo/market_system/models"

type StorageI interface {
	Category() CategoryRepoI
}

type CategoryRepoI interface {
	Create(req *models.CreateCategory) (*models.Category, error)
	GetByID(req *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Update(req *models.UpdateCategory) (int64, error)
	Delete(req *models.CategoryPrimaryKey) error
}
