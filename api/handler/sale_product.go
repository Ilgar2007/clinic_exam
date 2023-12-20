package handler

import (
	"clinics/config"
	"clinics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSaleProduct godoc
// @ID				create_sale_product
// @Router		/sale_product [POST]
// @Summary		Create SaleProduct
// @Description	Create SaleProduct
// @Tags			SaleProduct
// @Accept		json
// @Produce		json
// @Param			object	body		models.CreateSaleProduct	true	"CreateSaleProductRequestBody"
// @Success		201		{object}	Response{data=models.SaleProduct}	"SaleProductBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateSaleProduct(c *gin.Context) {

	var sale_product models.CreateSaleProduct
	err := c.ShouldBindJSON(&sale_product)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(sale_product)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.SaleProduct().Create(ctx, sale_product)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdSaleProduct godoc
// @ID				get_by_id_sale_product
// @Router		/sale_product/{id} [GET]
// @Summary		GetById SaleProduct
// @Description	GetById SaleProduct
// @Tags			SaleProduct
// @Accept		json
// @Produce		json
// @Param	 id path string		true	"GetByIdSaleProductRequestBody"
// @Success		200		{object}	Response{data=models.SaleProduct}	"SaleProductBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdSaleProduct(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.SaleProduct().GetById(c, models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListSaleProduct godoc
// @ID				GetList_sale_product
// @Router		/sale_product [GET]
// @Summary		GetList SaleProduct
// @Description	GetList SaleProduct
// @Tags			SaleProduct
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success		200		{object}	Response{data=models.GetListSaleProductResponse} "SaleProductBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListSaleProduct(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	resp, err := h.strg.SaleProduct().GetList(c, models.GetListSaleProductRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateSaleProduct godoc
// @ID update_baranch
// @Router 			/sale_product/{id} [PUT]
// @Summary 		Update SaleProduct
// @Description Update SaleProduct
// @Tags 				SaleProduct
// @Accept 			json
// @Produce 		json
// @Param id path string true "SaleProductPrimaryKey_ID"
// @Param object body models.UpdateSaleProduct true "UpdateSaleProductBody"
// @Success  200 {object} Response{data=models.SaleProduct} "Updated SaleProduct"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateSaleProduct(c *gin.Context) {

	var (
		baranch = models.UpdateSaleProduct{}
		id      = c.Param("id")
	)
	err := c.ShouldBindJSON(&baranch)
	fmt.Println(err)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	baranch.Id = id
	fmt.Println(id)
	resp, err := h.strg.SaleProduct().Update(c, baranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "SaleProduct does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteSaleProduct godoc
// @ID delete_baranch
// @Router 			/sale_product/{id} [DELETE]
// @Summary	 		Delete SaleProduct
// @Description Delete SaleProduct
// @Tags 				SaleProduct
// @Accept 			json
// @Produce 		json
// @Param id path string true "DeleteSaleProductPath"
// @Success  200 {object} Response{data=string} "Deleted SaleProduct"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteSaleProduct(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.SaleProduct().Delete(c, models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "SaleProduct does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
