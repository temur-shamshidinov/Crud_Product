package models

import "github.com/google/uuid"

type Product struct {
	ProductID   uuid.UUID `json:"product_id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
}

type GetProductListResp struct {
	Products []*Product `json:"products"`
	Count    int32      `json:"count"`
}
