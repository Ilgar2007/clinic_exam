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

// CreateBranch godoc
// @ID				create_branch
// @Router			/branch [POST]
// @Summary		Create Branch
// @Description	Create Branch
// @Tags			Branch
// @Accept			json
// @Produce		json
// @Param			object	body		models.CreateBranch	true	"CreateBranchRequestBody"
// @Success		201		{object}	Response{data=models.Branch}	"BranchBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateBranch(c *gin.Context) {

	var branch models.CreateBranch
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(branch)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.Branch().Create(ctx, branch)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdBranch godoc
// @ID				get_by_id_branch
// @Router			/branch/{id} [GET]
// @Summary		GetById Branch
// @Description	GetById Branch
// @Tags			Branch
// @Accept			json
// @Produce		json
// @Param	 id path string		true	"GetByIdBranchRequestBody"
// @Success		200		{object}	Response{data=models.Branch}	"BranchBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdBranch(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.Branch().GetById(c, models.BranchPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListBranch godoc
// @ID				GetList_branch
// @Router		/branch [GET]
// @Summary		GetList Branch
// @Description	GetList Branch
// @Tags			Branch
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param name query string false "name"
// @Param phone_number query string false "phone_number"
// @Success		200		{object}	Response{data=models.GetListBranchResponse} "BranchBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListBranch(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	name := c.Query("name")
	phone_number := c.Query("phone_number")

	resp, err := h.strg.Branch().GetList(c, models.GetListBranchRequest{Offset: offset, Limit: limit, Name: name, PhoneNumber: phone_number})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateBranch godoc
// @ID update_baranch
// @Router /branch/{id} [PUT]
// @Summary Update Branch
// @Description Update Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "BranchPrimaryKey_ID"
// @Param object body models.UpdateBranch true "UpdateBranchBody"
// @Success  200 {object} Response{data=models.Branch} "Updated Branch"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateBranch(c *gin.Context) {

	var (
		baranch = models.UpdateBranch{}
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
	resp, err := h.strg.Branch().Update(c, baranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "Branch does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteBranch godoc
// @ID delete_baranch
// @Router /branch/{id} [DELETE]
// @Summary Delete Branch
// @Description Delete Branch
// @Tags Branch
// @Accept json
// @Produce json
// @Param id path string true "DeleteBranchPath"
// @Success  200 {object} Response{data=string} "Deleted Branch"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteBranch(c *gin.Context) {
	fmt.Println("SSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.Branch().Delete(c, models.BranchPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "Branch does not delete: "+err.Error())
		return
	}

	fmt.Println("DELETEDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
