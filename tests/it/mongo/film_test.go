//go:build integration

package mongo

import (
	"context"
	"time"

	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/mspecification"
	"github.com/behrouz-rfa/kentech/internal/core/model"
	specification "github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/filters"
)

func (s *MongoTestSuite) TestFilmCreate() {
	testuser := &model.UserInput{
		Firstname: "master",
		Lastname:  "rfa",
		Username:  "master",
	}

	id, err := s.db.CreateUser(context.Background(), testuser)
	s.NoError(err)
	s.Require().NotEmpty(id, "id should not be empty")
	s.Require().NotEqual(id, "", "id should not be equal to ``")

	testFilm := &model.FilmInput{
		Title:       "Master",
		Director:    "KenTech",
		ReleaseDate: time.Now(),
		Cast:        []string{"Cast1"},
		Genre:       "Horro movies",
		Synopsis:    "",
		CreatorID:   id,
	}

	filmID, err := s.db.CreateFilm(context.Background(), testFilm)

	s.NoError(err)
	s.Require().NotEmpty(filmID, "id should not be empty")
	s.Require().NotEqual(filmID, "", "id should not be equal to ``")

	spec := NewFilmSpecification(context.Background()).By(filters.FilmBy{ID: &filmID})

	sch, err := s.db.GetFilm(context.Background(), spec)
	s.Require().NoError(err)
	s.Require().NotNil(sch, "user should not be nil")
	s.Require().Equal(testFilm.Title, sch.Title, "Title should be equal")

	s.Require().Equal(sch.ID, filmID, "id should be equal")

	// Delete successfully
	err = s.db.DeleteFilm(context.Background(), filmID)
	s.Require().NoError(err)

	err = s.db.DeleteUser(context.Background(), id)
	s.Require().NoError(err)
}

func (s *MongoTestSuite) TestFilmUpdate() {
	testuser := &model.UserInput{
		Firstname: "master",
		Lastname:  "rfa",
		Username:  "master",
	}

	id, err := s.db.CreateUser(context.Background(), testuser)
	s.NoError(err)
	s.Require().NotEmpty(id, "id should not be empty")
	s.Require().NotEqual(id, "", "id should not be equal to ``")

	Title := "newMoview"

	testCreateFilm := &model.FilmInput{
		Title:       "Master",
		Director:    "KenTech",
		ReleaseDate: time.Now(),
		Cast:        []string{"Cast1"},
		Genre:       "Horro movies",
		Synopsis:    "",
		CreatorID:   id,
	}

	filmId, err := s.db.CreateFilm(context.Background(), testCreateFilm)

	s.NoError(err)
	s.Require().NotEmpty(id, "id should not be empty")
	s.Require().NotEqual(id, "", "id should not be equal to ``")
	updateFilm := &model.FilmUpdateInput{
		Title: Title,
	}
	// update name successfully
	err = s.db.UpdateFilm(context.Background(), filmId, updateFilm)
	s.Require().NoError(err)

	s.Require().NoError(err)

	s.Require().NoError(err)

	spec := NewFilmSpecification(context.Background()).By(filters.FilmBy{ID: &filmId})
	sch, err := s.db.GetFilm(context.Background(), spec)
	s.Require().NoError(err)
	s.Require().NotNil(sch, "film should not be nil")
	s.Require().Equal(sch.Title, Title)
	s.Require().Equal(sch.Genre, testCreateFilm.Genre)

	// Delete successfully
	err = s.db.DeleteFilm(context.Background(), filmId)
	s.Require().NoError(err)

	err = s.db.DeleteUser(context.Background(), id)
	s.Require().NoError(err)
}

func NewFilmSpecification(context context.Context) specification.FilmSpecification {
	spec := mspecification.NewFilmSpecification()
	spec.WithContext(context)

	return spec
}
