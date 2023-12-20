package postgres

import (
	"clinics/models"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ReportRepo struct {
	db *pgxpool.Pool
}

func NewReportRepo(db *pgxpool.Pool) *ReportRepo {
	return &ReportRepo{
		db: db,
	}
}

func (r *ReportRepo) GetListReport(ctx context.Context) (*models.GetListClientResponse, error) {
	var (
		respons models.GetListClientResponse
		query   = `
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
						FROM client ORDER BY created_at DESC`
	)

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

func (r *ReportRepo) GetListSaleBranch(ctx context.Context) (*models.AllSaleReport, error) {
	var (
		respons models.AllSaleReport
		query   = `
				SELECT 
					B.name,
					SUM(SP.quantity)AS quantity,
					SUM("price") AS price
				FROM "sale_product" AS SP
				JOIN "sale" AS S ON S.id = SP.sale_id
				JOIN "branch" AS B ON B.id = S.branch_id
				GROUP BY B.name;`
	)

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			Name     sql.NullString
			Quantity sql.NullInt64
			Price    sql.NullFloat64
		)

		err := rows.Scan(
			&Name,
			&Quantity,
			&Price,
		)
		if err != nil {
			return nil, err
		}

		respons.SaleReports = append(respons.SaleReports, &models.SaleReport{
			Name:     Name.String,
			Quantity: Quantity.Int64,
			Price:    Price.Float64,
		})
	}

	return &respons, nil

}
