package postgres

import (
	"clinics/models"
	"clinics/pkg/helpers"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SaleProductRepo struct {
	db *pgxpool.Pool
}

func NewSaleProductRepo(db *pgxpool.Pool) *SaleProductRepo {
	return &SaleProductRepo{
		db: db,
	}
}

func (r *SaleProductRepo) Create(ctx context.Context, req models.CreateSaleProduct) (*models.SaleProduct, error) {
	// fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@", req)

	var (
		incremantIdNext, _ = helpers.NewIncrementId(r.db, "sale_product", "SP", 7)
		sale_productId     = uuid.New().String()
		query              = `
			INSERT INTO "sale_product"(
				"id",
				"increment_id",
				"product_id",
				"sale_id",
				"price",
				"quantity",
				"total",
				"updated_at"
			) VALUES ($1,$2,$3,$4,$5,$6,$7,NOW())`
	)
	/********************************************** task 1 **********************************************/
	// 1. Продажада товар сотилетганда агар остаток да товар йетарли булмаса, товар сотилмаслиги ва продажа создать булмаслиги кк
	var (
		Sql = `SELECT 
		R.id,
		R.quantity,     
		R.selling_price 
		FROM "sale" AS S
		LEFT JOIN "remainder" AS R ON R.branch_id = S.branch_id
		WHERE S.id = $1
;
		`
		updateSql = `
				UPDATE "remainder" SET  
						"quantity" = $2,
						"updated_at" = NOW()
				WHERE "id" = $1`

		Id sql.NullString
		// Name         sql.NullString
		Quantity sql.NullInt64
		// ArrivalPrice sql.NullFloat64
		SellingPrice sql.NullFloat64
		// BranchID sql.NullString
	)

	rowRemaninder := r.db.QueryRow(ctx, Sql, req.SaleID)
	rowRemaninder.Scan(
		&Id,
		// &Name,
		&Quantity,
		// &ArrivalPrice,
		&SellingPrice,
		// &BranchID,
	)
	fmt.Println(Quantity.Int64, req.Quantity)
	fmt.Println(Id.String, Quantity.Int64, SellingPrice.Float64)

	if req.Quantity > Quantity.Int64 {
		fmt.Println(req.Quantity)
		fmt.Println(Quantity.Int64)
		return nil, errors.New("There is not enough product")
	}

	/********************************************** task 1 end **********************************************/
	/********************************************** task 2 **********************************************/
	// Продажада товар сотиб олганида агар толамокчи болган пули общий толаши кк булган пулни 50 % дан кам булса, продажа булмаслиги кк
	fmt.Println(float64(req.Quantity) * SellingPrice.Float64 / 2)
	if req.Price < float64(req.Quantity)*SellingPrice.Float64/2 {
		fmt.Println(req.Price)
		return nil, errors.New("Not enough money")
	}
	/********************************************** task 2 end **********************************************/
	upQuantity := Quantity.Int64 - req.Quantity
	_, err := r.db.Exec(ctx, updateSql,
		Id,
		upQuantity,
	)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Exec(ctx, query,
		sale_productId,
		incremantIdNext(),
		req.ProductID,
		req.SaleID,
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
	return r.GetById(ctx, models.SaleProductPrimaryKey{Id: sale_productId})
}

func (c *SaleProductRepo) GetById(ctx context.Context, req models.SaleProductPrimaryKey) (*models.SaleProduct, error) {

	var (
		sale_product = models.SaleProduct{}
		query        = `
		SELECT 
				"id",
				"increment_id",
				"product_id",
				"sale_id",
				"price",
				"quantity",
				"total",
				"created_at",
				"updated_at"
		FROM "sale_product" WHERE id=$1`
	)

	var (
		Id          sql.NullString
		IncrementID sql.NullString
		ProductID   sql.NullString
		SaleID      sql.NullString
		Price       sql.NullFloat64
		Quantity    sql.NullInt64
		Total       sql.NullFloat64
		CreatedAt   sql.NullString
		UpdatedAt   sql.NullString
	)
	// fmt.Println(query)
	resp := c.db.QueryRow(ctx, query, req.Id)
	// fmt.Println("*********************", resp)

	err := resp.Scan(
		&Id,
		&IncrementID,
		&ProductID,
		&SaleID,
		&Price,
		&Quantity,
		&Total,
		&CreatedAt,
		&UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	sale_product = models.SaleProduct{
		Id:          Id.String,
		IncrementID: IncrementID.String,
		ProductID:   ProductID.String,
		SaleID:      SaleID.String,
		Price:       Price.Float64,
		Quantity:    Quantity.Int64,
		Total:       Total.Float64,
		CreatedAt:   CreatedAt.String,
		UpdatedAt:   UpdatedAt.String,
	}
	return &sale_product, nil
}

func (r *SaleProductRepo) GetList(ctx context.Context, req models.GetListSaleProductRequest) (*models.GetListSaleProductResponse, error) {
	var (
		respons  models.GetListSaleProductResponse
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
							"sale_id",
							"price",
							"quantity",
							"total",
							"created_at",
							"updated_at"
						FROM "sale_product"`
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if len(req.Search) > 0 {
		where += " AND increment_id ILIKE" + " '%" + req.Search + "%'"
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
			Id          sql.NullString
			IncrementID sql.NullString
			ProductID   sql.NullString
			SaleID      sql.NullString
			Price       sql.NullFloat64
			Quantity    sql.NullInt64
			Total       sql.NullFloat64
			CreatedAt   sql.NullString
			UpdatedAt   sql.NullString
		)

		err := rows.Scan(
			&respons.Count,
			&Id,
			&IncrementID,
			&ProductID,
			&SaleID,
			&Price,
			&Quantity,
			&Total,
			&CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		respons.SaleProducts = append(respons.SaleProducts, &models.SaleProduct{
			Id:          Id.String,
			IncrementID: IncrementID.String,
			ProductID:   ProductID.String,
			SaleID:      SaleID.String,
			Price:       Price.Float64,
			Quantity:    Quantity.Int64,
			Total:       Total.Float64,
			CreatedAt:   CreatedAt.String,
			UpdatedAt:   UpdatedAt.String,
		})
	}

	return &respons, nil

}

func (r *SaleProductRepo) Update(ctx context.Context, req models.UpdateSaleProduct) (*models.SaleProduct, error) {
	query := `
						UPDATE "sale_product" SET  
							"product_id" = $2,
							"sale_id" = $3,
							"price" = $4,
							"quantity" = $5,
							"total" = $6,
							"updated_at" = NOW()
						WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query,
		req.Id,
		req.ProductID,
		req.SaleID,
		req.Price,
		req.Quantity,
		req.Total,
	)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, models.SaleProductPrimaryKey{Id: req.Id})

}

func (r *SaleProductRepo) Delete(ctx context.Context, req models.SaleProductPrimaryKey) error {
	query := `DELETE FROM "sale_product" WHERE "id" = $1`
	fmt.Println(query)
	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil

}
