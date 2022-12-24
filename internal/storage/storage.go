package storage

import (
	"context"
	"dating/internal/constant"
	"dating/internal/constant/model"
)

type DatabaseCollection string

// database collection constants
const (
	Profile   DatabaseCollection = "profiles"
	Country   DatabaseCollection = "countries"
	State     DatabaseCollection = "states"
	Ethnicity DatabaseCollection = "ethnicities"
	Like      DatabaseCollection = "likes"
	Favorite  DatabaseCollection = "favorites"
)

type ProfileStorage interface {
	Create(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	Update(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	Get(ctx context.Context, id string) (*model.Profile, error)
	GetCustomers(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Profile, error)

	LikeProfile(ctx context.Context, userID string, profileID string) error
	UnlikeProfile(ctx context.Context, userID string, profileID string) error

	MakeFavorite(ctx context.Context, userID string, profileID string) error
	RemoveFavorite(ctx context.Context, userID string, profileID string) error
}

type MescStorage interface {
	CreateCountry(ctx context.Context, profile *model.Country) (*model.Country, error)
	DeleteCountry(ctx context.Context, id string) error
	GetCountries(ctx context.Context, page int, perPage int) ([]*model.Country, *model.MetaData, error)

	DeleteState(ctx context.Context, id string) error
	CreateState(ctx context.Context, profile *model.State) (*model.State, error)
	GetStates(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.State, error)

	CreateEthnicity(ctx context.Context, profile *model.Ethnicity) (*model.Ethnicity, error)
	DeleteEthnicity(ctx context.Context, id string) error
	GetEthnicities(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Ethnicity, error)
}
