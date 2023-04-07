package models

type Order struct {
	Id            string          `json:"id"`
	ClientId      string          `json:"client_id"`
	ClientData    *Client         `json:"client_data"`
	Price         float64         `json:"price"`
	Status        string          `json:"status"`
	CreatedAt     string          `json:"created_at"`
	UpdatedAt     string          `json:"updated_at"`
	OrderProducts []*OrderProduct `json:"order_products"`
}

type OrderPrimaryKey struct {
	Id string `json:"id"`
}

type CreateOrder struct {
	ClientId  string  `json:"client_id"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type UpdateOrder struct {
	Id        string  `json:"id"`
	ClientId  string  `json:"client_id"`
	Price     float64 `json:"price"`
	Status    string  `json:"status"`
	UpdatedAt string  `json:"updated_at"`
}

type GetListOrderRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListOrderResponse struct {
	Count  int      `json:"count"`
	Orders []*Order `json:"orders"`
}

// -----------------------ITEM------------------
type OrderProduct struct {
	Id          string   `json:"id"`
	OrderId     int      `json:"order_id"`
	ProductId   int      `json:"product_id"`
	ProductData *Product `json:"product_data"`
}

type OrderProductPrimaryKey struct {
	Id string `json:"id"`
}

type CreateOrderItem struct {
	OrderId   string `json:"order_id"`
	ProductId string `json:"product_id"`
}
