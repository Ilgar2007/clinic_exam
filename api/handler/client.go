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

// CreateUser godoc
// @ID				create_client
// @Router			/user [POST]
// @Summary		Create User
// @Description	Create User
// @Tags			User
// @Accept		json
// @Produce		json
// @Param			object	body		models.CreateClient	true	"CreateUserRequestBody"
// @Success		201		{object}	Response{data=models.Client}	"UserBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) CreateClient(c *gin.Context) {

	var client models.CreateClient
	err := c.ShouldBindJSON(&client)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	// fmt.Println(client)

	// fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&")
	resp, err := h.strg.Client().Create(ctx, client)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdUser godoc
// @ID				get_by_id_client
// @Router			/user/{id} [GET]
// @Summary		GetById User
// @Description	GetById User
// @Tags			User
// @Accept			json
// @Produce		json
// @Param	 id path string		true	"GetByIdUserRequestBody"
// @Success		200		{object}	Response{data=models.Client}	"UserBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetByIdClient(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.strg.Client().GetById(c, models.ClientPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListUser godoc
// @ID				get_list_client
// @Router		/user [GET]
// @Summary		GetList User
// @Description	GetList User
// @Tags			User
// @Accept		json
// @Produce		json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param first_name query string false "first_name"
// @Param last_name query string false "last_name"
// @Param branch_id query string false "branch_id"
// @Param phone_number query string false "phone_number"
// @Success		200		{object}	Response{data=models.GetListClientResponse} "UserBody"
// @Response	400		{object}	Response{data=string}	"Invalid Argument"
// @Failure		500		{object}	Response{data=string}
func (h *Handler) GetListClient(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	phone_number := c.Query("phone_number")
	branch_id := c.Query("branch_id")
	resp, err := h.strg.Client().GetList(c, models.GetListClientRequest{Offset: offset, Limit: limit, FirstName: first_name, LastName: last_name, BranchID: branch_id, PhoneNumber: phone_number})
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateUser godoc
// @ID update_client
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "UserPrimaryKey_ID"
// @Param object body models.UpdateClient true "UpdateUserBody"
// @Success  200 {object} Response{data=models.Client} "Updated User"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateClient(c *gin.Context) {

	var (
		baranch = models.UpdateClient{}
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
	resp, err := h.strg.Client().Update(c, baranch)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, "Client does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteUser godoc
// @ID delete_client
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "DeleteUserPath"
// @Success  200 {object} Response{data=string} "Deleted User"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure  500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteClient(c *gin.Context) {
	// fmt.Println("SSSSSSSSSSSSSSSSSSSSSSSSSSSSS")
	id := c.Param("id")
	fmt.Println(id)
	err := h.strg.Client().Delete(c, models.ClientPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, int(http.StatusInternalServerError), "Client does not delete: "+err.Error())
		return
	}

	// fmt.Println("DELETEDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")

	handleResponse(c, http.StatusAccepted, "SUCCESS DELETED")
}
