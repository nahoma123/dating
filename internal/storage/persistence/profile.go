package persistence

import (
	"context"
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/storage"
	"dating/platform/logger"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type profile struct {
	db *mongo.Database
}

func InitProfileDB(db *mongo.Database) storage.ProfileStorage {
	return &profile{
		db: db,
	}
}

func (p *profile) Create(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	// create database

	// createIndex( { "hostname": 1 }, { unique: true } )
	id, _ := uuid.NewV4()
	profile.ProfileID = id.String()

	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	profile.Status = constant.Active
	_, err := p.db.Collection(string(storage.Profile)).InsertOne(ctx, profile)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}
	}

	return profile, err
}

func (p *profile) Update(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	// create database

	// createIndex( { "hostname": 1 }, { unique: true } )

	profile.UpdatedAt = time.Now()
	updateProfile, err := bson.Marshal(profile)
	if err != nil {
		return nil, err
	}

	_, err = p.db.Collection(string(storage.Profile)).UpdateOne(ctx, bson.M{"profile_id": profile.ProfileID}, bson.M{"$set": updateProfile})
	if err != nil {
		return nil, err
	}

	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}
	}

	return profile, err
}

func (p *profile) Get(ctx context.Context, id string) (*model.Profile, error) {
	profile := &model.Profile{}
	err := p.db.Collection(string(storage.Profile)).FindOne(ctx, bson.M{"profile_id": id}).Decode(profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNoRecordFound.New(errors.RecordNotfound)
		}
		return nil, err
	}
	return profile, nil
}
