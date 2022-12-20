package module

import (
	"context"
	"crypto/rsa"
	"dating/internal/constant"
	"dating/internal/constant/model"

	"github.com/golang-jwt/jwt/v4"
)

type AuthModule interface {
	VerifyToken(signingMethod jwt.SigningMethod, token string, pk *rsa.PublicKey) (bool, *jwt.RegisteredClaims)
	GetUserStatus(ctx context.Context, Id string) (string, error)
}

type ProfileModule interface {
	GetUserProfile(ctx context.Context, Id string) (*model.Profile, error)
	RegisterUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error)
	UpdateUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error)

	GetCustomers(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Profile, error)
}

type MescModule interface {
	CreateCountry(ctx context.Context, profile *model.Country) (*model.Country, error)
	DeleteCountry(ctx context.Context, countryId string) error
	GetCountries(ctx context.Context, page int, perPage int) ([]*model.Country, *model.MetaData, error)

	CreateState(ctx context.Context, profile *model.State) (*model.State, error)
	DeleteState(ctx context.Context, id string) error
	GetStates(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.State, error)

	CreateEthnicity(ctx context.Context, profile *model.Ethnicity) (*model.Ethnicity, error)
	DeleteEthnicity(ctx context.Context, id string) error
	GetEthnicities(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Ethnicity, error)
}
