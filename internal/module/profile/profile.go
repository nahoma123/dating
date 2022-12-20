package profile

import (
	"context"
	"dating/internal/constant"
	"dating/internal/constant/errors"
	"dating/internal/constant/model"
	"dating/internal/module"
	"dating/internal/storage"
	"dating/platform/logger"

	"go.uber.org/zap"
)

type profile struct {
	profileStorage storage.ProfileStorage
	logger         logger.Logger
}

func InitProfile(logger logger.Logger, profileStorage storage.ProfileStorage) module.ProfileModule {
	return &profile{
		logger:         logger,
		profileStorage: profileStorage,
	}
}

func (o *profile) RegisterUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	//
	if err := profile.ValidateRegisterProfile(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		o.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	profile, err := o.profileStorage.Create(ctx, profile)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return profile, nil
}

func (o *profile) UpdateUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	if err := profile.ValidateUpdateProfile(); err != nil {
		err = errors.ErrInvalidInput.Wrap(err, "invalid input")
		o.logger.Info(ctx, "invalid input", zap.Error(err))
		return nil, err
	}

	profile, err := o.profileStorage.Update(ctx, profile)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return profile, nil
}

func (o *profile) GetUserProfile(ctx context.Context, id string) (*model.Profile, error) {
	profile, err := o.profileStorage.Get(ctx, id)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return profile, nil
}

func (o *profile) GetUsers(ctx context.Context, id string) (*model.Profile, error) {
	profile, err := o.profileStorage.Get(ctx, id)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return profile, nil
}

func (o *profile) GetCustomers(ctx context.Context, filterPagination *constant.FilterPagination) ([]model.Profile, error) {
	customers, err := o.profileStorage.GetCustomers(ctx, filterPagination)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
