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

// CreateComingTable godoc
// @ID				create_coming_table
// @Router			/coming [POST]
// @Summary			Create ComingTable
// @Description	Create ComingTable
// @Tags			ComingTable
// @Accept		json
// @Produce		json
// @Param			object	body		models.CreateComingTable	true	"CreateComingTableRequestBody"
// @Success		201		{object}	Response{data=models.ComingTable}	"ComingTableBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateComingTable(c *gin.Context) {

	var coming_table models.CreateComingTable
	err := c.ShouldBindJSON(&coming_table)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(coming_table)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.ComingTable().Create(ctx, coming_table)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdComingTable godoc
// @ID				get_by_id_coming_table
// @Router		/coming/{id} [GET]
// @Summary		GetById ComingTable
// @Description	GetById ComingTable
// @Tags			ComingTable
// @Accept			json
// @Produce		json
// @Param	 id path string		true	"GetByIdComingTableRequestBody"
// @Success		200		{object}	Response{data=models.ComingTable}	"ComingTableBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdComingTable(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.ComingTable().GetById(c, models.ComingTablePrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListComingTable godoc
// @ID				GetList_coming_table
// @Router		/coming [GET]
// @Summary		GetList ComingTable
// @Description	GetList ComingTable
// @Tags			ComingTable
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success		200		{object}	Response{data=models.GetListComingTableResponse} "ComingTableBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListComingTable(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	resp, err := h.strg.ComingTable().GetList(c, models.GetListComingTableRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateComingTable godoc
// @ID update_baranch
// @Router /coming/{id} [PUT]
// @Summary Update ComingTable
// @Description Update ComingTable
// @Tags ComingTable
// @Accept json
// @Produce json
// @Param id path string true "ComingTablePrimaryKey_ID"
// @Param object body models.UpdateComingTable true "UpdateComingTableBody"
// @Success  200 {object} Response{data=models.ComingTable} "Updated ComingTable"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateComingTable(c *gin.Context) {

	var (
		baranch = models.UpdateComingTable{}
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
	resp, err := h.strg.ComingTable().Update(c, baranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "ComingTable does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteComingTable godoc
// @ID delete_baranch
// @Router /coming/{id} [DELETE]
// @Summary Delete ComingTable
// @Description Delete ComingTable
// @Tags ComingTable
// @Accept json
// @Produce json
// @Param id path string true "DeleteComingTablePath"
// @Success  200 {object} Response{data=string} "Deleted ComingTable"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteComingTable(c *gin.Context) {
	// fmt.Println("SSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.ComingTable().Delete(c, models.ComingTablePrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "ComingTable does not delete: "+err.Error())
		return
	}

	// fmt.Println("DELETEDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
