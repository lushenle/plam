package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lushenle/plam/pkg/db"
	"github.com/lushenle/plam/pkg/util"
)

// createUserRequest is a struct that represents the request to create a user.
//
//	@swagger:model
type createUserRequest struct {
	// Username of the user.
	// Required: true
	// example: john_doe
	// in: body
	// minLength: 1
	// maxLength: 255
	Username string `json:"username" binding:"required,alphanum"`

	// Password of the user.
	// Required: true
	// example: password123
	// in: body
	// minLength: 6
	// maxLength: 32
	Password string `json:"password" binding:"required,min=6,max=32"`

	// Full name of the user.
	// Required: true
	// example: John Doe
	// in: body
	// minLength: 1
	// maxLength: 255
	FullName string `json:"full_name" binding:"required"`

	// Email of the user.
	// Required: true
	// example: john_doe@example.com
	// in: body
	// format: email
	Email string `json:"email" binding:"required,email"`
}

// userResponse represents the response structure for user data.
//
//	@swagger:model
type userResponse struct {
	// Username of the user.
	// example: john_doe
	Username string `json:"username"`

	// Full name of the user.
	// example: John Doe
	FullName string `json:"full_name"`

	// Email of the user.
	// example: john_doe@example.com
	Email string `json:"email"`

	// PasswordChangedAt represents the timestamp when the password was last changed.
	// swagger:strfmt date-time
	PasswordChangedAt time.Time `json:"password_changed_at"`

	// CreatedAt represents the timestamp when the user was created.
	// swagger:strfmt date-time
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt represents the timestamp when the user was last updated.
	// swagger:strfmt date-time
	UpdatedAt time.Time `json:"updated_at"`
}

// signupUser creates a new user.
//
//	@Summary		Creates a new user.
//	@Description	Creates a new user.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		createUserRequest	true	"User creation request"
//	@Success		200		{object}	userResponse		"User creation response"
//	@Failure		400		{object}	errorResponse		"Bad request"
//	@Failure		403		{object}	errorResponse		"Forbidden"
//	@Failure		500		{object}	errorResponse		"Internal server error"
//	@Router			/users/signup [post]
func (server *Server) signupUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrUniqueViolation) {
			ctx.JSON(http.StatusForbidden, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

// loginUserRequest represents the request structure for user login.
//
//	@swagger:model
type loginUserRequest struct {
	// Username of the user.
	// Required: true
	// example: john_doe
	// in: body
	Username string `json:"username" binding:"required,alphanum"`

	// Password of the user.
	// Required: true
	// example: password123
	// in: body
	// minLength: 6
	// maxLength: 32
	Password string `json:"password" binding:"required,min=6,max=32"`
}

// loginUserResponse represents the response structure for user login.
//
//	@swagger:model
type loginUserResponse struct {
	// Access token ID.
	// swagger:strfmt uuid
	// example: 123e4567-e89b-12d3-a456-426614174000
	AccessTokenID uuid.UUID `json:"access_token_id"`

	// Access token.
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
	AccessToken string `json:"access_token"`

	// Access token expiration time.
	// swagger:strfmt date-time
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`

	// User information.
	User userResponse `json:"user"`
}

// loginUser logs in a user.
//
//	@Summary		Logs in a user.
//	@Description	Logs in a user.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		loginUserRequest	true	"User login request"
//	@Success		200		{object}	loginUserResponse	"User login response"
//	@Failure		400		{object}	errorResponse		"Bad request"
//	@Failure		401		{object}	errorResponse		"Unauthorized"
//	@Failure		404		{object}	errorResponse		"Not found"
//	@Failure		500		{object}	errorResponse		"Internal server error"
//	@Router			/users/login [post]
func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(user.Username, user.Role, server.config.Server.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessTokenID:        accessPayload.ID,
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		User:                 newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
