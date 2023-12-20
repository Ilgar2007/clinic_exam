package models

type Branch struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}
type CreateBranch struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateBranch struct {
	Id          string `json:"-"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type GetListBranchRequest struct {
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Query       string `json:"-"`
}

type GetListBranchResponse struct {
	Count    int       `json:"count"`
	Branches []*Branch `json:"branches"`
}
