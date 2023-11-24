package api

import (
	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/market_system/api/handler"
	"github.com/asadbekGo/market_system/config"
	"github.com/asadbekGo/market_system/storage"
)

func SetUpApi(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {

	handler := handler.NewHandler(cfg, strg)

	// Category ...
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIDCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

}
