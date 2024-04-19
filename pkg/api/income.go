package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lushenle/plam/pkg/db"
)

// createIncomeRequest is a struct that represents the request to create an income.
//
//	@swagger:model
type createIncomeRequest struct {
	// Payee of the income.
	// Required: true
	// example: john_doe
	// in: body
	Payee string `json:"payee" binding:"required"`

	// Amount of the income.
	// Required: true
	// example: 1000
	// in: body
	Amount float32 `json:"balance" binding:"required"`

	// ProjectID of the income.
	// Required: true
	// swagger:strfmt uuid
	// example: 123e4567-e89b-12d3-a456-426614174000
	// in: body
	ProjectID string `json:"project_id" binding:"required,uuid"`
}

// createIncome creates a new income.
//
//	@Summary		Create an income
//	@Description	Create a new income.
//	@Tags			incomes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createIncomeRequest	true	"Create Income Request"
//	@Success		200		{object}	db.Income			"Income created"
//	@Failure		400		{object}	errorResponse		"Bad Request"
//	@Failure		401		{object}	errorResponse		"Unauthorized"
//	@Failure		403		{object}	errorResponse		"Forbidden"
//	@Router			/incomes [post]
//	@security		ApiKeyAuth
func (server *Server) createIncome(ctx *gin.Context) {
	var req createIncomeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateIncomeParams{
		Payee:     req.Payee,
		Amount:    req.Amount,
		ProjectID: uuid.MustParse(req.ProjectID),
	}
	income, err := server.store.CreateIncome(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, income)
}

// listIncomes lists all incomes.
//
//	@Summary		List all incomes
//	@Description	List all incomes.
//	@Tags			incomes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		listRequest	true	"List Request"
//	@Success		200		{object}	[]db.Income
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/incomes/all [post]
//	@security		ApiKeyAuth
func (server *Server) listIncomes(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListIncomesParams{
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}

	incomes, err := server.store.ListIncomes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, incomes)
}

// getIncome gets an income by ID.
//
//	@Summary		Get an income
//	@Description	Get an income by ID.
//	@Tags			incomes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Income ID"
//	@Success		200	{object}	db.Income		"Income found"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/incomes/{id} [get]
//	@security		ApiKeyAuth
func (server *Server) getIncome(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	income, err := server.store.GetIncome(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, income)
}

// searchIncomes searches incomes by payee.
//
//	@Summary		Search incomes
//	@Description	Search incomes by payee.
//	@Tags			incomes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		searchRequest	true	"Search Request"
//	@Success		200		{object}	[]db.Income		"Incomes found"
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/incomes/search [post]
//	@security		ApiKeyAuth
func (server *Server) searchIncomes(ctx *gin.Context) {
	var req searchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.SearchIncomesParams{
		Column1: pgtype.Text{
			String: req.Query,
			Valid:  true,
		},
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}

	incomes, err := server.store.SearchIncomes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, incomes)
}

// deleteIncome deletes an income by ID.
//
//	@Summary		Delete an income
//	@Description	Delete an income by ID.
//	@Tags			incomes
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Income ID"
//	@Success		200	{object}	db.Income		"Deleted income"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Router			/incomes/{id} [delete]
//	@security		ApiKeyAuth
func (server *Server) deleteIncome(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	income, err := server.store.DeleteIncome(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, income)
}
