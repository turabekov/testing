package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(ctx context.Context, req *models.CreateProduct) (string, error) {
	var (
		query string
		id    string
	)

	id = uuid.NewString()

	query = `
		INSERT INTO product(
			id, 
			name, 
			category_id,
			description,
			price,
			quantity,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.CategoryId,
		req.Description,
		req.Price,
		req.Quantity,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {

	var (
		query   string
		product models.Product
	)

	query = `
		SELECT
			p.id, 
			p.name, 

			p.category_id,
			c.id,
			c.name,
			CAST(c.created_at::timestamp AS VARCHAR),
			CAST(c.updated_at::timestamp AS VARCHAR),
			
			p.description,
			p.price,
			p.quantity,
			CAST(p.created_at::timestamp AS VARCHAR),
			CAST(p.updated_at::timestamp AS VARCHAR)
		FROM product AS p
		JOIN category AS c ON c.id = p.category_id
		WHERE p.id = $1
	`
	product.CategoryData = &models.Category{}

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.CategoryId,
		&product.CategoryData.Id,
		&product.CategoryData.Name,
		&product.CategoryData.CreatedAt,
		&product.CategoryData.UpdatedAt,
		&product.Description,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepo) GetList(ctx context.Context, req *models.GetListProductRequest) (resp *models.GetListProductResponse, err error) {

	resp = &models.GetListProductResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT
		COUNT(*) OVER(),
		p.id, 
		p.name, 

		p.category_id,
		c.id,
		c.name,
		CAST(c.created_at::timestamp AS VARCHAR),
		CAST(c.updated_at::timestamp AS VARCHAR),

		p.description,
		p.price,
		p.quantity,
		CAST(p.created_at::timestamp AS VARCHAR),
		CAST(p.updated_at::timestamp AS VARCHAR)
	FROM product AS p
	JOIN category AS c ON c.id = p.category_id
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		product.CategoryData = &models.Category{}
		err = rows.Scan(
			&resp.Count,
			&product.Id,
			&product.Name,
			&product.CategoryId,
			&product.CategoryData.Id,
			&product.CategoryData.Name,
			&product.CategoryData.CreatedAt,
			&product.CategoryData.UpdatedAt,
			&product.Description,
			&product.Price,
			&product.Quantity,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Products = append(resp.Products, &product)
	}

	return resp, nil
}

func (r *productRepo) Update(ctx context.Context, req *models.UpdateProduct) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		product
		SET
			id = :id, 
			name = :name, 
			category_id = :category_id,
			description = :description,
			price = :price,
			quantity = :quantity,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":          req.Id,
		"name":        req.Name,
		"category_id": req.CategoryId,
		"description": req.Description,
		"price":       req.Price,
		"quantity":    req.Quantity,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM product
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
