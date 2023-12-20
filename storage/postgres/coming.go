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

type ComingTableRepo struct {
	db *pgxpool.Pool
}

func NewComingTableRepo(db *pgxpool.Pool) *ComingTableRepo {
	return &ComingTableRepo{
		db: db,
	}
}

func (r *ComingTableRepo) Create(ctx context.Context, req models.CreateComingTable) (*models.ComingTable, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)

	var (
		incrementIdNext, _ = helpers.NewIncrementId(r.db, "coming", "CD", 7)
		coming_tableId     = uuid.New().String()
		query              = `
			INSERT INTO "coming"(
				"id",
				"increment_id",
				"dated",
				"branch_id",
				"updated_at"
			) VALUES ($1,$2,$3,$4,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		coming_tableId,
		incrementIdNext(),
		req.Dated,
		req.BranchID,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.ComingTablePrimaryKey{Id: coming_tableId})
}

func (c *ComingTableRepo) GetById(ctx context.Context, req models.ComingTablePrimaryKey) (*models.ComingTable, error) {

	var (
		coming_table = models.ComingTable{}
		query        = `
		SELECT 
				"id",
				"increment_id",
				"dated",
				"branch_id",
				"created_at",
				"updated_at"
		FROM "coming" WHERE id=$1`
	)

	var (
		Id          sql.NullString
		IncrementID sql.NullString
		Dated       sql.NullString
		BranchID    sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&IncrementID,
		&Dated,
		&BranchID,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	coming_table = models.ComingTable{
		Id:          Id.String,
		IncrementID: IncrementID.String,
		Dated:       Dated.String,
		BranchID:    BranchID.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &coming_table, nil
}

func (r *ComingTableRepo) GetList(ctx context.Context, req models.GetListComingTableRequest) (*models.GetListComingTableResponse, error) {

	var (
		respons  models.GetListComingTableResponse
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
		"dated",
		"branch_id",
		"created_at",
		"updated_at"
		FROM "coming"`
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
			Dated       sql.NullString
			BranchID    sql.NullString
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err := rows.Scan(
			&respons.Count,
			&Id,
			&IncrementID,
			&Dated,
			&BranchID,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		respons.ComingTables = append(respons.ComingTables, &models.ComingTable{
			Id:          Id.String,
			IncrementID: IncrementID.String,
			Dated:       Dated.String,
			BranchID:    BranchID.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &respons, nil

}

func (r *ComingTableRepo) Update(ctx context.Context, req models.UpdateComingTable) (*models.ComingTable, error) {
	query := `
						UPDATE "coming" SET  
							"dated" = $2,
							"branch_id" = $3,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.Dated,
		req.BranchID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.ComingTablePrimaryKey{Id: req.Id})

}

func (r *ComingTableRepo) Delete(ctx context.Context, req models.ComingTablePrimaryKey) error {
	query := `DELETE FROM "coming" WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
