package http

import (
	"github.com/behrouz-rfa/kentech/internal/core/common"
	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/ports"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/gin-gonic/gin"
	"time"
)

// FilmHandler represents the HTTP handler for film requests
type FilmHandler struct {
	svc ports.FilmService
}

// NewFilmHandler creates a new FilmHandler instance
func NewFilmHandler(svc ports.FilmService) *FilmHandler {
	return &FilmHandler{
		svc,
	}
}

// ListFilms godoc
//
//	@Summary		List films
//	@Description	List films with pagination
//	@Tags			Films
//	@Accept			json
//	@Produce		json
//	@Param			page	query		uint64			true	"Page"
//	@Param			limit	query		uint64			true	"Limit"
//	@Param			genre	query		string			false	"Genre"
//	@Param			title	query		string			false	"Title"
//	@Param			from	query		string		    false	"From example: 2021-02-18 "
//	@Param			to  	query		string		    false	"To   example: 2021-02-18 "
//	@Success		200		{object}	meta			"Films displayed"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/films [get]
//	@Security		BearerAuth
func (h *FilmHandler) ListFilms(ctx *gin.Context) {
	var req listFilmsRequest
	var filmsList []filmResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validationError(ctx, err)
		return
	}
	f := &filters.FilmFilter{}
	if req.Title != nil {
		f.Title = &filters.StringFilter{Contains: req.Title}
	}
	if req.From != nil {
		f.ReleaseDate = &filters.TimeRange{From: (*time.Time)(req.From), To: (*time.Time)(req.To)}
	}
	if req.Genre != nil {
		f.Genre = &filters.StringFilter{Contains: req.Genre}
	}

	films, err := h.svc.GetFilms(ctx, f, nil, &pagination.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		handleError(ctx, err)
		return
	}

	for _, film := range films {
		filmsList = append(filmsList, newFilmResponse(film))
	}

	total := uint64(len(filmsList))
	meta := newMeta(total, req.Limit, req.Page)
	rsp := toMap(meta, filmsList, "films")

	handleSuccess(ctx, rsp)
}

// GetFilm godoc
//
//	@Summary		Get a film
//	@Description	Get a film by id
//	@Tags			Films
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Film ID"
//	@Success		200	{object}	filmResponse	"Film displayed"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/films/{id} [get]
//	@Security		BearerAuth
func (h FilmHandler) GetFilm(ctx *gin.Context) {
	filmID := ctx.Param("id")

	film, err := h.svc.GetFilm(ctx, filters.FilmBy{ID: &filmID})
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newFilmResponse(film)

	handleSuccess(ctx, rsp)
}

// UpdateFilm godoc
//
//	@Summary		Update a film
//	@Description	Update a film's name, email, password, or role by id
//	@Tags			Films
//	@Accept			json
//	@Produce		json
//	@Param			id					path		string				true	"Film ID"
//	@Param			updateFilmRequest	body		updateFilmRequest	true	"Update film request"
//	@Success		200					{object}	filmResponse		"Film updated"
//	@Failure		400					{object}	errorResponse		"Validation error"
//	@Failure		401					{object}	errorResponse		"Unauthorized error"
//	@Failure		403					{object}	errorResponse		"Forbidden error"
//	@Failure		404					{object}	errorResponse		"Data not found error"
//	@Failure		500					{object}	errorResponse		"Internal server error"
//	@Router			/films/{id} [put]
//	@Security		BearerAuth
func (h FilmHandler) UpdateFilm(ctx *gin.Context) {
	var req updateFilmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	payload := GetAuthPayload(ctx, common.AuthorizationPayloadKey)

	filmID := ctx.Param("id")

	filmUpdateInput := &model.FilmUpdateInput{
		Title:    req.Title,
		Director: req.Director,
		Cast:     req.Cast,
		Genre:    req.Genre,
		Synopsis: req.Synopsis,
	}

	if req.ReleaseDate != nil {
		filmUpdateInput.ReleaseDate = (time.Time)(*req.ReleaseDate)
	}

	film, err := h.svc.UpdateFilm(ctx, filmID, filmUpdateInput, payload.UserID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newFilmResponse(film)

	handleSuccess(ctx, rsp)
}

// DeleteFilm godoc
//
//	@Summary		Delete a film
//	@Description	Delete a film by id
//	@Tags			Films
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string			true	"Film ID"
//	@Success		200	{object}	response		"Film deleted"
//	@Failure		400	{object}	errorResponse	"Validation error"
//	@Failure		401	{object}	errorResponse	"Unauthorized error"
//	@Failure		403	{object}	errorResponse	"Forbidden error"
//	@Failure		404	{object}	errorResponse	"Data not found error"
//	@Failure		500	{object}	errorResponse	"Internal server error"
//	@Router			/films/{id} [delete]
//	@Security		BearerAuth
func (h FilmHandler) DeleteFilm(ctx *gin.Context) {
	filmID := ctx.Param("id")

	err := h.svc.DeleteFilm(ctx, filmID)
	if err != nil {
		handleError(ctx, err)
		return
	}

	handleSuccess(ctx, nil)
}

// CreateFilm godoc
//
//	@Summary		Create a film
//	@Description	create a film
//	@Tags			Films
//	@Accept			json
//	@Produce		json
//	@Param			createFilmRequest	body		createFilmRequest	true	"Create film request"
//	@Success		200					{object}	filmResponse		"Film updated"
//	@Failure		400					{object}	errorResponse		"Validation error"
//	@Failure		401					{object}	errorResponse		"Unauthorized error"
//	@Failure		403					{object}	errorResponse		"Forbidden error"
//	@Failure		404					{object}	errorResponse		"Data not found error"
//	@Failure		500					{object}	errorResponse		"Internal server error"
//	@Router			/films [POST]
//	@Security		BearerAuth
func (h *FilmHandler) CreateFilm(ctx *gin.Context) {
	payload := GetAuthPayload(ctx, common.AuthorizationPayloadKey)

	var req createFilmRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError(ctx, err)
		return
	}

	filmUpdateInput := &model.FilmInput{
		Title:       req.Title,
		Director:    req.Director,
		ReleaseDate: time.Time(req.ReleaseDate),
		Cast:        req.Cast,
		Genre:       model.Genre(req.Genre),
		Synopsis:    req.Synopsis,
		CreatorID:   payload.UserID,
	}

	film, err := h.svc.CreateFilm(ctx, filmUpdateInput)
	if err != nil {
		handleError(ctx, err)
		return
	}

	rsp := newFilmResponse(film)

	handleSuccess(ctx, rsp)
}
