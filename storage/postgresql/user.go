package postgresql

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, req *models.CreateUser) (string, error) {
	var (
		query string
		id    string
	)
	id = uuid.NewString()

	query = `
		INSERT INTO users(
			id, 
			first_name,
			last_name,
			login,
			password,
			phone_number,
			updated_at 
		)
		VALUES ( $1, $2, $3, $4, $5, $6, now())
	`
	_, err := r.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.Login,
		req.Password,
		req.PhoneNumber,
	)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {

	var (
		query string
		user  models.User
	)

	if len(req.Login) > 0 {
		err := r.db.QueryRow(ctx, "SELECT id FROM users WHERE login = $1", req.Login).Scan(&req.Id)
		if err != nil {
			return nil, err
		}
	}

	query = `
		SELECT
			id, 
			first_name,
			last_name,
			login,
			password,
			phone_number,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM users
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Login,
		&user.Password,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error) {

	resp = &models.GetListUserResponse{}

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
			login,
			password,
			phone_number,
			CAST(created_at::timestamp AS VARCHAR),
			CAST(updated_at::timestamp AS VARCHAR)
		FROM users
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
		var user models.User
		err = rows.Scan(
			&resp.Count,
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Login,
			&user.Password,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return resp, nil
}

func (r *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		users
		SET
			id = :id,
			first_name = :first_name,
			last_name = :last_name,
			login = :login,
			password = :password,
			phone_number = :phone_number,
			updated_at = now()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"first_name":   req.FirstName,
		"last_name":    req.LastName,
		"login":        req.Login,
		"password":     req.Password,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM users
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}
