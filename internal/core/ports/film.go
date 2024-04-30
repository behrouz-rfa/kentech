package ports

import (
	"context"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/behrouz-rfa/kentech/internal/sort"
)

type FilmRepository interface {
	GetFilm(ctx context.Context, spec specification.FilmSpecification) (*model.Film, error)
	GetFilms(ctx context.Context, spec specification.FilmSpecification,
		paginate *pagination.Pagination) ([]*model.Film, error)
	CreateFilm(ctx context.Context, filmInput *model.FilmInput) (string, error)
	UpdateFilm(ctx context.Context, id string, filmInput *model.FilmUpdateInput) error
	DeleteFilm(ctx context.Context, id string) error
	NewFilmSpecification(ctx context.Context) specification.FilmSpecification
}

type FilmService interface {
	GetFilm(ctx context.Context, filter filters.FilmBy) (*model.Film, error)
	GetFilms(ctx context.Context, filter *filters.FilmFilter, sort *sort.FilmSort,
		paginate *pagination.Pagination) ([]*model.Film, error)
	CreateFilm(ctx context.Context, filmInput *model.FilmInput) (*model.Film, error)
	UpdateFilm(ctx context.Context, id string, filmInput *model.FilmUpdateInput) (*model.Film, error)
	DeleteFilm(ctx context.Context, id string) error
}
