package persistence

import (
	"context"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/storage"
	"dating/platform/logger"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

// mes
type mesc struct {
	db *mongo.Database
}

func InitMescDB(db *mongo.Database) storage.MescStorage {
	return &mesc{
		db: db,
	}
}

func (mesc *mesc) CreateCountry(ctx context.Context, country *model.Country) (*model.Country, error) {
	// create database

	// createIndex( { "hostname": 1 }, { unique: true } )
	id, _ := uuid.NewV4()
	country.CountryId = id.String()

	country.CreatedAt = time.Now()
	country.UpdatedAt = time.Now()

	_, err := mesc.db.Collection(string(storage.Country)).InsertOne(ctx, country)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.CountryIsAlreadyRegistered)
		}
	}

	return country, err
}

func (mesc *mesc) DeleteCountry(ctx context.Context, id int) error {
	// Delete the document from the Country collection
	result, err := mesc.db.Collection(string(storage.Country)).DeleteOne(
		ctx,
		bson.M{"country_id": id},
	)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)
	}

	if result.DeletedCount == 0 {
		return errors.ErrDataExists.NewWithNoMessage()
	}

	return nil
}

func (mesc *mesc) GetCountries(ctx context.Context, page int, perPage int) (*model.Country, error) {
	cur, err := mesc.db.Collection(string(storage.State)).Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println("result-", result)
		}
		// do something with result....
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return nil, nil
}

// CreateEthnicity implements storage.MescStorage
func (*mesc) CreateEthnicity(ctx context.Context, profile *model.Ethnicity) (*model.Ethnicity, error) {
	panic("unimplemented")
}

// CreateState implements storage.MescStorage
func (*mesc) CreateState(ctx context.Context, profile *model.State) (*model.State, error) {
	panic("unimplemented")
}

// DeleteEthnicity implements storage.MescStorage
func (*mesc) DeleteEthnicity(ctx context.Context, id int) error {
	panic("unimplemented")
}

// DeleteState implements storage.MescStorage
func (*mesc) DeleteState(ctx context.Context, id int) error {
	panic("unimplemented")
}

// GetEthnicities implements storage.MescStorage
func (*mesc) GetEthnicities(ctx context.Context, page int, perPage int) (*model.Ethnicity, error) {
	panic("unimplemented")
}

// GetStates implements storage.MescStorage
func (*mesc) GetStates(ctx context.Context, page int, perPage int) (*model.State, error) {
	panic("unimplemented")
}
