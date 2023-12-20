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

type PSRepo struct {
	db *pgxpool.Pool
}

func NewPSRepoRepo(db *pgxpool.Pool) *PSRepo {
	return &PSRepo{
		db: db,
	}
}

func (r *PSRepo) Create(ctx context.Context, req models.CreatePickingSheet) (*models.PickingSheet, error) {

	var (
		Id                 = uuid.New().String()
		incrementIdNext, _ = helpers.NewIncrementId(r.db, "coming", "PS", 7)
		query              = `
			INSERT INTO "client"(
				"id",
				"increment_id",
				"product_id",
				"coming_id",
				"price",
				"quantity",
				"total",
				"updated_at"
			) VALUES ($1,$2,$3,$4,$5,$6,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		Id,
		incrementIdNext,
		req.ProductID,
		req.ComingTableID,
		req.Price,
		req.Quantity,
		req.Total,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.PickingSheetPrimaryKey{Id: Id})
}

func (c *PSRepo) GetById(ctx context.Context, req models.PickingSheetPrimaryKey) (*models.PickingSheet, error) {

	var (
		sheet = models.PickingSheet{}
		query = `
		SELECT 
				"id",
				"increment_id",
				"product_id",
				"coming_id",
				"price",
				"quantity",
				"total",
				"created_at",
				"updated_at"
		FROM picking_sheet WHERE id=$1`
	)

	var (
		Id            sql.NullString
		IncrementID   sql.NullString
		ProductID     sql.NullString
		ComingTableID sql.NullString
		Price         sql.NullFloat64
		Quantity      sql.NullInt64
		Total         sql.NullFloat64
		CreatedAt     sql.NullString
		UpdatedAt     sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&IncrementID,
		&ProductID,
		&ComingTableID,
		&Price,
		&Quantity,
		&Total,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	sheet = models.PickingSheet{
		Id:            Id.String,
		IncrementID:   IncrementID.String,
		ProductID:     ProductID.String,
		ComingTableID: ComingTableID.String,
		Price:         Price.Float64,
		Quantity:      Quantity.Int64,
		Total:         Total.Float64,
		CreatedAt:     CreatedAt.String,
		UpdatedAt:     UpdatedAt.String,
	}
	return &sheet, nil
}

func (r *PSRepo) GetList(ctx context.Context, req models.GetListPickingSheetRequest) (*models.GetListPickingSheetResponse, error) {
	var (
		response models.GetListPickingSheetResponse
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
							"product_id",
							"coming_id",
							"price",
							"quantity",
							"total",
							"created_at",
							"updated_at"
						FROM picking_sheet`
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
			ProductID   sql.NullString
			ComingID    sql.NullString
			Price       sql.NullFloat64
			Quantity    sql.NullInt64
			Total       sql.NullFloat64
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err := rows.Scan(
			&response.Count,
			&Id,
			&IncrementID,
			&ProductID,
			&ComingID,
			&Price,
			&Quantity,
			&Total,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		response.PickingSheets = append(response.PickingSheets, &models.PickingSheet{
			Id:            Id.String,
			IncrementID:   IncrementID.String,
			ProductID:     ProductID.String,
			ComingTableID: ComingID.String,
			Price:         Price.Float64,
			Quantity:      Quantity.Int64,
			Total:         Total.Float64,
			CreatedAt:     CreatedAt.String,
			UpdatedAt:     UpdatedAt.String,
		})
	}

	return &response, nil

}

func (r *PSRepo) Update(ctx context.Context, req models.UpdatePickingSheet) (*models.PickingSheet, error) {
	query := `
						UPDATE picking_sheet SET  
							"increment_id"=$2,
							"product_id"=$3,
							"coming_id"=$4,
							"price"=$5,
							"quantity"=$6,
							"total"=$7,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.IncrementID,
		req.ProductID,
		req.ComingTableID,
		req.Price,
		req.Quantity,
		req.Total,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.PickingSheetPrimaryKey{Id: req.Id})

}

func (r *PSRepo) Delete(ctx context.Context, req models.PickingSheetPrimaryKey) error {
	query := `DELETE FROM picking_sheet WHERE "id" = $1`
	// fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
