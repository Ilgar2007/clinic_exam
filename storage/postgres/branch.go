package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"clinics/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{
		db: db,
	}
}

func (r *BranchRepo) Create(ctx context.Context, req models.CreateBranch) (*models.Branch, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)
	var (
		branchId = uuid.New().String()
		query    = `
			INSERT INTO "branch"(
				"id",
				"name",
				"address",
				"phone_number",
				"updated_at"
			) VALUES ($1,$2,$3,$4,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		branchId,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.BranchPrimaryKey{Id: branchId})
}

func (c *BranchRepo) GetById(ctx context.Context, req models.BranchPrimaryKey) (*models.Branch, error) {

	var (
		branch = models.Branch{}
		query  = `
		SELECT 
				"id",
				"name",
				"address",
				"phone_number",
				"created_at",
				"updated_at"
		FROM branch WHERE id=$1`
	)

	var (
		Id          sql.NullString
		Name        sql.NullString
		Address     sql.NullString
		PhoneNumber sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&Name,
		&Address,
		&PhoneNumber,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	branch = models.Branch{
		Id:          Id.String,
		Name:        Name.String,
		Address:     Address.String,
		PhoneNumber: PhoneNumber.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &branch, nil
}

func (r *BranchRepo) GetList(ctx context.Context, req models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	var (
		respons  models.GetListBranchResponse
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
							"address",
							"phone_number",
							"created_at",
							"updated_at"
						FROM branch`
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

	if len(req.Name) > 0 {
		where += " AND name ILIKE" + " '%" + req.Name + "%'"
	}

	if len(req.PhoneNumber) > 0 {
		where += " AND phone_number ILIKE" + " '%" + req.PhoneNumber + "%'"
	}

	query += where + querySql + sort + offset + limit
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			Id          sql.NullString
			Name        sql.NullString
			Address     sql.NullString
			PhoneNumber sql.NullString
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err := rows.Scan(
			&respons.Count,
			&Id,
			&Name,
			&Address,
			&PhoneNumber,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		respons.Branches = append(respons.Branches, &models.Branch{
			Id:          Id.String,
			Name:        Name.String,
			Address:     Address.String,
			PhoneNumber: PhoneNumber.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &respons, nil

}

func (r *BranchRepo) Update(ctx context.Context, req models.UpdateBranch) (*models.Branch, error) {
	query := `
						UPDATE branch SET  
							"name" = $2,
							"address" = $3,
							"phone_number" = $4,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.BranchPrimaryKey{Id: req.Id})

}

func (r *BranchRepo) Delete(ctx context.Context, req models.BranchPrimaryKey) error {
	query := `DELETE FROM branch WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
