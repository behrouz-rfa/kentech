package ports

import (
	"context"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	specification "github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/behrouz-rfa/kentech/internal/sort"
)

type UserRepository interface {
	GetUser(ctx context.Context, spec specification.UserSpecification) (*model.User, error)
	GetUsers(ctx context.Context, spec specification.UserSpecification,
		paginate *pagination.Pagination) ([]*model.User, error)
	CreateUser(ctx context.Context, filmInput *model.UserInput) (string, error)
	UpdateUser(ctx context.Context, id string, filmInput *model.UserUpdateInput) error
	DeleteUser(ctx context.Context, id string) error
	NewUserSpecification(ctx context.Context) specification.UserSpecification
}

type UserService interface {
	GetUser(ctx context.Context, filter filters.UserBy) (*model.User, error)
	GetUsers(ctx context.Context, filter *filters.UserFilter, sort *sort.UserSort,
		paginate *pagination.Pagination) ([]*model.User, error)
	CreateUser(ctx context.Context, filmInput *model.UserInput) (*model.User, error)
	Login(ctx context.Context, filmInput *model.UserLoginInput) (*model.User, error)
	UpdateUser(ctx context.Context, id string, filmInput *model.UserUpdateInput) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
}
