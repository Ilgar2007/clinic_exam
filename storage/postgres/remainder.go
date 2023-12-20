package postgres

import (
	"clinics/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RemainderRepo struct {
	db *pgxpool.Pool
}

func NewRemainderRepo(db *pgxpool.Pool) *RemainderRepo {
	return &RemainderRepo{
		db: db,
	}
}

func (r *RemainderRepo) Create(ctx context.Context, req models.CreateRemainder) (*models.Remainder, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)

	var (
		remainderId = uuid.New().String()
		query       = `
			INSERT INTO "remainder"(
				"id",
				"name",
				"quantity",
				"arrival_price",
				"selling_price",
				"branch_id",
				"updated_at"
			) VALUES ($1,$2,$3,$4,$5,$6,NOW())`
	)

	_, err := r.db.Exec(ctx, query,
		remainderId,
		req.Name,
		req.Quantity,
		req.ArrivalPrice,
		req.SellingPrice,
		req.BranchID,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.RemainderPrimaryKey{Id: remainderId})
}

func (c *RemainderRepo) GetById(ctx context.Context, req models.RemainderPrimaryKey) (*models.Remainder, error) {

	var (
		remainder = models.Remainder{}
		query     = `
		SELECT 
				"id",
				"name",
				"quantity",
				"arrival_price",
				"selling_price",
				"branch_id",
				"created_at",
				"updated_at"
		FROM "remainder" WHERE id=$1`
	)

	var (
		Id           sql.NullString
		Name         sql.NullString
		Quantity     sql.NullInt64
		ArrivalPrice sql.NullFloat64
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
		&Quantity,
		&ArrivalPrice,
		&SellingPrice,
		&BranchID,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	remainder = models.Remainder{
		Id:           Id.String,
		Name:         Name.String,
		Quantity:     Quantity.Int64,
		ArrivalPrice: ArrivalPrice.Float64,
		SellingPrice: SellingPrice.Float64,
		BranchID:     BranchID.String,
		CreatedAt:    CreatedAt.String,
		UpdatedAt:    UpdatedAt.String,
	}
	return &remainder, nil
}

func (r *RemainderRepo) GetList(ctx context.Context, req models.GetListRemainderRequest) (*models.GetListRemainderResponse, error) {

	var (
		response models.GetListRemainderResponse
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
			"quantity",
			"arrival_price",
			"selling_price",
			"branch_id",
			"created_at",
			"updated_at"
		FROM "remainder"`
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

	if len(req.BranchID) > 0 {
		where += " AND CAST(branch_id AS VARCHAR) ILIKE" + " '%" + req.BranchID + "%'"
	}

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
			Id           sql.NullString
			Name         sql.NullString
			Quantity     sql.NullInt64
			ArrivalPrice sql.NullFloat64
			SellingPrice sql.NullFloat64
			BranchID     sql.NullString
			CreatedAt    sql.NullString
			UpdatedAt    sql.NullString
		)

		err := rows.Scan(
			&response.Count,
			&Id,
			&Name,
			&Quantity,
			&ArrivalPrice,
			&SellingPrice,
			&BranchID,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		response.Remainders = append(response.Remainders, &models.Remainder{
			Id:           Id.String,
			Name:         Name.String,
			Quantity:     Quantity.Int64,
			ArrivalPrice: ArrivalPrice.Float64,
			SellingPrice: SellingPrice.Float64,
			BranchID:     BranchID.String,
			CreatedAt:    CreatedAt.String,
			UpdatedAt:    UpdatedAt.String,
		})
	}

	return &response, nil

}

func (r *RemainderRepo) Update(ctx context.Context, req models.UpdateRemainder) (*models.Remainder, error) {
	query := `
						UPDATE "remainder" SET  
							"name" = $2,
							"quantity" = $3,
							"arrival_price" = $4,
							"selling_price" = $5,
							"branch_id" = $6,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Quantity,
		req.ArrivalPrice,
		req.SellingPrice,
		req.BranchID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.RemainderPrimaryKey{Id: req.Id})

}

func (r *RemainderRepo) Delete(ctx context.Context, req models.RemainderPrimaryKey) error {
	query := `DELETE FROM "remainder" WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
