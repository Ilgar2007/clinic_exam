package models

type Remainder struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Quantity     int64   `json:"quantity"`
	ArrivalPrice float64 `json:"arrival_price"`
	SellingPrice float64 `json:"selling_price"`
	BranchID     string  `json:"branch_id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type RemainderPrimaryKey struct {
	Id string `json:"id"`
}
type CreateRemainder struct {
	Name         string  `json:"name"`
	Quantity     int64   `json:"quantity"`
	ArrivalPrice float64 `json:"arrival_price"`
	SellingPrice float64 `json:"selling_price"`
	BranchID     string  `json:"branch_id"`
}

type UpdateRemainder struct {
	Id           string  `json:"-"`
	Name         string  `json:"name"`
	Quantity     int64   `json:"quantity"`
	ArrivalPrice float64 `json:"arrival_price"`
	SellingPrice float64 `json:"selling_price"`
	BranchID     string  `json:"branch_id"`
}

type GetListRemainderRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	BranchID string `json:"branch_id"`
	Name     string `json:"name"`
	Query    string `json:"-"`
}

type GetListRemainderResponse struct {
	Count      int          `json:"count"`
	Remainders []*Remainder `json:"remainders"`
}
