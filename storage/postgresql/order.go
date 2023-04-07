package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (r *orderRepo) Create(ctx context.Context, req *models.CreateOrder) (string, error) {
	var (
		query string
		id    string
	)

	id = uuid.NewString()

	query = `
		INSERT INTO orders(
			id, 
			client_id, 
			price,
			status,
			updated_at
		)
		VALUES ($1, $2, $3, $4, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.ClientId,
		helper.NewNullInt32(int(req.Price)),
		helper.NewNullString(req.Status),
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *orderRepo) GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error) {
	var (
		query string
		order models.Order
	)

	query = `
		SELECT
			o.id, 
			o.client_id, 

			c.id, 
			c.first_name,
			c.last_name,
			c.phone_number,
			CAST(c.created_at::timestamp AS VARCHAR),
			CAST(c.updated_at::timestamp AS VARCHAR),

			COALESCE(o.price, 0),
			COALESCE(o.status, ''),
			CAST(o.created_at::timestamp AS VARCHAR),
			CAST(o.updated_at::timestamp AS VARCHAR)
		FROM "orders" AS o
		JOIN client AS c ON c.id = o.client_id
		WHERE o.id = $1
	`

	order.ClientData = &models.Client{}

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&order.Id,
		&order.ClientId,
		&order.ClientData.Id,
		&order.ClientData.FirstName,
		&order.ClientData.LastName,
		&order.ClientData.PhoneNumber,
		&order.ClientData.CreatedAt,
		&order.ClientData.UpdatedAt,

		&order.Price,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepo) GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error) {

	resp = &models.GetListOrderResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT
		COUNT(*) OVER(),
		o.id, 
		o.client_id, 

		c.id, 
		c.first_name,
		c.last_name,
		c.phone_number,
		CAST(c.created_at::timestamp AS VARCHAR),
		CAST(c.updated_at::timestamp AS VARCHAR),

		COALESCE(o.price, 0),
		COALESCE(o.status, ''),
		CAST(o.created_at::timestamp AS VARCHAR),
		CAST(o.updated_at::timestamp AS VARCHAR)
	FROM "orders" AS o
	JOIN client AS c ON c.id = o.client_id
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
		var order models.Order
		order.ClientData = &models.Client{}

		err = rows.Scan(
			&resp.Count,
			&order.Id,
			&order.ClientId,
			&order.ClientData.Id,
			&order.ClientData.FirstName,
			&order.ClientData.LastName,
			&order.ClientData.PhoneNumber,
			&order.ClientData.CreatedAt,
			&order.ClientData.UpdatedAt,

			&order.Price,
			&order.Status,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Orders = append(resp.Orders, &order)
	}

	return resp, nil
}

func (r *orderRepo) Update(ctx context.Context, req *models.UpdateOrder) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		orders
		SET
			id = :id, 
			client_id = :client_id, 
			price = :price,
			status = :status,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":        req.Id,
		"client_id": req.ClientId,
		"price":     req.Price,
		"status":    req.Status,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *orderRepo) Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM orders
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

// -------------ORDER_PRODUCTS-----------------------------------------------------------------------------------------------
func (r *orderRepo) AddOrderProduct(ctx context.Context, req *models.CreateOrderItem) (string, error) {
	id := uuid.NewString()

	query := `
		INSERT INTO order_products(
			id,
			order_id,
			product_id
		)
		VALUES (
			$1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.OrderId,
		req.ProductId,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *orderRepo) RemoveOrderItem(ctx context.Context, req *models.OrderProductPrimaryKey) (int64, error) {

	query := `
		DELETE FROM order_products WHERE id = $1
	`
	res, err := r.db.Exec(ctx, query, req.Id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
