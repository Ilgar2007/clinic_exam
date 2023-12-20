package models

type PickingSheet struct {
	Id            string  `json:"id"`
	IncrementID   string  `json:"increment_id"`
	ProductID     string  `json:"product_id"`
	ComingTableID string  `json:"coming_id"`
	Price         float64 `json:"price"`
	Quantity      int64   `json:"quantity"`
	Total         float64 `json:"total"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

type PickingSheetPrimaryKey struct {
	Id string `json:"id"`
}
type CreatePickingSheet struct {
	IncrementID   string  `json:"-"`
	ProductID     string  `json:"product_id"`
	ComingTableID string  `json:"coming_id"`
	Price         float64 `json:"price"`
	Quantity      int64   `json:"quantity"`
	Total         float64 `json:"total"`
}

type UpdatePickingSheet struct {
	Id            string  `json:"-"`
	IncrementID   string  `json:"-"`
	ProductID     string  `json:"product_id"`
	ComingTableID string  `json:"coming_id"`
	Price         float64 `json:"price"`
	Quantity      int64   `json:"quantity"`
	Total         float64 `json:"total"`
}

type GetListPickingSheetRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"-"`
}

type GetListPickingSheetResponse struct {
	Count         int             `json:"count"`
	PickingSheets []*PickingSheet `json:"picking_sheets"`
}
