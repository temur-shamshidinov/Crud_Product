package api

import (
	v1 "product/api/handlers/v1"
	log "product/pkg/logger"
	"product/storage"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Storage storage.StorageI
	Log     log.Log
}

func Api(opt Options) *gin.Engine {

	h := v1.NewHandler(v1.Handlers{Storage: opt.Storage, Log: opt.Log})

	engine := gin.Default()

	api := engine.Group("/api")

	api.GET("/ping", h.Ping)

	// products

	pr := api.Group("/pr")
	{
		pr.POST("/create-product", h.CreateProduct)
		pr.GET("/get-products", h.GetProductsList)
		pr.GET("/get-/product/:id", h.GetProduct)
		pr.PUT("/update-products/:id", h.UpdateProduct)
		pr.DELETE("/delete-products/:id", h.DeleteProduct)
	}

	return engine
}
