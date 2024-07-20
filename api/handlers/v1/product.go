package v1

import (
	"product/models"
	"product/pkg/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CreateProduct(ctx *gin.Context) {
	var reqBody models.Product

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error in binding req body", logger.Error(err))
		return
	}

	product := &models.Product{}

	helpers.DataParser(reqBody, &product)

	product.ProductID = uuid.New()

	product, err = h.storage.GetProductRepo().CreateProduct(ctx, product)
	if err != nil {
		h.log.Error("error on creating new product", logger.Error(err))
		return
	}

	ctx.JSON(201, product)

}

func (h *handlers) GetProductsList(ctx *gin.Context) {

	page := ctx.Query("page")
	limit := ctx.Query("limit")

	products, err := h.storage.GetProductRepo().GetProductsList(ctx, helpers.GetPage(page), helpers.GetLimit(limit))

	if err != nil {
		h.log.Error("error on creating new product", logger.Error(err))
		return
	}

	ctx.JSON(200, products)
}

func (h *handlers) GetProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	product := &models.Product{}

	product, err := h.storage.GetProductRepo().GetProduct(ctx, id)
	if err != nil {
		h.log.Error("error on getting product", logger.Error(err))
		return
	}

	ctx.JSON(201, product)
}

func (h *handlers) UpdateProduct(ctx *gin.Context) {

	id := ctx.Param("id")

	var reqBody models.Product

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		h.log.Error("error in binding req body", logger.Error(err))
		return
	}

	product, err := h.storage.GetProductRepo().GetProduct(ctx, id)
	if err != nil {
		h.log.Error("error product not found", logger.Error(err))
		return
	}

	helpers.DataParser(reqBody, &product)

	err = h.storage.GetProductRepo().UpdateProduct(ctx, product)
	if err != nil {
		h.log.Error("error: failed to update product", logger.Error(err))
		return
	}

	ctx.JSON(200, product)

}

func (h *handlers) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.storage.GetProductRepo().DeleteProduct(ctx, id)
	if err != nil {
		h.log.Error("error delete product", logger.Error(err))
		return
	}
}
