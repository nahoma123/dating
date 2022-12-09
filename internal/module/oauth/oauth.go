package oauth

import (
	"context"
	"dating/internal/module"
	"dating/platform/logger"
)

type oauth struct {
	logger logger.Logger
}

func InitOAuth(logger logger.Logger) module.AuthModule {
	return &oauth{
		logger,
	}
}

func (o *oauth) VerifyUserStatus(ctx context.Context, phone string) error {

	// logic from other microservice
	return nil
}
func (o *oauth) GetUserStatus(ctx context.Context, Id string) (string, error) {
	//
	return "", nil
}
