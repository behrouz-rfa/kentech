package http

import (
	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/ports"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/gin-gonic/gin"
)

// UserHandler represents the HTTP handler for film requests
type UserHandler struct {
	svc ports.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc ports.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}

// registerRequest represents the request body for creating a user
type registerRequest struct {
	Firstname string `json:"firstname" binding:"omitempty,required" example:"Jon"`
	Lastname  string `json:"lastname" binding:"omitempty,required" example:"doa"`
	Username  string `json:"username" binding:"omitempty,required" example:"jondoa"`
	Password  string `json:"password" binding:"required,min=8" example:"12345678"`
}

// Register godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account with default role "cashier"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		registerRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users/register [post]
func (h UserHandler) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := model.UserInput{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Password:  req.Password,
	}

	usr, err := h.svc.CreateUser(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUserResponse(usr))
}

// registerRequest represents the request body for creating a user
type loginRequest struct {
	Username string `json:"username" binding:"omitempty,required" example:"jondoa"`
	Password string `json:"password" binding:"required,min=8" example:"12345678"`
}

// Login godoc
//
//	@Summary		Register a new user
//	@Description	create a new user account with default role "cashier"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			loginRequest	body		loginRequest	true	"Register request"
//	@Success		200				{object}	userResponse	"User created"
//	@Failure		400				{object}	errorResponse	"Validation error"
//	@Failure		401				{object}	errorResponse	"Unauthorized error"
//	@Failure		404				{object}	errorResponse	"Data not found error"
//	@Failure		409				{object}	errorResponse	"Data conflict error"
//	@Failure		500				{object}	errorResponse	"Internal server error"
//	@Router			/users/login [post]
func (h UserHandler) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	user := model.UserLoginInput{
		Username: req.Username,
		Password: req.Password,
	}

	usr, err := h.svc.Login(ctx, &user)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, newUserResponse(usr))
}

// listUsersRequest represents the request body for listing users
type listUsersRequest struct {
	Page  uint64 `form:"page" binding:"required,min=0" example:"0"`
	Limit uint64 `form:"limit" binding:"required,min=5" example:"10"`
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	List users with pagination
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			page	query		uint64			true	"Page"
//	@Param			limit	query		uint64			true	"Limit"
//	@Success		200		{object}	meta			"Users displayed"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/users [get]
//	@Security		BearerAuth
func (h *UserHandler) ListUsers(ctx *gin.Context) {
	var req listUsersRequest
	var usersList []userResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}

	users, err := h.svc.GetUsers(ctx, nil, nil, &pagination.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, user := range users {
		usersList = append(usersList, newUserResponse(user))
	}

	total := uint64(len(usersList))
	meta := newMeta(total, req.Limit, req.Page)
	rsp := toMap(meta, usersList, "users")

	handleSuccess(ctx, rsp)
}

// GetUser godoc
//
//	@Summary		Get a user
//	@Description	Get a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	userResponse	"User displayed"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [get]
//	@Security		BearerAuth
func (h UserHandler) GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := h.svc.GetUser(ctx, filters.UserBy{ID: &userID})
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)

	handleSuccess(ctx, rsp)
}

// updateUserRequest represents the request body for updating a user
type updateUserRequest struct {
	Firstname *string `json:"firstname" binding:"omitempty" example:"Jon"`
	Lastname  *string `json:"lastname" binding:"omitempty" example:"doa"`
	Username  *string `json:"username" binding:"omitempty" example:"jondoa"`
}

// UpdateUser godoc
//
//	@Summary		Update a user
//	@Description	Update a user's name, email, password, or role by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string				true	"User ID"
//	@Param			updateUserRequest	body		updateUserRequest	true	"Update user request"
//	@Success		200					{object}	userResponse		"User updated"
//	@Failure		400					{object}	errorResponse		"Validation error"
//	@Failure		401					{object}	errorResponse		"Unauthorized error"
//	@Failure		403					{object}	errorResponse		"Forbidden error"
//	@Failure		404					{object}	errorResponse		"Data not found error"
//	@Failure		500					{object}	errorResponse		"Internal server error"
//	@Router			/users/{id} [put]
//	@Security		BearerAuth
func (h UserHandler) UpdateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	userID := ctx.Param("id")

	userUpdateInput := &model.UserUpdateInput{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	user, err := h.svc.UpdateUser(ctx, userID, userUpdateInput)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newUserResponse(user)

	handleSuccess(ctx, rsp)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Delete a user by id
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"User ID"
//	@Success		200	{object}	response		"User deleted"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/users/{id} [delete]
//	@Security		BearerAuth
func (h UserHandler) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	err := h.svc.DeleteUser(ctx, userID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}
