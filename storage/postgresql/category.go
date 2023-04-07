package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (r *categoryRepo) Create(ctx context.Context, req *models.CreateCategory) (string, error) {
	var (
		query string
		id    string
	)
	id = uuid.NewString()

	query = `
		INSERT INTO category(
			id, 
			name,
			updated_at 
		)
		VALUES ( $1, $2, now())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *categoryRepo) GetByID(ctx context.Context, req *models.CategoryPrimaryKey) (*models.Category, error) {

	var (
		query    string
		category models.Category
	)

	query = `
		SELECT
			id,
			name, 
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM category
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepo) GetList(ctx context.Context, req *models.GetListCategoryRequest) (resp *models.GetListCategoryResponse, err error) {

	resp = &models.GetListCategoryResponse{}

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
			name, 
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM category
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
		var category models.Category
		err = rows.Scan(
			&resp.Count,
			&category.Id,
			&category.Name,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Categories = append(resp.Categories, &category)
	}

	return resp, nil
}

func (r *categoryRepo) Update(ctx context.Context, req *models.UpdateCategory) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		category
		SET
			id = :id,
			name = :name,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":   req.Id,
		"name": req.Name,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *categoryRepo) Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error) {
	query := `
		DELETE
		FROM category
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
