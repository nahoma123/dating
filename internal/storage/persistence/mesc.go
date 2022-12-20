package persistence

import (
	"context"
	"dating/internal/constant"
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
func (mesc *mesc) CreateEthnicity(ctx context.Context, ethnicity *model.Ethnicity) (*model.Ethnicity, error) {
	countryFilter := bson.M{"country_id": ethnicity.CountryId}
	var foundCountry model.Country
	err := mesc.db.Collection(string(storage.Country)).FindOne(ctx, countryFilter).Decode(&foundCountry)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.ErrNoRecordFound.Wrap(err, errors.UnableToFindCountryByTheIdProvided)
		}
		logger.Log().Error(ctx, err.Error())
		return nil, errors.ErrInternalServerError.New(errors.UnknownDbError)
	}

	id, _ := uuid.NewV4()
	ethnicity.EthnicityId = id.String()

	ethnicity.CreatedAt = time.Now()
	ethnicity.UpdatedAt = time.Now()

	_, err = mesc.db.Collection(string(storage.Ethnicity)).InsertOne(ctx, ethnicity)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.ErrDataExists.Wrap(err, errors.CountryIsAlreadyRegistered)
		}
	}

	return ethnicity, err
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
func (msc *mesc) DeleteEthnicity(ctx context.Context, id string) error {
	result, err := msc.db.Collection(string(storage.Ethnicity)).DeleteOne(
		ctx,
		bson.M{"ethnicity_id": id},
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

// DeleteState implements storage.MescStorage
func (msc *mesc) DeleteState(ctx context.Context, id string) error {
	result, err := msc.db.Collection(string(storage.State)).DeleteOne(
		ctx,
		bson.M{"state_id": id},
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

// GetEthnicities implements storage.MescStorage
func (mesc *mesc) GetEthnicities(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Ethnicity, error) {
	var results []bson.M

	err := constant.GetResults(ctx, mesc.db, string(storage.Ethnicity), filterPagination, &results, nil)
	if err != nil {
		return nil, err
	}

	var ethnicities []model.Ethnicity
	for _, result := range results {
		var ethnicity model.Ethnicity
		b, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(b, &ethnicity)
		if err != nil {
			return nil, err
		}
		ethnicities = append(ethnicities, ethnicity)
	}
	return ethnicities, nil
}

// GetStates implements storage.MescStorage
func (mesc *mesc) GetStates(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.State, error) {
	var results []bson.M

	err := constant.GetResults(ctx, mesc.db, string(storage.State), filterPagination, &results, nil)
	if err != nil {
		return nil, err
	}

	var states []model.State
	for _, result := range results {
		var state model.State
		b, err := bson.Marshal(result)
		if err != nil {
			return nil, err
		}
		err = bson.Unmarshal(b, &state)
		if err != nil {
			return nil, err
		}
		states = append(states, state)
	}
	return states, nil
}
