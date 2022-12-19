package storage

import (
	"context"
	"dating/internal/constant/model"
)

type DatabaseCollection string

// database collection constants
const (
	Profile   DatabaseCollection = "profiles"
	Country   DatabaseCollection = "countries"
	State     DatabaseCollection = "states"
	Ethnicity DatabaseCollection = "ethnicities"
)

type ProfileStorage interface {
	Create(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	Update(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	Get(ctx context.Context, id string) (*model.Profile, error)
}

type MescStorage interface {
	CreateCountry(ctx context.Context, profile *model.Country) (*model.Country, error)
	DeleteCountry(ctx context.Context, id string) error
	GetCountries(ctx context.Context, page int, perPage int) ([]*model.Country, *model.MetaData, error)

	CreateState(ctx context.Context, profile *model.State) (*model.State, error)
	DeleteState(ctx context.Context, id int) error
	GetStates(ctx context.Context, page int, perPage int) (*model.State, error)

	CreateEthnicity(ctx context.Context, profile *model.Ethnicity) (*model.Ethnicity, error)
	DeleteEthnicity(ctx context.Context, id int) error
	GetEthnicities(ctx context.Context, page int, perPage int) (*model.Ethnicity, error)
}
