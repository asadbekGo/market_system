package handler

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/models"
	"github.com/asadbekGo/market_system/pkg/helpers"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {

	password := c.GetHeader("Password")
	if password != "1234" {
		handleResponse(c, http.StatusUnauthorized, "The request requires an user authentication.")
		return
	}

	var createCategory models.CreateCategory
	err := c.ShouldBindJSON(&createCategory)
	if err != nil {
		handleResponse(c, 400, "ShouldBindJSON err:"+err.Error())
		return
	}

	if createCategory.ParentID != "" {
		if !helpers.IsValidUUID(createCategory.ParentID) {
			handleResponse(c, http.StatusBadRequest, "parent id is not uuid")
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Category().Create(ctx, &createCategory)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

func (h *Handler) GetByIDCategory(c *gin.Context) {

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Category().GetByID(ctx, &models.CategoryPrimaryKey{Id: id})
	if err == sql.ErrNoRows {
		handleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}

	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) GetListCategory(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query("limit"), 10)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := c.Query("search")
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query search")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	// userId, _ := c.Get("user_id")

	resp, err := h.strg.Category().GetList(ctx, &models.GetListCategoryRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,

		// UserId: userId,
	})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

func (h *Handler) UpdateCategory(c *gin.Context) {

	var updateCategory models.UpdateCategory

	err := c.ShouldBindJSON(&updateCategory)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}
	updateCategory.Id = id

	if updateCategory.ParentID != "" {
		if !helpers.IsValidUUID(updateCategory.ParentID) {
			handleResponse(c, http.StatusBadRequest, "parent id is not uuid")
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.Category().Update(ctx, &updateCategory)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Category().GetByID(ctx, &models.CategoryPrimaryKey{Id: updateCategory.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	var id = c.Param("id")

	if !helpers.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Category().Delete(ctx, &models.CategoryPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
