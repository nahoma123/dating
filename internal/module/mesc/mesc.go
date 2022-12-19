package mesc

import (
	"context"
	"dating/internal/constant"
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
	// 	err = errors.ErrInvalidInput.Wrap(err, "invalid input")
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
func (msc *mesc) CreateEthnicity(ctx context.Context, profile *model.Ethnicity) (*model.Ethnicity, error) {
	ethnicity, err := msc.mescStorage.CreateEthnicity(ctx, profile)
	if err != nil {
		msc.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return ethnicity, nil
}

// CreateState implements module.MescModule
func (msc *mesc) CreateState(ctx context.Context, state *model.State) (*model.State, error) {
	state, err := msc.mescStorage.CreateState(ctx, state)
	if err != nil {
		msc.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return state, nil
}

// GetCountries implements module.MescModule
func (msc *mesc) GetCountries(ctx context.Context, page int, perPage int) ([]*model.Country, *model.MetaData, error) {
	countries, metaData, err := msc.mescStorage.GetCountries(ctx, page, perPage)
	if err != nil {
		return nil, nil, err
	}
	return countries, metaData, nil
}

// GetEthnicities implements module.MescModule
func (msc *mesc) GetEthnicities(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Ethnicity, error) {
	ethnicities, err := msc.mescStorage.GetEthnicities(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return ethnicities, nil
}

// GetStates implements module.MescModule
func (msc *mesc) GetStates(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.State, error) {
	states, err := msc.mescStorage.GetStates(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return states, nil
}

// DeleteCountry implements module.MescModule
func (msc *mesc) DeleteCountry(ctx context.Context, countryId string) error {
	return msc.mescStorage.DeleteCountry(ctx, countryId)
}

// UpdateEthnicity implements module.MescModule
func (msc *mesc) DeleteEthnicity(ctx context.Context, id string) error {
	return msc.mescStorage.DeleteEthnicity(ctx, id)
}

// delete State implements module.MescModule
func (msc *mesc) DeleteState(ctx context.Context, id string) error {
	return msc.mescStorage.DeleteState(ctx, id)
}
