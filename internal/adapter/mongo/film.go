package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/mspecification"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/utils"
	cerr "github.com/behrouz-rfa/kentech/internal/core/errors"
	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Film struct {
	DocumentBase `bson:",inline"`
	Title        string    `bson:"title"`
	Director     string    `bson:"director"`
	ReleaseDate  time.Time `bson:"releaseDate"`
	Cast         []string  `bson:"cast"`
	Genre        string    `bson:"genre"`
	Synopsis     string    `bson:"synopsis"`
	CreatorID    string    `bson:"creatorId"`
}

func (u Film) ToModel() *model.Film {
	s := &model.Film{
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		Title:       u.Title,
		Director:    u.Director,
		ReleaseDate: u.ReleaseDate,
		Cast:        u.Cast,
		Genre:       u.Genre,
		Synopsis:    u.Synopsis,
		CreatorID:   u.CreatorID,
	}

	return s
}

func (m Repository) filmCollection() *mongo.Collection {
	return m.db.Collection("films")
}

func (m Repository) GetFilm(ctx context.Context, spec specification.FilmSpecification) (*model.Film, error) {
	return FindOneBy(ctx, &FindByParams[Film, model.Film]{
		Collection: m.filmCollection(),
		Spec:       spec,
		ToModel:    Film.ToModel,
	})
}

func (m Repository) GetFilms(ctx context.Context, spec specification.FilmSpecification,
	paginate *pagination.Pagination,
) ([]*model.Film, error) {
	if spec == nil {
		spec = m.NewFilmSpecification(ctx)
	}

	spec.Paginate(paginate)

	return FindBy(ctx, &FindByParams[Film, model.Film]{
		Collection: m.filmCollection(),
		Spec:       spec,
		ToModel:    Film.ToModel,
	})
}

func (m Repository) CreateFilm(ctx context.Context, obj *model.FilmInput) (string, error) {
	col := m.filmCollection()
	data := utils.ToMap(obj)
	i, err := col.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return i.InsertedID.(string), nil
}

func (m Repository) UpdateFilm(ctx context.Context, id string, obj *model.FilmUpdateInput) error {
	filter := bson.M{"_id": id}
	data := utils.ToMap(obj, utils.MethodUpdate)

	col := m.filmCollection()

	_, err := col.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		m.lg.WithError(err).Error("failed to update film")
		return cerr.Wrap(err, cerr.ErrInternal)
	}

	return nil
}

func (m Repository) DeleteFilm(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}

	col := m.filmCollection()

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		m.lg.WithError(err).Error("failed to delete film")
		return cerr.Wrap(err, cerr.ErrInternal)
	}

	if result.DeletedCount == 0 {
		m.lg.WithError(err).Error("nothing found for delete")
		return cerr.Wrap(errors.New("nothing found for delete"), cerr.ErrNotFound)
	}

	return nil
}

func (m Repository) NewFilmSpecification(context context.Context) specification.FilmSpecification {
	spec := mspecification.NewFilmSpecification()
	spec.WithContext(context)

	return spec
}
