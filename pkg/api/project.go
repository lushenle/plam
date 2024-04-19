package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lushenle/plam/pkg/db"
)

// createProjectRequest is a struct that represents the request to create a project.
//
//	@swagger:model
type createProjectRequest struct {
	// Name of the project.
	// Required: true
	// example: project1
	// in: body
	Name string `json:"name" binding:"required"`

	// Description of the project.
	// Required: true
	// example: project1 description
	// in: body
	Description string `json:"description" binding:"required"`

	// Amount of the project.
	// Required: true
	// example: 1000
	// in: body
	Amount float32 `json:"amount" binding:"required,gt=0"`
}

// createProject creates a new project.
//
//	@Summary		Create a project
//	@Description	Create a new project.
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createProjectRequest	true	"Create Project Request"
//	@Success		200		{object}	db.Project				"Project created"
//	@Failure		400		{object}	errorResponse			"Bad Request"
//	@Failure		401		{object}	errorResponse			"Unauthorized"
//	@Failure		403		{object}	errorResponse			"Forbidden"
//	@Failure		500		{object}	errorResponse			"Internal Server Error"
//	@Router			/projects [post]
//	@security		ApiKeyAuth
func (server *Server) createProject(ctx *gin.Context) {
	var req createProjectRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateProjectParams{
		Name:        req.Name,
		Description: req.Description,
		Amount:      req.Amount,
	}

	project, err := server.store.CreateProject(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrUniqueViolation) {
			ctx.JSON(http.StatusForbidden, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, project)
}

// listRequest is a struct that represents the request to list projects.
//
//	@swagger:model
type listRequest struct {
	// PageID is the page number.
	// Required: true
	// example: 1
	// in: body
	// minimum: 1
	PageID int32 `json:"page_id" binding:"required,min=1"`

	// PageSize is the number of projects per page.
	// Required: true
	// example: 5
	// in: body
	// minimum: 5
	// maximum: 10
	PageSize int32 `json:"page_size" binding:"required,min=5,max=100"`
}

// listProjects lists all projects.
//
//	@Summary		List projects
//	@Description	List all projects.
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Param			request	body		listRequest		true	"List Request"
//	@Success		200		{array}		[]db.Project	"List of projects"
//	@Failure		400		{object}	errorResponse	"Bad Request"
//	@Failure		401		{object}	errorResponse	"Unauthorized"
//	@Failure		403		{object}	errorResponse	"Forbidden"
//	@Failure		500		{object}	errorResponse	"Internal Server Error"
//	@Router			/projects/all [post]
//	@security		ApiKeyAuth
func (server *Server) listProjects(ctx *gin.Context) {
	var req listRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListProjectsParams{
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}
	projects, err := server.store.ListProjects(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, projects)
}

// getRequest is a struct that represents the request to get a project.
//
//	@swagger:model
type getRequest struct {
	// ID is the project ID.
	// Required: true
	// swagger:strfmt uuid
	// example: 123e4567-e89b-12d3-a456-426614174000
	// in: path
	ID string `uri:"id" binding:"required,uuid"`
}

// getProject gets a project by ID.
//
//	@Summary		Get a project
//	@Description	Get a project.
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Project ID"
//	@Success		200	{object}	db.Project		"Project found"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/projects/{id} [get]
//	@security		ApiKeyAuth
func (server *Server) getProject(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	project, err := server.store.GetProject(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, project)
}

// searchRequest is a struct that represents the request to search projects.
//
//	@swagger:model
type searchRequest struct {
	// Query is the search query.
	// Required: true
	// example: project1
	// in: body
	Query string `json:"query" binding:"required"`

	// PageID is the page number.
	// Required: true
	// example: 1
	// in: body
	// minimum: 1
	PageID int32 `json:"page_id" binding:"required,min=1"`

	// PageSize is the number of projects per page.
	// Required: true
	// example: 5
	// in: body
	// minimum: 5
	// maximum: 10
	PageSize int32 `json:"page_size" binding:"required,min=5,max=10"`
}

// searchProjects searches projects.
//
//	@Summary		Search projects
//	@Description	Search projects.
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Param			request	body		listRequest		true	"Search Request"
//	@Success		200		{array}		[]db.Project	"List of projects"
//	@Failure		400		{object}	string			"Bad Request"
//	@Failure		401		{object}	string			"Unauthorized"
//	@Failure		403		{object}	string			"Forbidden"
//	@Failure		500		{object}	string			"Internal Server Error"
//	@Router			/projects/search [post]
//	@security		ApiKeyAuth
func (server *Server) searchProjects(ctx *gin.Context) {
	var req searchRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.SearchProjectsParams{
		Column1: pgtype.Text{
			String: req.Query,
			Valid:  true,
		},
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}

	projects, err := server.store.SearchProjects(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, projects)
}

// deleteProject deletes a project by ID.
//
//	@Summary		Delete a project
//	@Description	Delete a project.
//	@Tags			projects
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Project ID"
//	@Success		200	{object}	db.Project		"Project deleted"
//	@Failure		400	{object}	errorResponse	"Bad Request"
//	@Failure		401	{object}	errorResponse	"Unauthorized"
//	@Failure		403	{object}	errorResponse	"Forbidden"
//	@Failure		404	{object}	errorResponse	"Not Found"
//	@Failure		500	{object}	errorResponse	"Internal Server Error"
//	@Router			/projects/{id} [delete]
//	@security		ApiKeyAuth
func (server *Server) deleteProject(ctx *gin.Context) {
	var req getRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	project, err := server.store.DeleteProject(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, project)
}
