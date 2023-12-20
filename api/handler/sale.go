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

// CreateSale godoc
// @ID				create_sale
// @Router		/sale [POST]
// @Summary		Create Sale
// @Description	Create Sale
// @Tags			Sale
// @Accept		json
// @Produce		json
// @Param			object	body		models.CreateSale	true	"CreateSaleRequestBody"
// @Success		201		{object}	Response{data=models.Sale}	"SaleBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateSale(c *gin.Context) {

	var sale models.CreateSale
	err := c.ShouldBindJSON(&sale)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(sale)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.Sale().Create(ctx, sale)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdSale godoc
// @ID				get_by_id_sale
// @Router		/sale/{id} [GET]
// @Summary		GetById Sale
// @Description	GetById Sale
// @Tags			Sale
// @Accept		json
// @Produce		json
// @Param	 id path string		true	"GetByIdSaleRequestBody"
// @Success		200		{object}	Response{data=models.Sale}	"SaleBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdSale(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.Sale().GetById(c, models.SalePrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListSale godoc
// @ID				GetList_sale
// @Router		/sale [GET]
// @Summary		GetList Sale
// @Description	GetList Sale
// @Tags			Sale
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success		200		{object}	Response{data=models.GetListSaleResponse} "SaleBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListSale(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	resp, err := h.strg.Sale().GetList(c, models.GetListSaleRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateSale godoc
// @ID update_baranch
// @Router 			/sale/{id} [PUT]
// @Summary 		Update Sale
// @Description Update Sale
// @Tags 				Sale
// @Accept 			json
// @Produce 		json
// @Param id path string true "SalePrimaryKey_ID"
// @Param object body models.UpdateSale true "UpdateSaleBody"
// @Success  200 {object} Response{data=models.Sale} "Updated Sale"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateSale(c *gin.Context) {

	var (
		baranch = models.UpdateSale{}
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
	resp, err := h.strg.Sale().Update(c, baranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "Sale does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteSale godoc
// @ID delete_baranch
// @Router 			/sale/{id} [DELETE]
// @Summary	 		Delete Sale
// @Description Delete Sale
// @Tags 				Sale
// @Accept 			json
// @Produce 		json
// @Param id path string true "DeleteSalePath"
// @Success  200 {object} Response{data=string} "Deleted Sale"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteSale(c *gin.Context) {

	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.Sale().Delete(c, models.SalePrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "Sale does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
