package api

import (
	"clinics/api/handler"
	"clinics/config"
	"clinics/storage"

	"github.com/gin-gonic/gin"

	_ "clinics/api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpAPI(r *gin.Engine, cfg *config.Config, strg storage.StorageI) {
	handler := handler.NewHandler(cfg, strg)

	r.Use(customCORSMiddleware())

	//	@title		MARKET SYSTEM API
	//	@version	1.0
	//	@host		localhost:8080
	//	@BasePath	/api/v1
	v1 := r.Group("/api/v1")
	{
		// Branch
		v1.POST("/branch", handler.CreateBranch)
		v1.GET("/branch", handler.GetListBranch)
		v1.PUT("/branch/:id", handler.UpdateBranch)
		v1.GET("/branch/:id", handler.GetByIdBranch)
		v1.DELETE("/branch/:id", handler.DeleteBranch)

		// User
		v1.POST("/user", handler.CreateClient)
		v1.GET("/user", handler.GetListClient)
		v1.PUT("/user/:id", handler.UpdateClient)
		v1.GET("/user/:id", handler.GetByIdClient)
		v1.DELETE("/user/:id", handler.DeleteClient)

		// ComingTable
		v1.POST("/coming", handler.CreateComingTable)
		v1.GET("/coming", handler.GetListComingTable)
		v1.PUT("/coming/:id", handler.UpdateComingTable)
		v1.GET("/coming/:id", handler.GetByIdComingTable)
		v1.DELETE("/coming/:id", handler.DeleteComingTable)

		// Remainder
		v1.POST("/remainder", handler.CreateRemainder)
		v1.GET("/remainder", handler.GetListRemainder)
		v1.PUT("/remainder/:id", handler.UpdateRemainder)
		v1.GET("/remainder/:id", handler.GetByIdRemainder)
		v1.DELETE("/remainder/:id", handler.DeleteRemainder)

		// Picking Sheet
		v1.POST("/picking_sheet", handler.CreatePickingSheet)
		v1.GET("/picking_sheet", handler.GetListPickingSheet)
		v1.PUT("/picking_sheet/:id", handler.UpdatePickingSheet)
		v1.GET("/picking_sheet/:id", handler.GetByIdPickingSheet)
		v1.DELETE("/picking_sheet/:id", handler.DeletePickingSheet)

		// Product
		v1.POST("/product", handler.CreateProduct)
		v1.GET("/product", handler.GetListProduct)
		v1.PUT("/product/:id", handler.UpdateProduct)
		v1.GET("/product/:id", handler.GetByIdProduct)
		v1.DELETE("/product/:id", handler.DeleteProduct)

		//Sale Product
		v1.POST("/sale_product", handler.CreateSaleProduct)
		v1.GET("/sale_product", handler.GetListSaleProduct)
		v1.PUT("/sale_product/:id", handler.UpdateSaleProduct)
		v1.GET("/sale_product/:id", handler.GetByIdSaleProduct)
		v1.DELETE("/sale_product/:id", handler.DeleteSaleProduct)

		// Sale
		v1.POST("/sale", handler.CreateSale)
		v1.GET("/sale", handler.GetListSale)
		v1.PUT("/sale/:id", handler.UpdateSale)
		v1.GET("/sale/:id", handler.GetByIdSale)
		v1.DELETE("/sale/:id", handler.DeleteSale)

		//Overall Report
		v1.GET("/report", handler.OverallReport)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Password, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
