package models

type SaleProduct struct {
	Id          string  `json:"id"`
	IncrementID string  `json:"increment_id"`
	ProductID   string  `json:"product_id"`
	SaleID      string  `json:"sale_id"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	Total       float64 `json:"total"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type SaleProductPrimaryKey struct {
	Id string `json:"id"`
}
type CreateSaleProduct struct {
	IncrementID string  `json:"-"`
	ProductID   string  `json:"product_id"`
	SaleID      string  `json:"sale_id"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	Total       float64 `json:"total"`
}

type UpdateSaleProduct struct {
	Id          string  `json:"-"`
	IncrementID string  `json:"-"`
	ProductID   string  `json:"product_id"`
	SaleID      string  `json:"sale_id"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	Total       float64 `json:"total"`
}

type GetListSaleProductRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"-"`
}

type GetListSaleProductResponse struct {
	Count        int            `json:"count"`
	SaleProducts []*SaleProduct `json:"saleproducts"`
}
