package persistence

import (
	"context"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/storage"

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
	_, err := p.db.Collection(string(storage.Profile)).InsertOne(ctx, profile)
	if err != nil {
		return nil, errors.ErrDataExists.Wrap(err, "")
	}

	return profile, err
}
