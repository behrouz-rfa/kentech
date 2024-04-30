package services

import (
	"context"
	"fmt"
	specification "github.com/behrouz-rfa/kentech/internal/core/specefication"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/ports"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/behrouz-rfa/kentech/internal/sort"
	"github.com/behrouz-rfa/kentech/pkg/logger"
)

// FilmService is a service that handles film-related operations.
type FilmService struct {
	filmRepo ports.FilmRepository
	lg       *logger.Entry
}

// FilmServiceOption is a function that configures the FilmService.
type FilmServiceOption func(*FilmService)

// WithFilmRepository sets the film repository for the FilmService.
func WithFilmRepository(repo ports.FilmRepository) FilmServiceOption {
	return func(s *FilmService) {
		s.filmRepo = repo
	}
}

// NewFilmService creates a new instance of FilmService with the provided options.
func NewFilmService(opts ...FilmServiceOption) *FilmService {
	s := &FilmService{
		lg: logger.General.Component("FilmService"),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// GetFilm retrieves a film by the specified filter.
func (s *FilmService) GetFilm(ctx context.Context, filter filters.FilmBy) (*model.Film, error) {
	spec := s.filmRepo.NewFilmSpecification(ctx).By(filter)
	return s.filmRepo.GetFilm(ctx, spec)
}

// GetFilms retrieves a list of films based on the specified filter, sort, and pagination options.
func (s *FilmService) GetFilms(ctx context.Context, filter *filters.FilmFilter, sort *sort.FilmSort, paginate *pagination.Pagination) ([]*model.Film, error) {
	spec := s.filmRepo.NewFilmSpecification(ctx).
		Filter(filter).
		SortBy(sort).(specification.FilmSpecification)
	return s.filmRepo.GetFilms(ctx, spec, paginate)
}

// CreateFilm creates a new film with the provided input.
func (s *FilmService) CreateFilm(ctx context.Context, input *model.FilmInput) (*model.Film, error) {
	filmID, err := s.filmRepo.CreateFilm(ctx, input)
	if err != nil {
		return nil, err
	}

	return s.GetFilm(ctx, filters.FilmBy{ID: &filmID})
}

// UpdateFilm updates a film with the provided input.
func (s *FilmService) UpdateFilm(ctx context.Context, id string, input *model.FilmUpdateInput) (*model.Film, error) {
	if err := s.filmRepo.UpdateFilm(ctx, id, input); err != nil {
		return nil, err
	}

	return s.GetFilm(ctx, filters.FilmBy{ID: &id})
}

// DeleteFilm deletes a film with the specified ID.
func (s *FilmService) DeleteFilm(ctx context.Context, id string) error {
	existingFilm, err := s.GetFilm(ctx, filters.FilmBy{ID: &id})
	if err != nil {
		return err
	}

	if existingFilm == nil {
		return fmt.Errorf("film with ID %s not found", id)
	}

	return s.filmRepo.DeleteFilm(ctx, id)
}
