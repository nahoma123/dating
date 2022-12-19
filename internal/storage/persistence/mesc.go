package persistence

import (
	"context"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/storage"
	"dating/platform/logger"
	"math"
	"time"

	"github.com/gofrs/uuid"
	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (mesc *mesc) DeleteCountry(ctx context.Context, id string) error {
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
		return errors.ErrNoRecordFound.New(errors.RecordNotfound)
	}

	return nil
}

func (mesc *mesc) GetCountries(ctx context.Context, page int, perPage int) ([]*model.Country, *model.MetaData, error) {
	var countries []*model.Country

	// Set the pagination parameters for the query
	skip := int64((page - 1) * perPage)
	limit := int64(perPage)

	// Find all documents in the "Country" collection, skipping the specified number of
	// documents and limiting the result to the specified number of documents
	filter := bson.D{}
	cur, err := mesc.db.Collection(string(storage.Country)).Find(ctx, filter, options.Find().SetSkip(skip).SetLimit(limit))
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, nil, errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &countries); err != nil {
		return nil, nil, errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)
	}

	// Retrieve the total number of documents in the "Country" collection
	totalCount, err := mesc.db.Collection(string(storage.Country)).CountDocuments(ctx, filter)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, nil, errors.ErrInternalServerError.Wrap(err, errors.UnknownDbError)
	}

	// Calculate the total number of pages
	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

	// Create the MetaData struct
	metaData := &model.MetaData{
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
		TotalCount: int(totalCount),
	}

	return countries, metaData, nil
}

// CreateEthnicity implements storage.MescStorage
func (*mesc) CreateEthnicity(ctx context.Context, profile *model.Ethnicity) (*model.Ethnicity, error) {
	panic("unimplemented")
}

// CreateState implements storage.MescStorage
func (mesc *mesc) CreateState(ctx context.Context, state *model.State) (*model.State, error) {
	// Check if the country with the specified ID exists
	countryFilter := bson.M{"country_id": state.CountryId}
	var foundCountry model.Country
	err := mesc.db.Collection(string(storage.Country)).FindOne(ctx, countryFilter).Decode(&foundCountry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNoRecordFound.Wrap(err, errors.UnableToFindCountryByTheIdProvided)
		}
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New(errors.UnknownDbError)
	}

	// create database

	// createIndex( { "hostname": 1 }, { unique: true } )
	id, _ := uuid.NewV4()
	state.StateId = id.String()

	state.CreatedAt = time.Now()
	state.UpdatedAt = time.Now()

	_, err = mesc.db.Collection(string(storage.State)).InsertOne(ctx, state)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.StateIsAlreadyRegistered)
		}
	}

	return state, err
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
