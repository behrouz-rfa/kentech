package services

import (
	"context"
	"fmt"
	specification "github.com/behrouz-rfa/kentech/internal/core/specefication"

	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/ports"

	"github.com/behrouz-rfa/kentech/internal/core/util"
	"github.com/behrouz-rfa/kentech/internal/filters"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"github.com/behrouz-rfa/kentech/internal/sort"
	"github.com/behrouz-rfa/kentech/pkg/logger"
)

// UserService is a service that handles user-related operations.
type UserService struct {
	userRepo ports.UserRepository
	auth     ports.Auth
	lg       *logger.Entry
}

// UserServiceOption is a function that configures the UserService.
type UserServiceOption func(*UserService)

// WithUserRepository sets the user repository for the UserService.
func WithUserRepository(repo ports.UserRepository) UserServiceOption {
	return func(s *UserService) {
		s.userRepo = repo
	}
}

// WithAuth sets the authentication service for the UserService.
func WithAuth(auth ports.Auth) UserServiceOption {
	return func(s *UserService) {
		s.auth = auth
	}
}

// WithLogger sets the logger for the UserService.
func WithLogger(lg *logger.Entry) UserServiceOption {
	return func(s *UserService) {
		s.lg = lg
	}
}

// NewUserService creates a new instance of UserService with the provided options.
func NewUserService(opts ...UserServiceOption) *UserService {
	s := &UserService{
		lg: logger.General.Component("UserService"),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// GetUser retrieves a user by the specified filter.
func (s *UserService) GetUser(ctx context.Context, filter filters.UserBy) (*model.User, error) {
	spec := s.userRepo.NewUserSpecification(ctx).
		By(filter).
		Prefetch("films").(specification.UserSpecification)

	return s.userRepo.GetUser(ctx, spec)
}

// GetUsers retrieves a list of users based on the specified filter, sort, and pagination options.
func (s *UserService) GetUsers(ctx context.Context, filter *filters.UserFilter, sort *sort.UserSort, paginate *pagination.Pagination) ([]*model.User, error) {
	spec := s.userRepo.NewUserSpecification(ctx).
		Filter(filter).
		SortBy(sort).
		Prefetch("films").(specification.UserSpecification)

	return s.userRepo.GetUsers(ctx, spec, paginate)
}

// CreateUser creates a new user with the provided input.
func (s *UserService) CreateUser(ctx context.Context, input *model.UserInput) (*model.User, error) {
	existingUser, err := s.GetUser(ctx, filters.UserBy{Username: &input.Username})
	if err != nil && err.Error() != "could not found" {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("%w: user with username %s already exists", model.ErrConflictingData, input.Username)
	}

	hashedPassword, err := util.HashPassword(input.Password)
	if err != nil {
		return nil, model.ErrInternal
	}
	input.Password = hashedPassword

	userID, err := s.userRepo.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	token, err := s.auth.Create(model.TokenPayload{UserID: userID, Username: input.Username})
	if err != nil {
		return nil, model.ErrInternal
	}

	user, err := s.GetUser(ctx, filters.UserBy{ID: &userID})
	if err != nil {
		return nil, err
	}
	user.JwtToken = token

	return user, nil
}

// Login authenticates a user with the provided login input.
func (s *UserService) Login(ctx context.Context, input *model.UserLoginInput) (*model.User, error) {
	user, err := s.GetUser(ctx, filters.UserBy{Username: &input.Username})
	if err != nil {
		return nil, err
	}

	if err := util.ComparePassword(input.Password, user.Password); err != nil {
		return nil, model.ErrInternal
	}

	token, err := s.auth.Create(model.TokenPayload{Username: user.Username, UserID: user.ID})
	if err != nil {
		return nil, model.ErrInternal
	}

	user.JwtToken = token
	return user, nil
}

// UpdateUser updates a user with the provided input.
func (s *UserService) UpdateUser(ctx context.Context, id string, input *model.UserUpdateInput) (*model.User, error) {
	if err := s.userRepo.UpdateUser(ctx, id, input); err != nil {
		return nil, err
	}

	return s.GetUser(ctx, filters.UserBy{ID: &id})
}

// DeleteUser deletes a user with the specified ID.
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}
