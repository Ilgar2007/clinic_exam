package storage

import (
	"clinics/models"
	"context"
)

type StorageI interface {
	Branch() BranchRepoI
	Client() ClientRepoI
	ComingTable() ComingTableRepoI
	PickingSheet() PickingSheetRepoI
	Product() ProductRepoI
	Remainder() RemainderRepoI
	SaleProduct() SaleProductRepoI
	Sale() SaleRepoI
	Report() ReportRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req models.CreateBranch) (*models.Branch, error)
	GetList(ctx context.Context, req models.GetListBranchRequest) (*models.GetListBranchResponse, error)
	GetById(ctx context.Context, req models.BranchPrimaryKey) (*models.Branch, error)
	Update(ctx context.Context, req models.UpdateBranch) (*models.Branch, error)
	Delete(ctx context.Context, req models.BranchPrimaryKey) error
}

type ClientRepoI interface {
	Create(ctx context.Context, req models.CreateClient) (*models.Client, error)
	GetList(ctx context.Context, req models.GetListClientRequest) (*models.GetListClientResponse, error)
	GetById(ctx context.Context, req models.ClientPrimaryKey) (*models.Client, error)
	Update(ctx context.Context, req models.UpdateClient) (*models.Client, error)
	Delete(ctx context.Context, req models.ClientPrimaryKey) error
}

type ComingTableRepoI interface {
	Create(ctx context.Context, req models.CreateComingTable) (*models.ComingTable, error)
	GetList(ctx context.Context, req models.GetListComingTableRequest) (*models.GetListComingTableResponse, error)
	GetById(ctx context.Context, req models.ComingTablePrimaryKey) (*models.ComingTable, error)
	Update(ctx context.Context, req models.UpdateComingTable) (*models.ComingTable, error)
	Delete(ctx context.Context, req models.ComingTablePrimaryKey) error
}

type PickingSheetRepoI interface {
	Create(ctx context.Context, req models.CreatePickingSheet) (*models.PickingSheet, error)
	GetList(ctx context.Context, req models.GetListPickingSheetRequest) (*models.GetListPickingSheetResponse, error)
	GetById(ctx context.Context, req models.PickingSheetPrimaryKey) (*models.PickingSheet, error)
	Update(ctx context.Context, req models.UpdatePickingSheet) (*models.PickingSheet, error)
	Delete(ctx context.Context, req models.PickingSheetPrimaryKey) error
}
type ProductRepoI interface {
	Create(ctx context.Context, req models.CreateProduct) (*models.Product, error)
	GetList(ctx context.Context, req models.GetListProductRequest) (*models.GetListProductResponse, error)
	GetById(ctx context.Context, req models.ProductPrimaryKey) (*models.Product, error)
	Update(ctx context.Context, req models.UpdateProduct) (*models.Product, error)
	Delete(ctx context.Context, req models.ProductPrimaryKey) error
}

type RemainderRepoI interface {
	Create(ctx context.Context, req models.CreateRemainder) (*models.Remainder, error)
	GetList(ctx context.Context, req models.GetListRemainderRequest) (*models.GetListRemainderResponse, error)
	GetById(ctx context.Context, req models.RemainderPrimaryKey) (*models.Remainder, error)
	Update(ctx context.Context, req models.UpdateRemainder) (*models.Remainder, error)
	Delete(ctx context.Context, req models.RemainderPrimaryKey) error
}

type SaleProductRepoI interface {
	Create(ctx context.Context, req models.CreateSaleProduct) (*models.SaleProduct, error)
	GetList(ctx context.Context, req models.GetListSaleProductRequest) (*models.GetListSaleProductResponse, error)
	GetById(ctx context.Context, req models.SaleProductPrimaryKey) (*models.SaleProduct, error)
	Update(ctx context.Context, req models.UpdateSaleProduct) (*models.SaleProduct, error)
	Delete(ctx context.Context, req models.SaleProductPrimaryKey) error
}

type SaleRepoI interface {
	Create(ctx context.Context, req models.CreateSale) (*models.Sale, error)
	GetList(ctx context.Context, req models.GetListSaleRequest) (*models.GetListSaleResponse, error)
	GetById(ctx context.Context, req models.SalePrimaryKey) (*models.Sale, error)
	Update(ctx context.Context, req models.UpdateSale) (*models.Sale, error)
	Delete(ctx context.Context, req models.SalePrimaryKey) error
}

type ReportRepoI interface {
	GetListReport(ctx context.Context) (*models.GetListClientResponse, error)
	GetListSaleBranch(ctx context.Context) (*models.AllSaleReport, error)
}
