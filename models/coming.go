package models

type ComingTable struct {
	Id          string `json:"id"`
	IncrementID string `json:"increment_id"`
	Dated       string `json:"dated"`
	BranchID    string `json:"branch_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ComingTablePrimaryKey struct {
	Id string `json:"id"`
}

type CreateComingTable struct {
	IncrementID string `json:"-"`
	Dated       string `json:"dated"`
	BranchID    string `json:"branch_id"`
}

type UpdateComingTable struct {
	Id          string `json:"-"`
	IncrementID string `json:"-"`
	Dated       string `json:"dated"`
	BranchID    string `json:"branch_id"`
}

type GetListComingTableRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Query  string `json:"-"`
}

type GetListComingTableResponse struct {
	Count        int            `json:"count"`
	ComingTables []*ComingTable `json:"coming_tables"`
}
