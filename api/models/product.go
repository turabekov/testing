package models

type Product struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	CategoryId   string    `json:"category_id"`
	CategoryData *Category `json:"category_data"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Quantity     int       `json:"quantity"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}
type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateProduct struct {
	Name        string  `json:"name"`
	CategoryId  string  `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type UpdateProduct struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	CategoryId  string  `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	UpdatedAt   string  `json:"updated_at"`
}

type GetListProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
