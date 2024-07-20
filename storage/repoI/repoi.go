package repoi

import (
	"context"
	"product/models"
)

type ProductRepoI interface {
	CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	GetProductsList(ctx context.Context, page, limit int32) (*models.GetProductListResp, error)
	GetProduct(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, id string) error
}
