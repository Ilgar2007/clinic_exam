package postgres

import (
	"clinics/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(ctx context.Context, req models.CreateProduct) (*models.Product, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)

	var (
		productId = uuid.New().String()
		query     = `
			INSERT INTO "product"(
				"id",
				"name",
				"selling_price",
				"branch_id",
				"updated_at"
			) VALUES ($1,$2,$3,$4,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		productId,
		req.Name,
		req.SellingPrice,
		req.BranchID,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.ProductPrimaryKey{Id: productId})
}

func (c *ProductRepo) GetById(ctx context.Context, req models.ProductPrimaryKey) (*models.Product, error) {

	var (
		product = models.Product{}
		query   = `
		SELECT 
				"id",
				"name",
				"selling_price",
				"branch_id",
				"created_at",
				"updated_at"
		FROM "product" WHERE id=$1`
	)

	var (
		Id           sql.NullString
		Name         sql.NullString
		SellingPrice sql.NullFloat64
		BranchID     sql.NullString
		CreatedAt    sql.NullString
		UpdatedAt    sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Name,
		&SellingPrice,
		&BranchID,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	product = models.Product{
		Id:           Id.String,
		Name:         Name.String,
		SellingPrice: SellingPrice.Float64,
		BranchID:     BranchID.String,
		CreatedAt:    CreatedAt.String,
		UpdatedAt:    UpdatedAt.String,
	}
	return &product, nil
}

func (r *ProductRepo) GetList(ctx context.Context, req models.GetListProductRequest) (*models.GetListProductResponse, error) {

	var (
		response models.GetListProductResponse
		where    = " WHERE TRUE"
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		sort     = " ORDER BY created_at DESC"
		querySql string
		query    = `
		SELECT 
		COUNT(*) OVER(),
			"id",
			"name",
			"selling_price",
			"branch_id",
			"created_at",
			"updated_at"
		FROM "product"`
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	// if len(req.Search) > 0 {
	// 	where += " AND title ILIKE" + " '%" + req.Search + "%'"
	// }

	if len(req.Name) > 0 {
		where += " AND name ILIKE" + " '%" + req.Name + "%'"
	}

	if len(req.Query) > 0 {
		querySql = fmt.Sprintf(" AND %s", req.Query)
	}

	if len(req.BranchID) > 0 {
		where += " AND CAST(branch_id AS VARCHAR) ILIKE" + " '%" + req.BranchID + "%'"
	}

	query += where + querySql + sort + offset + limit
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			Id           sql.NullString
			Name         sql.NullString
			SellingPrice sql.NullFloat64
			BranchID     sql.NullString
			CreatedAt    sql.NullString
			UpdatedAt    sql.NullString
		)

		err := rows.Scan(
			&response.Count,
			&Id,
			&Name,
			&SellingPrice,
			&BranchID,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		response.Products = append(response.Products, &models.Product{
			Id:           Id.String,
			Name:         Name.String,
			SellingPrice: SellingPrice.Float64,
			BranchID:     BranchID.String,
			CreatedAt:    CreatedAt.String,
			UpdatedAt:    UpdatedAt.String,
		})
	}

	return &response, nil

}

func (r *ProductRepo) Update(ctx context.Context, req models.UpdateProduct) (*models.Product, error) {
	query := `
						UPDATE "product" SET  
							"name" = $2,
							"selling_price" = $3,
							"branch_id" = $4,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.SellingPrice,
		req.BranchID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.ProductPrimaryKey{Id: req.Id})

}

func (r *ProductRepo) Delete(ctx context.Context, req models.ProductPrimaryKey) error {
	query := `DELETE FROM "product" WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
