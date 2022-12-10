package profile

import (
	"context"
	"dating/internal/constant/model"
	"dating/internal/module"
	"dating/platform/logger"
)

type profile struct {
	logger logger.Logger
}

func InitProfile(logger logger.Logger) module.ProfileModule {
	return &profile{
		logger,
	}
}

func (o *profile) GetUserProfile(ctx context.Context, Id string) (*model.Profile, error) {

	// logic from other microservice
	return nil, nil
}
func (o *profile) RegisterUserProfile(ctx context.Context, profile *model.Profile) error {
	//
	return nil
}
