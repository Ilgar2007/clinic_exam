package models

type Sale struct {
	Id          string  `json:"id"`
	IncrementID string  `json:"increment_id"`
	ClientID    string  `json:"client_id"`
	BranchId    string  `json:"branch_id"`
	Total       float64 `json:"total"`
	Debt        float64 `json:"debt"`
	Paid        float64 `json:"paid"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type SalePrimaryKey struct {
	Id string `json:"id"`
}
type CreateSale struct {
	IncrementID string  `json:"-"`
	ClientID    string  `json:"client_id"`
	BranchId    string  `json:"branch_id"`
	Total       float64 `json:"total"`
	Debt        float64 `json:"debt"`
	Paid        float64 `json:"paid"`
}

type UpdateSale struct {
	Id          string  `json:"-"`
	IncrementID string  `json:"-"`
	ClientID    string  `json:"client_id"`
	BranchId    string  `json:"branch_id"`
	Total       float64 `json:"total"`
	Debt        float64 `json:"debt"`
	Paid        float64 `json:"paid"`
}

type GetListSaleRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"-"`
}

type GetListSaleResponse struct {
	Count int     `json:"count"`
	Sales []*Sale `json:"sales"`
}
