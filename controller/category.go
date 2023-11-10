package controller

import (
	"errors"
	"log"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
)

func (c *Conn) CreateCategory(req *models.CreateCategory) (*models.Category, error) {

	log.Println(config.Info, "Create Category resp >>>>> ", req)
	resp, err := c.storage.Category().Create(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetByIDCategory(req *models.CategoryPrimaryKey) (*models.Category, error) {

	log.Println(config.Info, "GetByID Category resp >>>>> ", req)
	resp, err := c.storage.Category().GetByID(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) GetListCategory(req *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error) {

	log.Println(config.Info, "GetList Category resp >>>>> ", req)
	resp, err := c.storage.Category().GetList(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) UpdateCategory(req *models.UpdateCategory) (*models.Category, error) {

	log.Println(config.Info, "Update Category resp >>>>> ", req)
	rowsAffected, err := c.storage.Category().Update(req)
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, errors.New("no rows affected")
	}

	resp, err := c.storage.Category().GetByID(&models.CategoryPrimaryKey{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Conn) DeleteCategory(req *models.CategoryPrimaryKey) error {

	log.Println(config.Info, "Delete Category resp >>>>> ", req)
	err := c.storage.Category().Delete(req)
	if err != nil {
		return err
	}

	return nil
}
