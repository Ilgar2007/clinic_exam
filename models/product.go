package models

type Product struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	SellingPrice float64 `json:"selling_price"`
	BranchID     string  `json:"branch_id"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type ProductPrimaryKey struct {
	Id string `json:"id"`
}
type CreateProduct struct {
	Name         string  `json:"name"`
	SellingPrice float64 `json:"selling_price"`
	BranchID     string  `json:"branch_id"`
}

type UpdateProduct struct {
	Id           string  `json:"-"`
	Name         string  `json:"name"`
	SellingPrice float64 `json:"selling_price"`
	BranchID     string  `json:"branch_id"`
}

type GetListProductRequest struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Name     string `json:"name"`
	BranchID string `json:"branch_id"`
	Query    string `json:"-"`
}

type GetListProductResponse struct {
	Count    int        `json:"count"`
	Products []*Product `json:"products"`
}
