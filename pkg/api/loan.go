package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lushenle/plam/pkg/db"
)

// createLoanRequest is a struct that represents the request to create a loan.
//
//	@swagger:model
type createLoanRequest struct {
	// Borrower of the loan.
	// Required: true
	// example: john_doe
	// in: body
	Borrower string `json:"borrower" binding:"required"`

	// Subject of the loan.
	// Required: true
	// example: loan1
	// in: body
	Subject string `json:"subject" binding:"required"`

	// Amount of the loan.
	// Required: true
	// example: 1000
	// in: body
	Amount float32 `json:"amount" binding:"required"`
}

// createLoan creates a new loan.
//
//	@Summary		Create a loan
//	@Description	Create a new loan.
//	@Tags			loans
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createLoanRequest	true	"Create Loan Request"
//	@Success		200		{object}	db.Loan				"Loan created"
//	@Failure		400		{object}	errorResponse		"Bad Request"
//	@Failure		401		{object}	errorResponse		"Unauthorized"
//	@Failure		403		{object}	errorResponse		"Forbidden"
//	@Router			/loans [post]
//	@security		ApiKeyAuth
func (server *Server) createLoan(ctx *gin.Context) {
	var req createLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateLoanParams{
		Borrower: req.Borrower,
		Subject:  req.Subject,
		Amount:   req.Amount,
	}

	loan, err := server.store.CreateLoan(ctx, arg)
	if err != nil {
		if err == db.ErrUniqueViolation {
			ctx.JSON(http.StatusForbidden, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

// listLoans lists all loans.
//
//	@Summary		List all loans
//	@Description	List all loans.
//	@Tags			loans
//	@Accept			json
//	@Produce		json
//	@Param			request	body		listRequest		true	"List Request"
//	@Success		200		{object}	[]db.Loan		"List of loans"
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/loans/all [post]
//	@security		ApiKeyAuth
func (server *Server) listLoans(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListLoansParams{
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}

	loans, err := server.store.ListLoans(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

// getLoan gets a loan by ID.
//
//	@Summary		Get a loan
//	@Description	Get a loan by ID.
//	@Tags			loans
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Loan ID"
//	@Success		200	{object}	db.Loan			"Loan found"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/loans/{id} [get]
//	@security		ApiKeyAuth
func (server *Server) getLoan(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	loan, err := server.store.GetLoan(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

// searchLoans searches loans by borrower.
//
//	@Summary		Search loans
//	@Description	Search loans by borrower.
//	@Tags			loans
//	@Accept			json
//	@Produce		json
//	@Param			request	body		searchRequest	true	"Search Request"
//	@Success		200		{object}	[]db.Loan		"List of loans"
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/loans/search [post]
//	@security		ApiKeyAuth
func (server *Server) searchLoans(ctx *gin.Context) {
	var req searchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.SearchLoansParams{
		Borrower: req.Query,
		Offset:   (req.PageID - 1) * req.PageSize,
		Limit:    req.PageSize,
	}

	loans, err := server.store.SearchLoans(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

// deleteLoan deletes a loan by ID.
//
//	@Summary		Delete a loan
//	@Description	Delete a loan by ID.
//	@Tags			loans
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Loan ID"
//	@Success		200	{object}	db.Loan			"Lone deleted"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/loans/{id} [delete]
//	@security		ApiKeyAuth
func (server *Server) deleteLoan(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	loan, err := server.store.DeleteLoan(ctx, uuid.MustParse(req.ID))
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, loan)
}
