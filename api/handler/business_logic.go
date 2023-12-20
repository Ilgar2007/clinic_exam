package handler

import (
	"net/http"

	"clinics/models"

	"github.com/gin-gonic/gin"
)

// OverallReport godoc
// @ID				OverallReport
// @Router		/report [GET]
// @Summary		OverallReport
// @Description	OverallReport
// @Tags			OverallReport
// @Accept		json
// @Produce		json
// @Success		201		{object}	Response{data=models.Client}	"OverallReportBody"
// @Response	400		{object}	Response{data=string}		"Invalid Argument"
// @Failure		500		{object}	Response{data=string}	"Server Error"
func (h *Handler) OverallReport(c *gin.Context) {
	clent_resp, err := h.strg.Report().GetListReport(c)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err)
	}
	sale_resp, err := h.strg.Report().GetListSaleBranch(c)
	handleResponse(c, http.StatusOK, models.OverallReport{
		Clients:           clent_resp.Clients,
		BranchSaleReports: *sale_resp,
	})
}
