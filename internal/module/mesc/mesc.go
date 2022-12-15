package mesc

import (
	"context"
	"dating/internal/constant/model"
	"dating/internal/module"
	"dating/internal/storage"
	"dating/platform/logger"
)

type mesc struct {
	mescStorage storage.MescStorage
	logger      logger.Logger
}

func InitMesc(logger logger.Logger, mescStorage storage.MescStorage) module.MescModule {
	return &mesc{
		logger:      logger,
		mescStorage: mescStorage,
	}
}

func (msc *mesc) CreateCountry(ctx context.Context, country *model.Country) (*model.Country, error) {
	//
	// if err := country.ValidateRegisterProfile(); err != nil {
	// 	err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
	// 	msc.logger.Info(ctx, "invalid input", zap.Error(err))
	// 	return nil, err
	// }

	country, err := msc.mescStorage.CreateCountry(ctx, country)
	if err != nil {
		msc.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return country, nil
}

// CreateEthnicity implements module.MescModule
func (*mesc) CreateEthnicity(ctx context.Context, profile *model.Country) (*model.Country, error) {
	panic("unimplemented")
}

// CreateState implements module.MescModule
func (*mesc) CreateState(ctx context.Context, profile *model.Country) (*model.Country, error) {
	panic("unimplemented")
}

// GetCountries implements module.MescModule
func (*mesc) GetCountries(ctx context.Context, page int, perPage int) (*model.Country, error) {
	panic("unimplemented")
}

// GetEthnicities implements module.MescModule
func (*mesc) GetEthnicities(ctx context.Context, page int, perPage int) (*model.Country, error) {
	panic("unimplemented")
}

// GetStates implements module.MescModule
func (*mesc) GetStates(ctx context.Context, page int, perPage int) (*model.Country, error) {
	panic("unimplemented")
}

// UpdateCountry implements module.MescModule
func (*mesc) UpdateCountry(ctx context.Context, profile *model.Country) (*model.Country, error) {
	panic("unimplemented")
}

// UpdateEthnicity implements module.MescModule
func (*mesc) UpdateEthnicity(ctx context.Context, profile *model.Country) (*model.Country, error) {
	panic("unimplemented")
}

// UpdateState implements module.MescModule
func (*mesc) UpdateState(ctx context.Context, profile *model.Country) (*model.Country, error) {
	panic("unimplemented")
}
