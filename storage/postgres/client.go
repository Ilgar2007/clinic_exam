package postgres

import (
	"clinics/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) *ClientRepo {
	return &ClientRepo{
		db: db,
	}
}

func (r *ClientRepo) Create(ctx context.Context, req models.CreateClient) (*models.Client, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)
	var (
		clientId = uuid.New().String()
		query    = `
			INSERT INTO "client"(
				"id",
				"first_name",
				"last_name",
				"phone_number",
				"birthday",
				"is_active",
				"gender",
				"branch_id",
				"updated_at"
			) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW())`
	)
	_, err := r.db.Exec(ctx, query,
		clientId,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.BirthDay,
		req.IsActive,
		req.Gender,
		req.BranchID,
	)
	fmt.Println(query)
	// defer r.db.Close()
	if err != nil {
		return nil, err
	}

	// fmt.Println("CREATED")
	return r.GetById(ctx, models.ClientPrimaryKey{Id: clientId})
}

func (c *ClientRepo) GetById(ctx context.Context, req models.ClientPrimaryKey) (*models.Client, error) {

	var (
		client = models.Client{}
		query  = `
		SELECT 
				"id",
				"first_name",
				"last_name",
				"phone_number",
				"birthday",
				"is_active",
				"gender",
				"branch_id",
				"created_at",
				"updated_at"
		FROM client WHERE id=$1`
	)

	var (
		Id          sql.NullString
		FirstName   sql.NullString
		LastName    sql.NullString
		PhoneNumber sql.NullString
		BirthDay    sql.NullString
		IsActive    sql.NullBool
		Gender      sql.NullString
		BranchID    sql.NullString
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&FirstName,
		&LastName,
		&PhoneNumber,
		&BirthDay,
		&IsActive,
		&Gender,
		&BranchID,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	client = models.Client{
		Id:          Id.String,
		FirstName:   FirstName.String,
		LastName:    LastName.String,
		PhoneNumber: PhoneNumber.String,
		BirthDay:    BirthDay.String,
		IsActive:    IsActive.Bool,
		Gender:      Gender.String,
		BranchID:    BranchID.String,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &client, nil
}

func (r *ClientRepo) GetList(ctx context.Context, req models.GetListClientRequest) (*models.GetListClientResponse, error) {
	var (
		respons  models.GetListClientResponse
		where    = " WHERE TRUE"
		offset   = " OFFSET 0"
		limit    = " LIMIT 10"
		sort     = " ORDER BY created_at DESC"
		querySql string
		query    = `
						SELECT 
							COUNT(*) OVER(),
							"id",
							"first_name",
							"last_name",
							"phone_number",
							"birthday",
							"is_active",
							"gender",
							"branch_id",
							"created_at",
							"updated_at"
						FROM client`
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

	if len(req.FirstName) > 0 {
		where += " AND first_name ILIKE" + " '%" + req.FirstName + "%'"
	}

	if len(req.LastName) > 0 {
		where += " AND last_name ILIKE" + " '%" + req.LastName + "%'"
	}

	if len(req.PhoneNumber) > 0 {
		where += " AND phone_number ILIKE" + " '%" + req.PhoneNumber + "%'"
	}

	if len(req.BranchID) > 0 {
		where += " AND CAST(branch_id AS VARCHAR) ILIKE" + " '%" + req.BranchID + "%'"
	}

	if len(req.Query) > 0 {
		querySql = fmt.Sprintf(" AND %s", req.Query)
	}

	query += where + querySql + sort + offset + limit
	fmt.Println(query)
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			Id          sql.NullString
			FirstName   sql.NullString
			LastName    sql.NullString
			PhoneNumber sql.NullString
			BirthDay    sql.NullString
			IsActive    sql.NullBool
			Gender      sql.NullString
			BranchID    sql.NullString
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err := rows.Scan(
			&respons.Count,
			&Id,
			&FirstName,
			&LastName,
			&PhoneNumber,
			&BirthDay,
			&IsActive,
			&Gender,
			&BranchID,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		respons.Clients = append(respons.Clients, &models.Client{
			Id:          Id.String,
			FirstName:   FirstName.String,
			LastName:    LastName.String,
			PhoneNumber: PhoneNumber.String,
			BirthDay:    BirthDay.String,
			IsActive:    IsActive.Bool,
			Gender:      Gender.String,
			BranchID:    BranchID.String,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &respons, nil

}

func (r *ClientRepo) Update(ctx context.Context, req models.UpdateClient) (*models.Client, error) {
	query := `
						UPDATE client SET  
							"first_name" = $2,
							"last_name" = $3,
							"phone_number" = $4,
							"birthday" = $5,
							"is_active" = $6,
							"gender" = $7,
							"branch_id" = $8,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.BirthDay,
		req.IsActive,
		req.Gender,
		req.BranchID,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.ClientPrimaryKey{Id: req.Id})

}

func (r *ClientRepo) Delete(ctx context.Context, req models.ClientPrimaryKey) error {
	query := `DELETE FROM client WHERE "id" = $1`
	// fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
