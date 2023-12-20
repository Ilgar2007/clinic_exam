package models

type Client struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	BirthDay    string `json:"birthday"`
	IsActive    bool   `json:"is_active"`
	Gender      string `json:"gender"`
	BranchID    string `json:"branch_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ClientPrimaryKey struct {
	Id string `json:"id"`
}
type CreateClient struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	BirthDay    string `json:"birthday"`
	IsActive    bool   `json:"is_active"`
	Gender      string `json:"gender"`
	BranchID    string `json:"branch_id"`
}

type UpdateClient struct {
	Id          string `json:"-"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	BirthDay    string `json:"birthday"`
	IsActive    bool   `json:"is_active"`
	Gender      string `json:"gender"`
	BranchID    string `json:"branch_id"`
}

type GetListClientRequest struct {
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BranchID    string `json:"branch_id"`
	PhoneNumber string `json:"phone_number"`
	Query       string `json:"-"`
}

type GetListClientResponse struct {
	Count   int       `json:"count"`
	Clients []*Client `json:"clients"`
}
