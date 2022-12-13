package profile

import (
	"context"
	"dating/internal/constant/model"
	"dating/internal/module"
	"dating/internal/storage"
	"dating/platform/logger"
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

func (o *profile) GetUserProfile(ctx context.Context, Id string) (*model.Profile, error) {

	// logic from other microservice
	return nil, nil
}
func (o *profile) RegisterUserProfile(ctx context.Context, profile *model.Profile) (*model.Profile, error) {
	//

	profile, err := o.profileStorage.Create(ctx, profile)
	if err != nil {
		o.logger.Warn(ctx, err.Error())
		return nil, err
	}
	return profile, nil
}
