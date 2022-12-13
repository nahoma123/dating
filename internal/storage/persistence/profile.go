package persistence

import (
	"context"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/storage"
	"dating/platform/logger"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
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
	profile.ProfileID = id
	profile.CreatedAt = time.Now()
	_, err := p.db.Collection(string(storage.Profile)).InsertOne(ctx, profile)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.UserIsAlreadyRegistered)
		}
	}

	return profile, err
}
