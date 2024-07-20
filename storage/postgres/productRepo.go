package postgres

import (
	"context"
	"product/models"
	log "product/pkg/logger"
	repoi "product/storage/repoI"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type productRepo struct {
	db  *pgx.Conn
	log log.Log
}

func NewProductRepo(db *pgx.Conn, log log.Log) repoi.ProductRepoI {
	return &productRepo{db, log}
}

func (p *productRepo) CreateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	p.log.Debug("request in CreateProduct")

	product.ProductID = uuid.New()

	query := `
		INSERT INTO 
			products (
				product_id,
				name,
				price,
				description

		) VALUES ($1,$2,$3,$4)
	`
	_, err := p.db.Exec(
		ctx, query,
		product.ProductID,
		product.Name,
		product.Price,
		product.Description,
	)
	if err != nil {
		p.log.Error("error on Creating new Product", logger.Error(err))
		return nil, err
	}

	prod, err := p.GetProduct(ctx, product.ProductID.String())
	if err != nil {
		p.log.Error("error on Getting new Product", logger.Error(err))
		return nil, err
	}
	return prod, nil
}
func (p *productRepo) GetProductsList(ctx context.Context, limit, page int32) (*models.GetProductListResp, error) {
	p.log.Debug("request in GetProductsList.")

	query := `
		SELECT
			* 
		FROM 
			products 
		LIMIT 
			$1
		OFFSET
			$2
	`
	offset := (page - 1) * limit
	rows, err := p.db.Query(ctx, query, limit, offset)
	if err != nil {
		p.log.Error("error on Getting all Product ", logger.Error(err))
		return nil, err
	}

	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.ProductID,
			&product.Name,
			&product.Price,
			&product.Description,
		)
		if err != nil {
			p.log.Error("error on scaning  Product ", logger.Error(err))
			return nil, err
		}

		products = append(products, &product)
	}

	var count int32

	err = p.db.QueryRow(ctx, "SELECT count(*) FROM products").Scan(&count)
	if err != nil {
		p.log.Error("error on scaning  Product count ", logger.Error(err))
		return nil, err
	}

	return &models.GetProductListResp{
		Products: products,
		Count:    count,
	}, nil
}

func (p *productRepo) GetProduct(ctx context.Context, id string) (*models.Product, error) {
	p.log.Debug("request in GetProduct.")

	var product models.Product

	query := `SELECT
				 * 
			  FROM 
			  	products 
			  WHERE 
			  	product_id = $1 `

	err := p.db.QueryRow(ctx, query, id).Scan(
		&product.ProductID,
		&product.Name,
		&product.Price,
		&product.Description,
	)

	if err != nil {
		p.log.Error("error on Getting  Product by id", logger.Error(err))
		return nil, err
	}
	return &product, nil
}
func (p *productRepo) UpdateProduct(ctx context.Context, product *models.Product) error {
	p.log.Debug("request in UpdateProduct.")

	query := `
		UPDATE 
			products
		SET 
			name = $1, 
			price = $2, 
			description = $3
		WHERE 
			product_id = $4
	`
	_, err := p.db.Exec(
		ctx, query,
		product.Name,
		product.Price,
		product.Description,
		product.ProductID,
	)
	if err != nil {
		p.log.Error("error on Updating  Product ", logger.Error(err))
		return err
	}
	return nil
}
func (p *productRepo) DeleteProduct(ctx context.Context, id string) error {
	p.log.Debug("request in DeleteProduct.")

	query := `
	
		DELETE 
		FROM 
			products
		WHERE 
			product_id = $1
	`

	_, err := p.db.Exec(ctx, query, id)

	if err != nil {
		p.log.Error("error on Delete Product ", logger.Error(err))
		return err
	}

	return nil
}
