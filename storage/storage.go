package storage

import (
	log "product/pkg/logger"
	"product/storage/postgres"
	repoi "product/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface{
	GetProductRepo() repoi.ProductRepoI
}

type storage struct{
	productRepo  repoi.ProductRepoI
}

func NewStorage(db *pgx.Conn,log log.Log) StorageI {
	return &storage{
		productRepo: postgres.NewProductRepo(db,log),
	}
}

func (s *storage) GetProductRepo() repoi.ProductRepoI{
	return s.productRepo
}