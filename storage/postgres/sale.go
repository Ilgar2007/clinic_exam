package postgres

import (
	"clinics/models"
	"clinics/pkg/helpers"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SaleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *SaleRepo {
	return &SaleRepo{
		db: db,
	}
}

func (r *SaleRepo) Create(ctx context.Context, req models.CreateSale) (*models.Sale, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)
	var (
		incremantIdNext, _ = helpers.NewIncrementId(r.db, "sale", "S", 7)
		saleId             = uuid.New().String()
		query              = `
			INSERT INTO "sale"(
				"id",
				"increment_id",
				"client_id",
				"branch_id",
				"total",
				"debt",
				"paid",
				"updated_at"
			) VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		saleId,
		incremantIdNext(),
		req.ClientID,
		req.BranchId,
		req.Total,
		req.Debt,
		req.Paid,
	)
	// fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.SalePrimaryKey{Id: saleId})
}

func (c *SaleRepo) GetById(ctx context.Context, req models.SalePrimaryKey) (*models.Sale, error) {

	var (
		sale  = models.Sale{}
		query = `
		SELECT 
				"id",
				"increment_id",
				"client_id",
				"branch_id",
				"total",
				"debt",
				"paid",
				"created_at",
				"updated_at"
		FROM "sale" WHERE id=$1`
	)

	var (
		Id          sql.NullString
		IncrementID sql.NullString
		ClientID    sql.NullString
		BranchId    sql.NullString
		Total       sql.NullFloat64
		Debt        sql.NullFloat64
		Paid        sql.NullFloat64
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&IncrementID,
		&ClientID,
		&BranchId,
		&Total,
		&Debt,
		&Paid,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	sale = models.Sale{
		Id:          Id.String,
		IncrementID: IncrementID.String,
		ClientID:    ClientID.String,
		BranchId:    BranchId.String,
		Total:       Total.Float64,
		Debt:        Debt.Float64,
		Paid:        Paid.Float64,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &sale, nil
}

func (r *SaleRepo) GetList(ctx context.Context, req models.GetListSaleRequest) (*models.GetListSaleResponse, error) {
	var (
		respons  models.GetListSaleResponse
		where    = " WHERE TRUE"
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		sort     = " ORDER BY created_at DESC"
		querySql string
		query    = `
						SELECT 
							COUNT(*) OVER(),
							"id",
							"increment_id",
							"client_id",
							"branch_id",
							"total",
							"debt",
							"paid",
							"created_at",
							"updated_at"
						FROM "sale"`
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	// if len(req.Search) > 0 {
	// 	where += " AND name ILIKE" + " '%" + req.Search + "%'"
	// }

	if len(req.Query) > 0 {
		querySql = fmt.Sprintf(" AND %s", req.Query)
	}

	query += where + querySql + sort + offset + limit
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			Id          sql.NullString
			IncrementID sql.NullString
			ClientID    sql.NullString
			BranchId    sql.NullString
			Total       sql.NullFloat64
			Debt        sql.NullFloat64
			Paid        sql.NullFloat64
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err := rows.Scan(
			&respons.Count,
			&Id,
			&IncrementID,
			&ClientID,
			&BranchId,
			&Total,
			&Debt,
			&Paid,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		respons.Sales = append(respons.Sales, &models.Sale{
			Id:          Id.String,
			IncrementID: IncrementID.String,
			ClientID:    ClientID.String,
			BranchId:    BranchId.String,
			Total:       Total.Float64,
			Debt:        Debt.Float64,
			Paid:        Paid.Float64,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &respons, nil

}

func (r *SaleRepo) Update(ctx context.Context, req models.UpdateSale) (*models.Sale, error) {
	query := `
						UPDATE "sale" SET  
							"client_id" =$2,
							"branch_id" =$3,
							"total" =$4,
							"debt" =$5,
							"paid" =$6,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.ClientID,
		req.BranchId,
		req.Total,
		req.Debt,
		req.Paid,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.SalePrimaryKey{Id: req.Id})

}

func (r *SaleRepo) Delete(ctx context.Context, req models.SalePrimaryKey) error {
	query := `DELETE FROM "sale" WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
