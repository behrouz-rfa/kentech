//go:build integration

package mongo

import (
	"context"

	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/mspecification"
	"github.com/behrouz-rfa/kentech/internal/core/model"
	specification "github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/filters"
)

func (s *MongoTestSuite) TestUserCreate() {
	testuser := &model.UserInput{
		Firstname: "master",
		Lastname:  "rfa",
		Username:  "master",
	}

	id, err := s.db.CreateUser(context.Background(), testuser)

	s.NoError(err)
	s.Require().NotEmpty(id, "id should not be empty")
	s.Require().NotEqual(id, "", "id should not be equal to ``")

	spec := NewUserSpecification(context.Background()).By(filters.UserBy{ID: &id})

	sch, err := s.db.GetUser(context.Background(), spec)
	s.Require().NoError(err)
	s.Require().NotNil(sch, "user should not be nil")
	s.Require().Equal(testuser.Username, sch.Username, "Username should be equal")
	s.Require().Equal(testuser.Firstname, sch.Firstname, "Firstname should be equal")
	s.Require().Equal(testuser.Lastname, sch.Lastname, "Lastname should be equal")
	s.Require().Equal(id, sch.ID, "id should be equal")

	// Delete successfully
	err = s.db.DeleteUser(context.Background(), id)
	s.Require().NoError(err)
}

func (s *MongoTestSuite) TestUserUpdate() {
	firstName := "Behrouz"
	lastName := "R.FA"
	username := "master2"

	testCreateUser := &model.UserInput{
		Firstname: "master",
		Lastname:  "rfa",
		Username:  "master",
	}

	id, err := s.db.CreateUser(context.Background(), testCreateUser)

	s.NoError(err)
	s.Require().NotEmpty(id, "id should not be empty")
	s.Require().NotEqual(id, "", "id should not be equal to ``")
	updateUser := &model.UserUpdateInput{
		Username:  &username,
		Firstname: &firstName,
		Lastname:  &lastName,
	}
	// update name successfully
	err = s.db.UpdateUser(context.Background(), id, updateUser)
	s.Require().NoError(err)

	s.Require().NoError(err)

	s.Require().NoError(err)

	spec := NewUserSpecification(context.Background()).By(filters.UserBy{ID: &id})
	sch, err := s.db.GetUser(context.Background(), spec)
	s.Require().NoError(err)
	s.Require().NotNil(sch, "user should not be nil")
	s.Require().Equal(sch.Firstname, firstName)
	s.Require().Equal(sch.Lastname, lastName)

	// Delete successfully
	err = s.db.DeleteUser(context.Background(), id)
	s.Require().NoError(err)
}

func NewUserSpecification(context context.Context) specification.UserSpecification {
	spec := mspecification.NewUserSpecification()
	spec.WithContext(context)

	return spec
}
