package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *clientRepo {
	return &clientRepo{
		db: db,
	}
}

func (r *clientRepo) Create(ctx context.Context, req *models.CreateClient) (string, error) {
	var (
		query string
		id    string
	)
	id = uuid.NewString()

	query = `
		INSERT INTO client(
			id, 
			first_name,
			last_name,
			phone_number,
			updated_at 
		)
		VALUES ( $1, $2, $3, $4, now())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *clientRepo) GetByID(ctx context.Context, req *models.ClientPrimaryKey) (*models.Client, error) {

	var (
		query  string
		client models.Client
	)

	query = `
		SELECT
			id, 
			first_name,
			last_name,
			phone_number,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM client
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&client.Id,
		&client.FirstName,
		&client.LastName,
		&client.PhoneNumber,
		&client.CreatedAt,
		&client.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &client, nil
}

func (r *clientRepo) GetList(ctx context.Context, req *models.GetListClientRequest) (resp *models.GetListClientResponse, err error) {

	resp = &models.GetListClientResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, 
			first_name,
			last_name,
			phone_number,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM client
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
		var client models.Client
		err = rows.Scan(
			&resp.Count,
			&client.Id,
			&client.FirstName,
			&client.LastName,
			&client.PhoneNumber,
			&client.CreatedAt,
			&client.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Clients = append(resp.Clients, &client)
	}

	return resp, nil
}

func (r *clientRepo) Update(ctx context.Context, req *models.UpdateClient) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		client
		SET
			id = :id,
			first_name = :first_name,
			last_name = :last_name,
			phone_number = :phone_number,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *clientRepo) Delete(ctx context.Context, req *models.ClientPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM client
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
