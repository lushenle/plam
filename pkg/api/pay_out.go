package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lushenle/plam/pkg/db"
)

// createPayOutRequest is a struct that represents the request to create a pay out.
//
//	@swagger:model
type createPayOutRequest struct {
	// Owner of the pay out.
	// Required: true
	// example: john_doe
	// in: body
	Owner string `json:"owner" binding:"required"`

	// Amount of the pay out.
	// Required: true
	// example: 1000
	// in: body
	Amount float32 `json:"amount" binding:"required"`

	// Subject of the pay out.
	// Required: true
	// example: pay_out1
	// in: body
	Subject string `json:"subject" binding:"required"`
}

// createPayOut creates a new pay out.
//
//	@Summary		Create a pay out
//	@Description	Create a new pay out.
//	@Tags			pay_outs
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createPayOutRequest	true	"Create Pay Out Request"
//	@Success		200		{object}	db.PayOut			"Pay Out created"
//	@Failure		400		{object}	errorResponse		"Bad Request"
//	@Failure		401		{object}	errorResponse		"Unauthorized"
//	@Failure		403		{object}	errorResponse		"Forbidden"
//	@Failure		500		{object}	errorResponse		"Internal Server Error"
//	@Router			/pay_outs [post]
//	@security		ApiKeyAuth
func (server *Server) createPayOut(ctx *gin.Context) {
	var req createPayOutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreatePayOutParams{
		Owner:   req.Owner,
		Amount:  req.Amount,
		Subject: req.Subject,
	}

	payOut, err := server.store.CreatePayOut(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrUniqueViolation) {
			ctx.JSON(http.StatusForbidden, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payOut)
}

// listPayOuts lists all pay outs.
//
//	@Summary		List all pay outs
//	@Description	List all pay outs.
//	@Tags			pay_outs
//	@Accept			json
//	@Produce		json
//	@Param			request	body		listRequest		true	"List Request"
//	@Success		200		{object}	[]db.PayOut		"List of pay outs"
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/pay_outs/all [post]
//	@security		ApiKeyAuth
func (server *Server) listPayOuts(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListPayOutsParams{
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}

	payOuts, err := server.store.ListPayOuts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payOuts)
}

// getPayOut gets a pay out by ID.
//
//	@Summary		Get a pay out
//	@Description	Get a pay out by ID.
//	@Tags			pay_outs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Pay Out ID"
//	@Success		200	{object}	db.PayOut		"Pay Out found"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/pay_outs/{id} [get]
//	@security		ApiKeyAuth
func (server *Server) getPayOut(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	payOut, err := server.store.GetPayOut(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payOut)
}

// searchPayOuts searches pay outs by owner.
//
//	@Summary		Search pay outs
//	@Description	Search pay outs by owner.
//	@Tags			pay_outs
//	@Accept			json
//	@Produce		json
//	@Param			request	body		searchRequest	true	"Search Request"
//	@Success		200		{object}	[]db.PayOut		"Pay Outs found"
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		404		{object}	errorResponse	"Not Found"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/pay_outs/search [post]
//	@security		ApiKeyAuth
func (server *Server) searchPayOuts(ctx *gin.Context) {
	var req searchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.SearchPayOutsParams{
		Owner:  req.Query,
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}

	payOuts, err := server.store.SearchPayOuts(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payOuts)
}

// deletePayOut deletes a pay out by ID.
//
//	@Summary		Delete a pay out
//	@Description	Delete a pay out by ID.
//	@Tags			pay_outs
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Pay Out ID"
//	@Success		200	{object}	db.PayOut		"Pay Out deleted"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/pay_outs/{id} [delete]
//	@security		ApiKeyAuth
func (server *Server) deletePayOut(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	payOut, err := server.store.DeletePayOut(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payOut)
}
