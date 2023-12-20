package handler

import (
	"clinics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"clinics/config"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @ID				create_product
// @Router			/product [POST]
// @Summary		Create Product
// @Description	Create Product
// @Tags			Product
// @Accept		json
// @Produce		json
// @Param			object	body		models.CreateProduct	true	"CreateProductRequestBody"
// @Success		201		{object}	Response{data=models.Product}	"ProductBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {

	var product models.CreateProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(Product)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.Product().Create(ctx, product)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdProduct godoc
// @ID				get_by_id_product
// @Router			/product/{id} [GET]
// @Summary		GetById Product
// @Description	GetById Product
// @Tags			Product
// @Accept			json
// @Produce		json
// @Param	 id path string		true	"GetByIdProductRequestBody"
// @Success		200		{object}	Response{data=models.Product}	"ProductBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdProduct(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.Product().GetById(c, models.ProductPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListProduct godoc
// @ID				get_list_product
// @Router		/product [GET]
// @Summary		GetList Product
// @Description	GetList Product
// @Tags			Product
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param name query string false "name"
// @Param branch_id query string false "branch_id"
// @Success		200		{object}	Response{data=models.GetListProductResponse} "ProductBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListProduct(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	name := c.Query("name")
	branch_id := c.Query("branch_id")

	resp, err := h.strg.Product().GetList(c, models.GetListProductRequest{Offset: offset, Limit: limit, Name: name, BranchID: branch_id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateProduct godoc
// @ID update_product
// @Router /product/{id} [PUT]
// @Summary Update Product
// @Description Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "ProductPrimaryKey_ID"
// @Param object body models.UpdateProduct true "UpdateProductBody"
// @Success  200 {object} Response{data=models.Client} "Updated Product"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	var (
		product = models.UpdateProduct{}
		id      = c.Param("id")
	)
	err := c.ShouldBindJSON(&product)
	fmt.Println(err)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	product.Id = id
	fmt.Println(id)
	resp, err := h.strg.Product().Update(c, product)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "Product does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteProduct godoc
// @ID delete_product
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "DeleteProductPath"
// @Success  200 {object} Response{data=string} "Deleted Product"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.Product().Delete(c, models.ProductPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "Product does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
