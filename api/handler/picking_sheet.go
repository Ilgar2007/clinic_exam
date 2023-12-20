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

// CreatePickingSheet godoc
// @ID				create_picking_sheet
// @Router			/picking_sheet [POST]
// @Summary		Create PickingSheet
// @Description	Create PickingSheet
// @Tags			PickingSheet
// @Accept			json
// @Produce		json
// @Param			object	body		models.CreatePickingSheet	true	"CreatePickingSheetRequestBody"
// @Success		201		{object}	Response{data=models.PickingSheet}	"PickingSheetBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreatePickingSheet(c *gin.Context) {

	var picking_sheet models.CreatePickingSheet
	err := c.ShouldBindJSON(&picking_sheet)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.PickingSheet().Create(ctx, picking_sheet)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdPickingSheet godoc
// @ID				get_by_id_picking_sheet
// @Router			/picking_sheet/{id} [GET]
// @Summary		GetById PickingSheet
// @Description	GetById PickingSheet
// @Tags			PickingSheet
// @Accept			json
// @Produce		json
// @Param	 id path string		true	"GetByIdPickingSheetRequestBody"
// @Success		200		{object}	Response{data=models.PickingSheet}	"PickingSheetBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdPickingSheet(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.PickingSheet().GetById(c, models.PickingSheetPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListPickingSheet godoc
// @ID				GetList_picking_sheet
// @Router		/picking_sheet [GET]
// @Summary		GetList PickingSheet
// @Description	GetList PickingSheet
// @Tags			PickingSheet
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success		200		{object}	Response{data=models.GetListPickingSheetResponse} "PickingSheetBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListPickingSheet(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	resp, err := h.strg.PickingSheet().GetList(c, models.GetListPickingSheetRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdatePickingSheet godoc
// @ID update_picking_sheet
// @Router /picking_sheet/{id} [PUT]
// @Summary Update PickingSheet
// @Description Update PickingSheet
// @Tags PickingSheet
// @Accept json
// @Produce json
// @Param id path string true "PickingSheetPrimaryKey_ID"
// @Param object body models.UpdatePickingSheet true "UpdatePickingSheetBody"
// @Success  200 {object} Response{data=models.PickingSheet} "Updated PickingSheet"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePickingSheet(c *gin.Context) {

	var (
		picking_sheet = models.UpdatePickingSheet{}
		id            = c.Param("id")
	)
	err := c.ShouldBindJSON(&picking_sheet)
	fmt.Println(err)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	picking_sheet.Id = id
	fmt.Println(id)
	resp, err := h.strg.PickingSheet().Update(c, picking_sheet)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "PickingSheet does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeletePickingSheet godoc
// @ID delete_picking_sheet
// @Router /picking_sheet/{id} [DELETE]
// @Summary Delete PickingSheet
// @Description Delete PickingSheet
// @Tags PickingSheet
// @Accept json
// @Produce json
// @Param id path string true "DeletePickingSheetPath"
// @Success  200 {object} Response{data=string} "Deleted PickingSheet"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeletePickingSheet(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.PickingSheet().Delete(c, models.PickingSheetPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "PickingSheet does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
