package mongo

import (
	"context"
	"errors"

	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/mspecification"
	"github.com/behrouz-rfa/kentech/internal/adapter/mongo/utils"
	cerr "github.com/behrouz-rfa/kentech/internal/core/errors"
	"github.com/behrouz-rfa/kentech/internal/core/model"
	"github.com/behrouz-rfa/kentech/internal/core/specefication"
	"github.com/behrouz-rfa/kentech/internal/pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	DocumentBase `bson:",inline"`
	Username     string  `bson:"username" `
	Password     string  `bson:"password" `
	Firstname    string  `bson:"firstname" `
	Lastname     string  `bson:"lastname" `
	Films        []*Film `bson:"films" `
}

func (u User) ToModel() *model.User {
	s := &model.User{
		ID:        u.ID,
		Username:  u.Username,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Password:  u.Password,
	}

	for _, f := range u.Films {
		s.Films = append(s.Films, *f.ToModel())
	}

	return s
}

func (m Repository) userCollection() *mongo.Collection {
	return m.db.Collection("users")
}

func (m Repository) GetUser(ctx context.Context, spec specification.UserSpecification) (*model.User, error) {
	return FindOneBy(ctx, &FindByParams[User, model.User]{
		Collection: m.userCollection(),
		Spec:       spec,
		ToModel:    User.ToModel,
	})
}

func (m Repository) GetUsers(ctx context.Context, spec specification.UserSpecification,
	paginate *pagination.Pagination,
) ([]*model.User, error) {
	if spec == nil {
		spec = m.NewUserSpecification(ctx)
	}

	spec.Paginate(paginate)

	return FindBy(ctx, &FindByParams[User, model.User]{
		Collection: m.userCollection(),
		Spec:       spec,
		ToModel:    User.ToModel,
	})
}

func (m Repository) CreateUser(ctx context.Context, obj *model.UserInput) (string, error) {
	col := m.userCollection()
	data := utils.ToMap(obj)
	i, err := col.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return i.InsertedID.(string), nil
}

func (m Repository) UpdateUser(ctx context.Context, id string, obj *model.UserUpdateInput) error {
	filter := bson.M{"_id": id}
	data := utils.ToMap(obj, utils.MethodUpdate)

	col := m.userCollection()

	_, err := col.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		m.lg.WithError(err).Error("failed to update user")
		return cerr.Wrap(err, cerr.ErrInternal)
	}

	return nil
}

func (m Repository) DeleteUser(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}

	col := m.userCollection()

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		m.lg.WithError(err).Error("failed to delete user")
		return cerr.Wrap(err, cerr.ErrInternal)
	}

	if result.DeletedCount == 0 {
		m.lg.WithError(err).Error("nothing found for delete")
		return cerr.Wrap(errors.New("nothing found for delete"), cerr.ErrNotFound)
	}

	return nil
}

func (m Repository) NewUserSpecification(context context.Context) specification.UserSpecification {
	spec := mspecification.NewUserSpecification()
	spec.WithContext(context)

	return spec
}
