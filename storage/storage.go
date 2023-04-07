package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Category() CategoryRepoI
	Client() ClientRepoI
	Order() OrderRepoI
	User() UserRepoI
}
type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (string, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetListUserResponse, err error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) (int64, error)
}

type ProductRepoI interface {
	Create(context.Context, *models.CreateProduct) (string, error)
	GetByID(context.Context, *models.ProductPrimaryKey) (*models.Product, error)
	GetList(context.Context, *models.GetListProductRequest) (*models.GetListProductResponse, error)
	Update(ctx context.Context, req *models.UpdateProduct) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) (int64, error)
}

type CategoryRepoI interface {
	Create(context.Context, *models.CreateCategory) (string, error)
	GetByID(context.Context, *models.CategoryPrimaryKey) (*models.Category, error)
	GetList(context.Context, *models.GetListCategoryRequest) (*models.GetListCategoryResponse, error)
	Delete(ctx context.Context, req *models.CategoryPrimaryKey) (int64, error)
	Update(ctx context.Context, req *models.UpdateCategory) (int64, error)
}

type ClientRepoI interface {
	Create(ctx context.Context, req *models.CreateClient) (string, error)
	GetByID(ctx context.Context, req *models.ClientPrimaryKey) (*models.Client, error)
	GetList(ctx context.Context, req *models.GetListClientRequest) (resp *models.GetListClientResponse, err error)
	Update(ctx context.Context, req *models.UpdateClient) (int64, error)
	Delete(ctx context.Context, req *models.ClientPrimaryKey) (int64, error)
}

type OrderRepoI interface {
	Create(ctx context.Context, req *models.CreateOrder) (string, error)
	GetByID(ctx context.Context, req *models.OrderPrimaryKey) (*models.Order, error)
	GetList(ctx context.Context, req *models.GetListOrderRequest) (resp *models.GetListOrderResponse, err error)
	Update(ctx context.Context, req *models.UpdateOrder) (int64, error)
	Delete(ctx context.Context, req *models.OrderPrimaryKey) (int64, error)
	AddOrderProduct(ctx context.Context, req *models.CreateOrderItem) (string, error)
	RemoveOrderItem(ctx context.Context, req *models.OrderProductPrimaryKey) (int64, error)
}
